package wechat

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// AppId = "wxba5fad812f8e6fb9"
//
// 服务器配置：
// URL = https://www.qq.com/revice
// Token = "AAAAA"
// EncodingAESKey = "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
// 数据格式为JSON
//
// 推送的URL链接：
// https://www.qq.com/recive?signature=6c5c811b55cc85e0e1b54100749188c20beb3f5d&timestamp=1714112445&nonce=415670741&openid=o9AgO5Kd5ggOC-bXrbNODIiE3bGY&encrypt_type=aes&msg_signature=046e02f8204d34f8ba5fa3b1db94908f3df2e9b3
//
// 推送的包体：
// {
//     "ToUserName": "gh_97417a04a28d",
//     "Encrypt": "+qdx1OKCy+5JPCBFWw70tm0fJGb2Jmeia4FCB7kao+/Q5c/ohsOzQHi8khUOb05JCpj0JB4RvQMkUyus8TPxLKJGQqcvZqzDpVzazhZv6JsXUnnR8XGT740XgXZUXQ7vJVnAG+tE8NUd4yFyjPy7GgiaviNrlCTj+l5kdfMuFUPpRSrfMZuMcp3Fn2Pede2IuQrKEYwKSqFIZoNqJ4M8EajAsjLY2km32IIjdf8YL/P50F7mStwntrA2cPDrM1kb6mOcfBgRtWygb3VIYnSeOBrebufAlr7F9mFUPAJGj04="
// }

func Test_OA_VerifyEventMsg(t *testing.T) {
	oa := NewOfficialAccount("wxba5fad812f8e6fb9", "secret", WithOASrvCfg("AAAAA", "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"))
	err := oa.VerifyEventMsg("046e02f8204d34f8ba5fa3b1db94908f3df2e9b3", "1714112445", "415670741", "+qdx1OKCy+5JPCBFWw70tm0fJGb2Jmeia4FCB7kao+/Q5c/ohsOzQHi8khUOb05JCpj0JB4RvQMkUyus8TPxLKJGQqcvZqzDpVzazhZv6JsXUnnR8XGT740XgXZUXQ7vJVnAG+tE8NUd4yFyjPy7GgiaviNrlCTj+l5kdfMuFUPpRSrfMZuMcp3Fn2Pede2IuQrKEYwKSqFIZoNqJ4M8EajAsjLY2km32IIjdf8YL/P50F7mStwntrA2cPDrM1kb6mOcfBgRtWygb3VIYnSeOBrebufAlr7F9mFUPAJGj04=")
	assert.Nil(t, err)
}

func Test_OA_DecodeEventMsg(t *testing.T) {
	oa := NewOfficialAccount("wxba5fad812f8e6fb9", "secret", WithOASrvCfg("AAAAA", "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"))
	b, err := oa.DecodeEventMsg("+qdx1OKCy+5JPCBFWw70tm0fJGb2Jmeia4FCB7kao+/Q5c/ohsOzQHi8khUOb05JCpj0JB4RvQMkUyus8TPxLKJGQqcvZqzDpVzazhZv6JsXUnnR8XGT740XgXZUXQ7vJVnAG+tE8NUd4yFyjPy7GgiaviNrlCTj+l5kdfMuFUPpRSrfMZuMcp3Fn2Pede2IuQrKEYwKSqFIZoNqJ4M8EajAsjLY2km32IIjdf8YL/P50F7mStwntrA2cPDrM1kb6mOcfBgRtWygb3VIYnSeOBrebufAlr7F9mFUPAJGj04=")
	assert.Nil(t, err)
	fmt.Println(string(b))
}
