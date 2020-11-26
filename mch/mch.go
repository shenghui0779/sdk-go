package mch

import (
	"crypto/rand"
	"crypto/tls"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"strconv"

	"golang.org/x/crypto/pkcs12"
)

// Action wechat pay action
type Action interface {
	WXML() func(appid, mchid, apikey, nonce string) (WXML, error)
	URL() string
	TLS() bool
}

// WechatAPI 微信支付API
type WechatAPI struct {
	wxml func(appid, mchid, apikey, nonce string) (WXML, error)
	url  string
	tls  bool
}

func (a *WechatAPI) WXML() func(appid, mchid, apikey, nonce string) (WXML, error) {
	return a.wxml
}

func (a *WechatAPI) URL() string {
	return a.url
}

func (a *WechatAPI) TLS() bool {
	return a.tls
}

// WechatMch 微信支付
type WechatMch struct {
	appid     string
	mchid     string
	apikey    string
	nonce     func(size int) string
	client    HTTPClient
	tlsClient HTTPClient
}

// New returns new wechat pay
func New(appid, mchid, apikey string) *WechatMch {
	c := NewHTTPClient(&tls.Config{InsecureSkipVerify: true})

	return &WechatMch{
		appid:  appid,
		mchid:  mchid,
		apikey: apikey,
		nonce: func(size int) string {
			nonce := make([]byte, size/2)
			io.ReadFull(rand.Reader, nonce)

			return hex.EncodeToString(nonce)
		},
		client:    c,
		tlsClient: c,
	}
}

// LoadCertFromP12File load cert from p12(pfx) file
func (w *WechatMch) LoadCertFromP12File(path string) error {
	p12, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	cert, err := w.pkcs12ToPem(p12)

	if err != nil {
		return err
	}

	w.tlsClient = NewHTTPClient(&tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	})

	return nil
}

// LoadCertFromPemFile load cert from PEM file
func (w *WechatMch) LoadCertFromPemFile(certFile, keyFile string) error {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)

	if err != nil {
		return err
	}

	w.tlsClient = NewHTTPClient(&tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	})

	return nil
}

// LoadCertFromPemBlock load cert from a pair of PEM encoded data
func (w *WechatMch) LoadCertFromPemBlock(certPEMBlock, keyPEMBlock []byte) error {
	cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)

	if err != nil {
		return err
	}

	w.tlsClient = NewHTTPClient(&tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	})

	return nil
}

// Do exec action
func (w *WechatMch) Do(action Action, options ...HTTPOption) (WXML, error) {
	body, err := action.WXML()(w.appid, w.mchid, w.apikey, w.nonce(16))

	if err != nil {
		return nil, err
	}

	if action.URL() == "" {
		return body, nil
	}

	var resp WXML

	if action.TLS() {
		resp, err = w.tlsClient.PostXML(action.URL(), body, options...)
	} else {
		resp, err = w.client.PostXML(action.URL(), body, options...)
	}

	if err != nil {
		return nil, err
	}

	if resp["return_code"] != ResultSuccess {
		return nil, errors.New(resp["return_msg"])
	}

	if err := w.VerifyWXReply(resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// APPAPI 用于APP拉起支付
func (w *WechatMch) APPAPI(prepayID string, timestamp int64) WXML {
	m := WXML{
		"appid":     w.appid,
		"partnerid": w.mchid,
		"prepayid":  prepayID,
		"package":   "Sign=WXPay",
		"noncestr":  w.nonce(16),
		"timestamp": strconv.FormatInt(timestamp, 10),
	}

	m["sign"] = SignWithMD5(m, w.apikey, true)

	return m
}

// JSAPI 用于JS拉起支付
func (w *WechatMch) JSAPI(prepayID string, timestamp int64) WXML {
	m := WXML{
		"appId":     w.appid,
		"nonceStr":  w.nonce(16),
		"package":   fmt.Sprintf("prepay_id=%s", prepayID),
		"signType":  SignMD5,
		"timeStamp": strconv.FormatInt(timestamp, 10),
	}

	m["paySign"] = SignWithMD5(m, w.apikey, true)

	return m
}

// MPRedpackJSAPI 小程序领取红包
func (w *WechatMch) MPRedpackJSAPI(pkg string, timestamp int64) WXML {
	m := WXML{
		"appId":     w.appid,
		"nonceStr":  w.nonce(16),
		"package":   url.QueryEscape(pkg),
		"timeStamp": strconv.FormatInt(timestamp, 10),
	}

	m["paySign"] = SignWithMD5(m, w.apikey, false)

	delete(m, "appId")
	m["signType"] = SignMD5

	return m
}

// VerifyWXReply 验证微信结果
func (w *WechatMch) VerifyWXReply(m WXML) error {
	if wxsign, ok := m["sign"]; ok {
		signature := ""

		if v, ok := m["sign_type"]; ok && v == SignHMacSHA256 {
			signature = SignWithHMacSHA256(m, w.apikey, true)
		} else {
			signature = SignWithMD5(m, w.apikey, true)
		}

		if wxsign != signature {
			return fmt.Errorf("signature verified failed, want: %s, got: %s", signature, wxsign)
		}
	}

	if appid, ok := m["appid"]; ok {
		if appid != w.appid {
			return fmt.Errorf("appid mismatch, want: %s, got: %s", w.appid, m["appid"])
		}
	}

	if mchid, ok := m["mch_id"]; ok {
		if mchid != w.mchid {
			return fmt.Errorf("mchid mismatch, want: %s, got: %s", w.mchid, m["mch_id"])
		}
	}

	return nil
}

func (w *WechatMch) pkcs12ToPem(p12 []byte) (tls.Certificate, error) {
	blocks, err := pkcs12.ToPEM(p12, w.mchid)

	if err != nil {
		return tls.Certificate{}, err
	}

	pemData := make([]byte, 0)

	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}

	// then use PEM data for tls to construct tls certificate:
	return tls.X509KeyPair(pemData, pemData)
}
