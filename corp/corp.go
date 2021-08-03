package corp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/shenghui0779/gochat/wx"
	"github.com/shenghui0779/yiigo"
	"github.com/tidwall/gjson"
)

type Corp struct {
	corpid         string
	token          string
	encodingAESKey string
	nonce          func(size uint) string
	client         wx.Client
}

func NewCorp(corpid string) *Corp {
	return &Corp{
		corpid: corpid,
		nonce:  wx.Nonce,
		client: wx.NewClient(wx.WithInsecureSkipVerify()),
	}
}

// SetServerConfig 设置服务器配置
// [参考](https://open.work.weixin.qq.com/api/doc/90000/90135/90930)
func (corp *Corp) SetServerConfig(token, encodingAESKey string) {
	corp.token = token
	corp.encodingAESKey = encodingAESKey
}

func (corp *Corp) CorpID() string {
	return corp.corpid
}

func (corp *Corp) AccessToken(ctx context.Context, secret string, options ...yiigo.HTTPOption) (*AccessToken, error) {
	resp, err := corp.client.Get(ctx, fmt.Sprintf("%s?corpid=%s&corpsecret=%s", CgiBinAccessTokenURL, corp.corpid, secret), options...)

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if code := r.Get("errcode").Int(); code != 0 {
		return nil, fmt.Errorf("%d|%s", code, r.Get("errmsg").String())
	}

	token := new(AccessToken)

	if err = json.Unmarshal(resp, token); err != nil {
		return nil, err
	}

	return token, nil
}

// Do exec action
func (corp *Corp) Do(ctx context.Context, accessToken string, action wx.Action, options ...yiigo.HTTPOption) error {
	var (
		resp []byte
		err  error
	)

	switch action.Method() {
	case wx.MethodGet:
		resp, err = corp.client.Get(ctx, action.URL(accessToken), options...)
	case wx.MethodPost:
		body, berr := action.Body()

		if berr != nil {
			return berr
		}

		resp, err = corp.client.Post(ctx, action.URL(accessToken), body, options...)
	case wx.MethodUpload:
		form, ferr := action.UploadForm()

		if ferr != nil {
			return ferr
		}

		resp, err = corp.client.Upload(ctx, action.URL(accessToken), form, options...)
	}

	if err != nil {
		return err
	}

	r := gjson.ParseBytes(resp)

	if code := r.Get("errcode").Int(); code != 0 {
		return fmt.Errorf("%d|%s", code, r.Get("errmsg").String())
	}

	if action.Decode() == nil {
		return nil
	}

	return action.Decode()(resp)
}
