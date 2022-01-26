package linkedcorp

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/corp"
	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/wx"
)

func TestGetUser(t *testing.T) {
	body := []byte(`{"userid":"CORPID/USERID"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "user_info": {
        "userid": "zhangsan",
        "name": "张三",
        "department": [
            "LINKEDID/1",
            "LINKEDID/2"
        ],
        "mobile": "+86 12345678901",
        "telephone": "10086",
        "email": "zhangsan@tencent.com",
        "position": "后台开发",
        "corpid": "xxxxxx",
        "extattr": {
            "attrs": [
                {
                    "name": "自定义属性(文本)",
                    "type": 0,
                    "text": {
                        "value": "10086"
                    }
                },
                {
                    "name": "自定义属性(网页)",
                    "type": 1,
                    "web": {
                        "url": "https://work.weixin.qq.com/",
                        "title": "官网"
                    }
                }
            ]
        }
    }
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/user/get?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultUserGet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetUser("CORPID", "USERID", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultUserGet{
		UserInfo: &UserInfo{
			UserID:     "zhangsan",
			Name:       "张三",
			Department: []string{"LINKEDID/1", "LINKEDID/2"},
			Mobile:     "+86 12345678901",
			Telephone:  "10086",
			EMail:      "zhangsan@tencent.com",
			Position:   "后台开发",
			CorpID:     "xxxxxx",
			ExtAttr: &ExtAttr{
				Attrs: []*Attr{
					{
						Name: "自定义属性(文本)",
						Type: 0,
						Text: &AttrText{
							Value: "10086",
						},
					},
					{
						Name: "自定义属性(网页)",
						Type: 1,
						Web: &AttrWeb{
							Title: "官网",
							URL:   "https://work.weixin.qq.com/",
						},
					},
				},
			},
		},
	}, result)
}

func TestListUserSimple(t *testing.T) {
	body := []byte(``)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/user/authsucc?access_token=ACCESS_TOKEN&userid=USERID", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN")

	assert.Nil(t, err)
}

func TestListUser(t *testing.T) {
	body := []byte(``)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/user/authsucc?access_token=ACCESS_TOKEN&userid=USERID", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN")

	assert.Nil(t, err)
}
