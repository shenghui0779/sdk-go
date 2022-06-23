package media

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/chenghonour/gochat/corp"
	"github.com/chenghonour/gochat/mock"
	"github.com/chenghonour/gochat/wx"
)

func TestUpload(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"type": "image",
	"media_id": "1G6nrLmr5EC3MMb_-zK1dDdzmd0p7cNliYu9V5w7o8K0",
	"created_at": "1380000000"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://qyapi.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=image", gomock.AssignableToTypeOf(wx.NewUploadForm())).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultUpload)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", Upload(MediaImage, "../../mock/test.jpg", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultUpload{
		Type:      MediaImage,
		MediaID:   "1G6nrLmr5EC3MMb_-zK1dDdzmd0p7cNliYu9V5w7o8K0",
		CreatedAt: "1380000000",
	}, result)
}

func TestUploadByURL(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"type": "image",
	"media_id": "1G6nrLmr5EC3MMb_-zK1dDdzmd0p7cNliYu9V5w7o8K0",
	"created_at": "1380000000"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://qyapi.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=image", gomock.AssignableToTypeOf(wx.NewUploadForm())).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultUpload)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UploadByURL(MediaImage, "test.png", "https://golang.google.cn/doc/gopher/pkg.png", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultUpload{
		Type:      MediaImage,
		MediaID:   "1G6nrLmr5EC3MMb_-zK1dDdzmd0p7cNliYu9V5w7o8K0",
		CreatedAt: "1380000000",
	}, result)
}

func TestUploadByByte(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"type": "image",
	"media_id": "1G6nrLmr5EC3MMb_-zK1dDdzmd0p7cNliYu9V5w7o8K0",
	"created_at": "1380000000"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://qyapi.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=file", gomock.AssignableToTypeOf(wx.NewUploadForm())).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultUpload)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UploadByByte(MediaFile, "test.txt", []byte("test123"), result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultUpload{
		Type:      MediaImage,
		MediaID:   "1G6nrLmr5EC3MMb_-zK1dDdzmd0p7cNliYu9V5w7o8K0",
		CreatedAt: "1380000000",
	}, result)
}

func TestUploadImg(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"url": "http://p.qpic.cn/pic_wework/3474110808/7a7c8471673ff0f178f63447935d35a5c1247a7f31d9c060/0"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://qyapi.weixin.qq.com/cgi-bin/media/uploadimg?access_token=ACCESS_TOKEN", gomock.AssignableToTypeOf(wx.NewUploadForm())).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultUploadImg)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UploadImg("../../mock/test.jpg", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultUploadImg{
		URL: "http://p.qpic.cn/pic_wework/3474110808/7a7c8471673ff0f178f63447935d35a5c1247a7f31d9c060/0",
	}, result)
}

func TestUploadImgByURL(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"url": "http://p.qpic.cn/pic_wework/3474110808/7a7c8471673ff0f178f63447935d35a5c1247a7f31d9c060/0"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://qyapi.weixin.qq.com/cgi-bin/media/uploadimg?access_token=ACCESS_TOKEN", gomock.AssignableToTypeOf(wx.NewUploadForm())).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultUploadImg)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UploadImgByURL("test.png", "https://golang.google.cn/doc/gopher/pkg.png", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultUploadImg{
		URL: "http://p.qpic.cn/pic_wework/3474110808/7a7c8471673ff0f178f63447935d35a5c1247a7f31d9c060/0",
	}, result)
}
