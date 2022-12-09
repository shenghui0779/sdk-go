package offia

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/urls"
)

func TestCreateQRCode(t *testing.T) {
	body := []byte(`{"action_name":"QR_SCENE","action_info":{"scene":{"scene_id":123}},"expire_seconds":60}`)

	resp := []byte(`{
	"ticket": "gQH47joAAAAAAAAAASxodHRwOi8vd2VpeGluLnFxLmNvbS9xL2taZ2Z3TVRtNzJXV1Brb3ZhYmJJAAIEZ23sUwMEmm3sUw==",
	"expire_seconds": 60,
	"url": "http://weixin.qq.com/q/kZgfwMTm72WWPkovabbI"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET", WithMockClient(client))

	params := &ParamsQRCodeCreate{
		ActionName: QRScene,
		ActionInfo: &QRCodeActionInfo{
			Scene: &QRCodeScene{
				SceneID: 123,
			},
		},
		ExpireSeconds: 60,
	}
	result := new(ResultQRCodeCreate)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", CreateQRCode(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultQRCodeCreate{
		Ticket:        "gQH47joAAAAAAAAAASxodHRwOi8vd2VpeGluLnFxLmNvbS9xL2taZ2Z3TVRtNzJXV1Brb3ZhYmJJAAIEZ23sUwMEmm3sUw==",
		ExpireSeconds: 60,
		URL:           "http://weixin.qq.com/q/kZgfwMTm72WWPkovabbI",
	}, result)
}

func TestShowQRCode(t *testing.T) {
	resp := []byte("BUFFER")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=TICKET", nil).Return(resp, nil)

	b, err := client.Do(context.TODO(), http.MethodGet, fmt.Sprintf("%s?ticket=%s", urls.OffiaQRCodeShow, url.QueryEscape("TICKET")), nil)

	assert.Nil(t, err)

	buf := bytes.NewBuffer(nil)

	if _, err = io.Copy(buf, bytes.NewReader(b)); err != nil {
		assert.Nil(t, err)
	}

	assert.Equal(t, "QlVGRkVS", base64.StdEncoding.EncodeToString(buf.Bytes()))
}

func TestShortURL(t *testing.T) {
	body := []byte(`{"action":"long2short","long_url":"http://wap.koudaitong.com/v2/showcase/goods?alias=128wi9shh&spm=h56083&redirect_count=1"}`)

	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"short_url": "http:\/\/w.url.cn\/s\/AvCo6Ih"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/shorturl?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET", WithMockClient(client))

	result := new(ResultShortURL)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", ShortURL("http://wap.koudaitong.com/v2/showcase/goods?alias=128wi9shh&spm=h56083&redirect_count=1", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultShortURL{
		ShortURL: "http://w.url.cn/s/AvCo6Ih",
	}, result)
}
