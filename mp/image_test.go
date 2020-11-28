package mp

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/internal"
	"github.com/stretchr/testify/assert"
)

func TestAICrop(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := internal.NewMockWechatClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cv/img/aicrop?access_token=ACCESS_TOKEN", gomock.AssignableToTypeOf(postBody)).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"results": [
			{
				"crop_left": 112,
				"crop_top": 0,
				"crop_right": 839,
				"crop_bottom": 727
			},
			{
				"crop_left": 0,
				"crop_top": 205,
				"crop_right": 965,
				"crop_bottom": 615
			}
		],
		"img_size": {
			"w": 966,
			"h": 728
		}
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(AICropResult)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", AICrop("test.jpg", dest))

	assert.Nil(t, err)
	assert.Equal(t, &AICropResult{
		Results: []CropPosition{
			{
				CropLeft:   112,
				CropTop:    0,
				CropRight:  839,
				CropBottom: 727,
			},
			{
				CropLeft:   0,
				CropTop:    205,
				CropRight:  965,
				CropBottom: 615,
			},
		},
		ImgSize: ImageSize{
			W: 966,
			H: 728,
		},
	}, dest)
}
