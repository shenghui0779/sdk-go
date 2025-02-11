package wechat

import (
	"github.com/tidwall/gjson"
)

type X map[string]any

type Form map[string]string

// APIResult API结果 (支付v3)
type APIResult struct {
	Code int // HTTP状态码
	Body gjson.Result
}

// DownloadResult 资源下载结果 (支付v3)
type DownloadResult struct {
	HashType  string
	HashValue string
	Buffer    []byte
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
