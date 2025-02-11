package ysepay

import (
	"errors"

	"github.com/yiigo/sdk-go/internal/value"
)

type V = value.V

// ErrSysAccepting 网关受理中
var ErrSysAccepting = errors.New("SYS001 | 网关受理中")

const (
	SysOK        = "SYS000" // 网关受理成功响应码
	SysAccepting = "SYS001" // 网关受理中响应码

	ComOK         = "COM000" // 业务受理成功
	ComProcessing = "COM004" // 业务处理中
)
