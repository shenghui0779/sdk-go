package utils

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"sync"
)

// CDATA XML CDATA section which is defined as blocks of text that are not parsed by the parser, but are otherwise recognized as markup.
type CDATA string

// MarshalXML encodes the receiver as zero or more XML elements.
func (c CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}

// ErrXMLWriterNil ...
var ErrXMLWriterNil = errors.New("nil xml writer")

// ErrXMLReaderNil ...
var ErrXMLReaderNil = errors.New("nil xml reader")

// WXML 微信返回结果
type WXML map[string]string

var ErrHTTPClientNil = errors.New("nil http client")

var bufferPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 16<<10)) // 16KB
	},
}

// PostXML XML POST 请求
func PostXML(client *HTTPClient, url string, body WXML) (WXML, error) {
	if client == nil {
		return nil, ErrHTTPClientNil
	}

	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()

	defer bufferPool.Put(buf)

	if err := FormatMap2XML(buf, body); err != nil {
		return nil, err
	}

	resp, err := client.Post(url, buf.Bytes(), WithRequestHeader("Content-Type", "text/xml; charset=utf-8"))

	if err != nil {
		return nil, err
	}

	wxml, err := ParseXML2Map(bytes.NewReader(resp))

	if err != nil {
		return nil, err
	}

	return wxml, nil
}

// FormatMap2XML format map to xml
func FormatMap2XML(xmlWriter io.Writer, m WXML) (err error) {
	if xmlWriter == nil {
		err = ErrXMLWriterNil

		return
	}

	if _, err = io.WriteString(xmlWriter, "<xml>"); err != nil {
		return
	}

	for k, v := range m {
		if _, err = io.WriteString(xmlWriter, fmt.Sprintf("<%s>", k)); err != nil {
			return
		}

		if err = xml.EscapeText(xmlWriter, []byte(v)); err != nil {
			return
		}

		if _, err = io.WriteString(xmlWriter, fmt.Sprintf("</%s>", k)); err != nil {
			return
		}
	}

	if _, err = io.WriteString(xmlWriter, "</xml>"); err != nil {
		return
	}

	return
}

// ParseXML2Map parse xml to map
func ParseXML2Map(xmlReader io.Reader) (m WXML, err error) {
	if xmlReader == nil {
		err = ErrXMLReaderNil

		return
	}

	m = make(WXML)

	var (
		d     = xml.NewDecoder(xmlReader)
		tk    xml.Token
		depth = 0 // current xml.Token depth
		key   string
		buf   bytes.Buffer
	)

	for {
		tk, err = d.Token()

		if err != nil {
			if err == io.EOF {
				err = nil

				return
			}

			return
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

					return
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
