package mch

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/shenghui0779/gochat/utils"
	"golang.org/x/crypto/pkcs12"
)

// WXMch 微信商户
type WXMch struct {
	appid      string
	mchid      string
	apikey     string
	tlsOptions []utils.TLSOption
	client     *utils.WXClient
	tlsClient  *utils.WXClient
}

func New(appid, mchid, apikey string) *WXMch {
	mch := &WXMch{
		appid:      appid,
		mchid:      mchid,
		apikey:     apikey,
		tlsOptions: make([]utils.TLSOption, 0),
	}

	mch.client = utils.NewWXClient(utils.WithInsecureSkipVerify(true))
	mch.tlsClient = utils.NewWXClient(utils.WithInsecureSkipVerify(true))

	return mch
}

// LoadP12CertFromFile load p12(pfx) cert from file
func (wx *WXMch) LoadP12CertFromFile(path string) error {
	p12, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	cert, err := wx.pkcs12ToPem(p12)

	if err != nil {
		return err
	}

	wx.tlsClient = utils.NewWXClient(utils.WithCertificates(cert), utils.WithInsecureSkipVerify(true))

	return nil
}

// LoadP12CertFromBase64 load p12(pfx) cert from base64
func (wx *WXMch) LoadP12CertFromBase64(cert string) error {
	p12, err := base64.StdEncoding.DecodeString(cert)

	if err != nil {
		return err
	}

	certificate, err := wx.pkcs12ToPem(p12)

	if err != nil {
		return err
	}

	wx.tlsClient = utils.NewWXClient(utils.WithCertificates(certificate), utils.WithInsecureSkipVerify(true))

	return nil
}

// Order returns new order
func (wx *WXMch) Order(options ...utils.RequestOption) *Order {
	return &Order{
		mch:     wx,
		options: options,
	}
}

// Refund returns new refund
func (wx *WXMch) Refund(options ...utils.RequestOption) *Refund {
	return &Refund{
		mch:     wx,
		options: options,
	}
}

// Pappay returns new pappay
func (wx *WXMch) Pappay(options ...utils.RequestOption) *Pappay {
	return &Pappay{
		mch:     wx,
		options: options,
	}
}

// Transfer returns new transfer
func (wx *WXMch) Transfer(options ...utils.RequestOption) *Transfer {
	return &Transfer{
		mch:     wx,
		options: options,
	}
}

// Redpack returns new redpack
func (wx *WXMch) Redpack(options ...utils.RequestOption) *Redpack {
	return &Redpack{
		mch:     wx,
		options: options,
	}
}

// APPAPI 用于APP拉起支付
func (wx *WXMch) APPAPI(prepayID string) utils.WXML {
	ch := utils.WXML{
		"appid":     wx.appid,
		"partnerid": wx.mchid,
		"prepayid":  prepayID,
		"package":   "Sign=WXPay",
		"noncestr":  utils.NonceStr(),
		"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
	}

	ch["sign"] = SignWithMD5(ch, wx.apikey)

	return ch
}

// JSAPI 用于JS拉起支付
func (wx *WXMch) JSAPI(prepayID string) utils.WXML {
	ch := utils.WXML{
		"appId":     wx.appid,
		"nonceStr":  utils.NonceStr(),
		"package":   fmt.Sprintf("prepay_id=%s", prepayID),
		"signType":  SignMD5,
		"timeStamp": strconv.FormatInt(time.Now().Unix(), 10),
	}

	ch["paySign"] = SignWithMD5(ch, wx.apikey)

	return ch
}

// VerifyWXReply 验证微信结果
func (wx *WXMch) VerifyWXReply(reply utils.WXML) error {
	if wxsign, ok := reply["sign"]; ok {
		signType := SignMD5

		if v, ok := reply["sign_type"]; ok {
			signType = v
		}

		signature := ""

		switch signType {
		case SignMD5:
			signature = SignWithMD5(reply, wx.apikey)
		case SignHMacSHA256:
			signature = SignWithHMacSHA256(reply, wx.apikey)
		default:
			return fmt.Errorf("invalid sign type: %s", signType)
		}

		if wxsign != signature {
			return fmt.Errorf("signature verified failed, want: %s, got: %s", signature, wxsign)
		}
	}

	if appid, ok := reply["appid"]; ok {
		if appid != wx.appid {
			return fmt.Errorf("appid mismatch, want: %s, got: %s", wx.appid, reply["appid"])
		}
	}

	if mchid, ok := reply["mch_id"]; ok {
		if mchid != wx.mchid {
			return fmt.Errorf("mchid mismatch, want: %s, got: %s", wx.mchid, reply["mch_id"])
		}
	}

	return nil
}

// RSAPublicKey 获取RSA加密公钥
func (wx *WXMch) RSAPublicKey(options ...utils.RequestOption) ([]byte, error) {
	body := utils.WXML{
		"mch_id":    wx.mchid,
		"nonce_str": utils.NonceStr(),
		"sign_type": SignMD5,
	}

	resp, err := wx.tlsPost(TransferBalanceOrderQueryURL, body, options...)

	if err != nil {
		return nil, err
	}

	pubKey, ok := resp["pub_key"]

	if !ok {
		return nil, errors.New("empty pub_key")
	}

	return []byte(pubKey), nil
}

func (wx *WXMch) pkcs12ToPem(p12 []byte) (tls.Certificate, error) {
	blocks, err := pkcs12.ToPEM(p12, wx.mchid)

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

func (wx *WXMch) post(reqURL string, body utils.WXML, options ...utils.RequestOption) (utils.WXML, error) {
	body["sign"] = SignWithMD5(body, wx.apikey)

	resp, err := wx.client.PostXML(reqURL, body, options...)

	if err != nil {
		return nil, err
	}

	if resp["return_code"] != ResultSuccess {
		return nil, errors.New(resp["return_msg"])
	}

	if err := wx.VerifyWXReply(resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (wx *WXMch) tlsPost(reqURL string, body utils.WXML, options ...utils.RequestOption) (utils.WXML, error) {
	body["sign"] = SignWithMD5(body, wx.apikey)

	resp, err := wx.tlsClient.PostXML(reqURL, body, options...)

	if err != nil {
		return nil, err
	}

	if resp["return_code"] != ResultSuccess {
		return nil, errors.New(resp["return_msg"])
	}

	if err := wx.VerifyWXReply(resp); err != nil {
		return nil, err
	}

	return resp, nil
}
