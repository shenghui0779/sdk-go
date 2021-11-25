package wx

import (
	"bytes"
	"crypto/rand"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/pkcs12"
)

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

// Nonce returns nonce string, param `size` better for even number.
func Nonce(size uint) string {
	nonce := make([]byte, size/2)
	io.ReadFull(rand.Reader, nonce)

	return hex.EncodeToString(nonce)
}

// FormatMap2XML format map to xml
func FormatMap2XML(m WXML) ([]byte, error) {
	var builder strings.Builder

	builder.WriteString("<xml>")

	for k, v := range m {
		builder.WriteString(fmt.Sprintf("<%s>", k))

		if err := xml.EscapeText(&builder, []byte(v)); err != nil {
			return nil, err
		}

		builder.WriteString(fmt.Sprintf("</%s>", k))
	}

	builder.WriteString("</xml>")

	return []byte(builder.String()), nil
}

// FormatMap2XML format map to xml with sorted keys for test
// func FormatMap2XML(m WXML) ([]byte, error) {
// 	ks := make([]string, 0, len(m))

// 	for k := range m {
// 		ks = append(ks, k)
// 	}

// 	sort.Strings(ks)

// 	var builder strings.Builder

// 	builder.WriteString("<xml>")

// 	for _, k := range ks {
// 		builder.WriteString(fmt.Sprintf("<%s>", k))

// 		if err := xml.EscapeText(&builder, []byte(m[k])); err != nil {
// 			return nil, err
// 		}

// 		builder.WriteString(fmt.Sprintf("</%s>", k))
// 	}

// 	builder.WriteString("</xml>")

// 	return []byte(builder.String()), nil
// }

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

// MarshalWithNoEscapeHTML marshal with no escape HTML
func MarshalWithNoEscapeHTML(v interface{}) ([]byte, error) {
	var buf bytes.Buffer

	jsonEncoder := json.NewEncoder(&buf)
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

// P12FileToCert 通过p12(pfx)证书文件生成Pem证书
func P12FileToCert(path, password string) (tls.Certificate, error) {
	fail := func(err error) (tls.Certificate, error) { return tls.Certificate{}, err }

	certPath, err := filepath.Abs(filepath.Clean(path))

	if err != nil {
		return fail(err)
	}

	p12, err := ioutil.ReadFile(certPath)

	if err != nil {
		return fail(err)
	}

	return P12BlockToCert(p12, password)
}

// P12BlockToCert 通过p12(pfx)证书内容生成Pem证书
func P12BlockToCert(pfxData []byte, password string) (tls.Certificate, error) {
	blocks, err := pkcs12.ToPEM(pfxData, password)

	if err != nil {
		return tls.Certificate{}, err
	}

	pemData := make([]byte, 0)

	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}

	return tls.X509KeyPair(pemData, pemData)
}
