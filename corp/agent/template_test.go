package agent

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/corp"
	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestSetWorkbenchKeyDataTemplate(t *testing.T) {
	body := []byte(`{"agentid":1000005,"type":"keydata","keydata":{"items":[{"key":"待审批","data":"2","jump_url":"http://www.qq.com","pagepath":"pages/index"},{"key":"带批阅作业","data":"4","jump_url":"http://www.qq.com","pagepath":"pages/index"},{"key":"成绩录入","data":"45","jump_url":"http://www.qq.com","pagepath":"pages/index"},{"key":"综合评价","data":"98","jump_url":"http://www.qq.com","pagepath":"pages/index"}]},"replace_user_data":true}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/agent/set_workbench_template?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	keydata := &WorkbenchKeyData{
		Items: []*KeyDataItem{
			{
				Key:      "待审批",
				Data:     "2",
				JumpURL:  "http://www.qq.com",
				PagePath: "pages/index",
			},
			{
				Key:      "带批阅作业",
				Data:     "4",
				JumpURL:  "http://www.qq.com",
				PagePath: "pages/index",
			},
			{
				Key:      "成绩录入",
				Data:     "45",
				JumpURL:  "http://www.qq.com",
				PagePath: "pages/index",
			},
			{
				Key:      "综合评价",
				Data:     "98",
				JumpURL:  "http://www.qq.com",
				PagePath: "pages/index",
			},
		},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SetWorkbenchKeyDataTemplate(1000005, keydata, true))

	assert.Nil(t, err)
}

func TestSetWorkbenchImageTemplate(t *testing.T) {
	body := []byte(`{"agentid":1000005,"type":"image","image":{"url":"xxxx","jump_url":"http://www.qq.com","pagepath":"pages/index"},"replace_user_data":true}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/agent/set_workbench_template?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	image := &WorkbenchImage{
		URL:      "xxxx",
		JumpURL:  "http://www.qq.com",
		PagePath: "pages/index",
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SetWorkbenchImageTemplate(1000005, image, true))

	assert.Nil(t, err)
}

func TestSetWorkbenchListTemplate(t *testing.T) {
	body := []byte(`{"agentid":1000005,"type":"list","list":{"items":[{"title":"智慧校园新版设计的小程序要来啦","jump_url":"http://www.qq.com","pagepath":"pages/index"},{"title":"植物百科，这是什么花","jump_url":"http://www.qq.com","pagepath":"pages/index"},{"title":"周一升旗通知，全体学生必须穿校服","jump_url":"http://www.qq.com","pagepath":"pages/index"}]},"replace_user_data":true}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/agent/set_workbench_template?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	list := &WorkbenchList{
		Items: []*ListItem{
			{
				Title:    "智慧校园新版设计的小程序要来啦",
				JumpURL:  "http://www.qq.com",
				PagePath: "pages/index",
			},
			{
				Title:    "植物百科，这是什么花",
				JumpURL:  "http://www.qq.com",
				PagePath: "pages/index",
			},
			{
				Title:    "周一升旗通知，全体学生必须穿校服",
				JumpURL:  "http://www.qq.com",
				PagePath: "pages/index",
			},
		},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SetWorkbenchListTemplate(1000005, list, true))

	assert.Nil(t, err)
}

func TestSetWorkbenchWebViewTemplate(t *testing.T) {
	body := []byte(`{"agentid":1000005,"type":"webview","webview":{"url":"http://www.qq.com","jump_url":"http://www.qq.com","pagepath":"pages/index"},"replace_user_data":true}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/agent/set_workbench_template?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	webview := &WorkbenchWebView{
		URL:      "http://www.qq.com",
		JumpURL:  "http://www.qq.com",
		PagePath: "pages/index",
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SetWorkbenchWebViewTemplate(1000005, webview, true))

	assert.Nil(t, err)
}

func TestGetWorkbenchTemplate(t *testing.T) {
	body := []byte(`{"agentid":1000005}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"type": "image",
	"image": {
		"url": "xxxx",
		"jump_url": "http://www.qq.com",
		"pagepath": "pages/index"
	},
	"replace_user_data": true
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/agent/get_workbench_template?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultWorkbenchTemplateGet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetWorkbenchTemplate(1000005, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultWorkbenchTemplateGet{
		Type: TplImage,
		Image: &WorkbenchImage{
			URL:      "xxxx",
			JumpURL:  "http://www.qq.com",
			PagePath: "pages/index",
		},
		ReplaceUserData: true,
	}, result)
}

func TestSetWorkbenchKeyData(t *testing.T) {
	body := []byte(`{"agentid":1000005,"userid":"test","type":"keydata","keydata":{"items":[{"key":"待审批","data":"2","jump_url":"http://www.qq.com","pagepath":"pages/index"},{"key":"带批阅作业","data":"4","jump_url":"http://www.qq.com","pagepath":"pages/index"},{"key":"成绩录入","data":"45","jump_url":"http://www.qq.com","pagepath":"pages/index"},{"key":"综合评价","data":"98","jump_url":"http://www.qq.com","pagepath":"pages/index"}]}}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/agent/set_workbench_data?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	keydata := &WorkbenchKeyData{
		Items: []*KeyDataItem{
			{
				Key:      "待审批",
				Data:     "2",
				JumpURL:  "http://www.qq.com",
				PagePath: "pages/index",
			},
			{
				Key:      "带批阅作业",
				Data:     "4",
				JumpURL:  "http://www.qq.com",
				PagePath: "pages/index",
			},
			{
				Key:      "成绩录入",
				Data:     "45",
				JumpURL:  "http://www.qq.com",
				PagePath: "pages/index",
			},
			{
				Key:      "综合评价",
				Data:     "98",
				JumpURL:  "http://www.qq.com",
				PagePath: "pages/index",
			},
		},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SetWorkbenchKeyData(1000005, "test", keydata))

	assert.Nil(t, err)
}

func TestSetWorkbenchImageData(t *testing.T) {
	body := []byte(`{"agentid":1000005,"userid":"test","type":"image","image":{"url":"xxxx","jump_url":"http://www.qq.com","pagepath":"pages/index"}}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/agent/set_workbench_data?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	image := &WorkbenchImage{
		URL:      "xxxx",
		JumpURL:  "http://www.qq.com",
		PagePath: "pages/index",
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SetWorkbenchImageData(1000005, "test", image))

	assert.Nil(t, err)
}

func TestSetWorkbenchListData(t *testing.T) {
	body := []byte(`{"agentid":1000005,"userid":"test","type":"list","list":{"items":[{"title":"智慧校园新版设计的小程序要来啦","jump_url":"http://www.qq.com","pagepath":"pages/index"},{"title":"植物百科，这是什么花","jump_url":"http://www.qq.com","pagepath":"pages/index"},{"title":"周一升旗通知，全体学生必须穿校服","jump_url":"http://www.qq.com","pagepath":"pages/index"}]}}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/agent/set_workbench_data?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	list := &WorkbenchList{
		Items: []*ListItem{
			{
				Title:    "智慧校园新版设计的小程序要来啦",
				JumpURL:  "http://www.qq.com",
				PagePath: "pages/index",
			},
			{
				Title:    "植物百科，这是什么花",
				JumpURL:  "http://www.qq.com",
				PagePath: "pages/index",
			},
			{
				Title:    "周一升旗通知，全体学生必须穿校服",
				JumpURL:  "http://www.qq.com",
				PagePath: "pages/index",
			},
		},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SetWorkbenchListData(1000005, "test", list))

	assert.Nil(t, err)
}

func TestSetWorkbenchWebViewData(t *testing.T) {
	body := []byte(`{"agentid":1000005,"userid":"test","type":"webview","webview":{"url":"http://www.qq.com","jump_url":"http://www.qq.com","pagepath":"pages/index"}}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/agent/set_workbench_data?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	webview := &WorkbenchWebView{
		URL:      "http://www.qq.com",
		JumpURL:  "http://www.qq.com",
		PagePath: "pages/index",
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SetWorkbenchWebViewData(1000005, "test", webview))

	assert.Nil(t, err)
}
