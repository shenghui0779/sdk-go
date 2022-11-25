package wx

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/crypto/pkcs12"
)

// M is a convenient alias for a map[string]interface{}.
type M map[string]interface{}

// WXML deal with xml for wechat
type WXML map[string]string

// CDATA XML CDATA section which is defined as blocks of text that are not parsed by the parser, but are otherwise recognized as markup.
type CDATA string

// MarshalXML encodes the receiver as zero or more XML elements.
func (c CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}

// 签名类型
type SignType string

func (st SignType) Do(key string, m WXML, toUpper bool) string {
	data := make([]string, 0, len(m))

	for k, v := range m {
		if k != "sign" && len(v) != 0 {
			data = append(data, k+"="+v)
		}
	}

	sort.Strings(data)

	data = append(data, "key="+key)

	sign := ""

	if st == SignHMacSHA256 {
		sign = HMacSHA256(strings.Join(data, "&"), key)
	} else {
		sign = MD5(strings.Join(data, "&"))
	}

	if toUpper {
		sign = strings.ToUpper(sign)
	}

	return sign
}

const (
	SignMD5        SignType = "MD5"
	SignHMacSHA256 SignType = "HMAC-SHA256"
)

// Nonce returns nonce string, param `size` better for even number.
func Nonce(size uint) string {
	nonce := make([]byte, size/2)
	io.ReadFull(rand.Reader, nonce)

	return hex.EncodeToString(nonce)
}

// MD5 calculates the md5 hash of a string.
func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}

// SHA256 calculates the sha256 hash of a string.
func SHA256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}

// HMacSHA256 generates a keyed sha256 hash value.
func HMacSHA256(s, key string) string {
	mac := hmac.New(sha256.New, []byte(key))

	mac.Write([]byte(s))

	return hex.EncodeToString(mac.Sum(nil))
}

// FormatMap2XML format map to xml
func FormatMap2XML(m WXML) ([]byte, error) {
	var builder strings.Builder

	builder.WriteString("<xml>")

	for k, v := range m {
		builder.WriteString("<" + k + ">")

		if err := xml.EscapeText(&builder, []byte(v)); err != nil {
			return nil, err
		}

		builder.WriteString("</" + k + ">")
	}

	builder.WriteString("</xml>")

	return []byte(builder.String()), nil
}

// FormatMap2XMLForTest 用于单元测试
func FormatMap2XMLForTest(m WXML) ([]byte, error) {
	ks := make([]string, 0, len(m))

	for k := range m {
		ks = append(ks, k)
	}

	sort.Strings(ks)

	var builder strings.Builder

	builder.WriteString("<xml>")

	for _, k := range ks {
		builder.WriteString("<" + k + ">")

		if err := xml.EscapeText(&builder, []byte(m[k])); err != nil {
			return nil, err
		}

		builder.WriteString("</" + k + ">")
	}

	builder.WriteString("</xml>")

	xmlStr := builder.String()

	fmt.Println("[XML]", xmlStr)

	return []byte(xmlStr), nil
}

// ParseXML2Map parse xml to map
func ParseXML2Map(b []byte) (WXML, error) {
	m := make(WXML)

	xmlReader := bytes.NewReader(b)

	var (
		d     = xml.NewDecoder(xmlReader)
		tk    xml.Token
		depth = 0 // current xml.Token depth
		key   string
		buf   bytes.Buffer
		err   error
	)

	d.Strict = false

	for {
		tk, err = d.Token()

		if err != nil {
			if err == io.EOF {
				return m, nil
			}

			return nil, err
		}

		switch v := tk.(type) {
		case xml.StartElement:
			depth++

			switch depth {
			case 2:
				key = v.Name.Local
				buf.Reset()
			case 3:
				if err = d.Skip(); err != nil {
					return nil, err
				}

				depth--
				key = "" // key == "" indicates that the node with depth==2 has children
			}
		case xml.CharData:
			if depth == 2 && key != "" {
				buf.Write(v)
			}
		case xml.EndElement:
			if depth == 2 && key != "" {
				m[key] = buf.String()
			}

			depth--
		}
	}
}

// EncodeUint32ToBytes 把整数 uint32 格式化成 4 字节的网络字节序
func EncodeUint32ToBytes(i uint32) []byte {
	b := make([]byte, 4)

	b[0] = byte(i >> 24)
	b[1] = byte(i >> 16)
	b[2] = byte(i >> 8)
	b[3] = byte(i)

	return b
}

// DecodeBytesToUint32 从 4 字节的网络字节序里解析出整数 uint32
func DecodeBytesToUint32(b []byte) uint32 {
	if len(b) != 4 {
		return 0
	}

	return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
}

// MarshalNoEscapeHTML marshal with no escape HTML
func MarshalNoEscapeHTML(v interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)

	jsonEncoder := json.NewEncoder(buf)
	jsonEncoder.SetEscapeHTML(false)

	if err := jsonEncoder.Encode(v); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	// 去掉 go std 给末尾加的 '\n'
	// @see https://github.com/golang/go/issues/7767
	if l := len(b); l != 0 && b[l-1] == '\n' {
		b = b[:l-1]
	}

	return b, nil
}

// LoadP12Cert 通过p12(pfx)证书文件生成Pem证书
func LoadP12Cert(pfxfile, mchid string) (tls.Certificate, error) {
	fail := func(err error) (tls.Certificate, error) { return tls.Certificate{}, err }

	certPath, err := filepath.Abs(filepath.Clean(pfxfile))

	if err != nil {
		return fail(err)
	}

	pfxdata, err := ioutil.ReadFile(certPath)

	if err != nil {
		return fail(err)
	}

	blocks, err := pkcs12.ToPEM(pfxdata, mchid)

	if err != nil {
		return fail(err)
	}

	pemData := make([]byte, 0)

	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}

	return tls.X509KeyPair(pemData, pemData)
}
