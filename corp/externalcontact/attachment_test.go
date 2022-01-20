package externalcontact

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/yiigo"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/corp"
	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/wx"
)

func TestUploadAttachment(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"type": "image",
	"media_id": "1G6nrLmr5EC3MMb_-zK1dDdzmd0p7cNliYu9V5w7o8K0",
	"created_at": 1380000000
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://qyapi.weixin.qq.com/cgi-bin/media/upload_attachment?access_token=accesstoken001&media_type=TYPE&attachment_type=1", gomock.AssignableToTypeOf(yiigo.NewUploadForm())).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsAttachmentUpload{
		MediaType:      MediaImage,
		AttachmentType: AttachmentMoment,
		Path:           "../mock/test.jpg",
	}
	result := new(ResultAttachmentUpload)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UploadAttachment(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultAttachmentUpload{
		Type:      MediaImage,
		MediaID:   "1G6nrLmr5EC3MMb_-zK1dDdzmd0p7cNliYu9V5w7o8K0",
		CreatedAt: 1380000000,
	}, result)
}

func TestUploadAttachmentByURL(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"type": "image",
	"media_id": "1G6nrLmr5EC3MMb_-zK1dDdzmd0p7cNliYu9V5w7o8K0",
	"created_at": 1380000000
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://qyapi.weixin.qq.com/cgi-bin/media/upload_attachment?access_token=accesstoken001&media_type=TYPE&attachment_type=1", gomock.AssignableToTypeOf(yiigo.NewUploadForm())).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsAttachmentUploadByURL{
		MediaType:      MediaImage,
		AttachmentType: AttachmentMoment,
		Filename:       "test.png",
		URL:            "https://golang.google.cn/doc/gopher/pkg.png",
	}
	result := new(ResultAttachmentUpload)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UploadAttachmentByURL(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultAttachmentUpload{
		Type:      MediaImage,
		MediaID:   "1G6nrLmr5EC3MMb_-zK1dDdzmd0p7cNliYu9V5w7o8K0",
		CreatedAt: 1380000000,
	}, result)
}
