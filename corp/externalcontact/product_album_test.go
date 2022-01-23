package externalcontact

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

func TestAddProductAlbum(t *testing.T) {
	body := []byte(`{"description":"世界上最好的商品","price":30000,"product_sn":"xxxxxxxx","attachments":[{"type":"image","image":{"media_id":"MEDIA_ID"}}]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "product_id": "xxxxxxxxxx"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_product_album?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsProductAlbumAdd{
		Description: "世界上最好的商品",
		Price:       30000,
		ProductSN:   "xxxxxxxx",
		Attachments: []*ProductAlbumAttachment{
			{
				Type: MediaImage,
				Image: &ProductAlbumImage{
					MediaID: "MEDIA_ID",
				},
			},
		},
	}

	result := new(ResultProductAlbumAdd)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", AddProductAlbum(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultProductAlbumAdd{
		ProductID: "xxxxxxxxxx",
	}, result)
}

func TestUpdateProductAlbum(t *testing.T) {
	body := []byte(`{"product_id":"xxxxxxxxxx","description":"世界上最好的商品","price":30000,"product_sn":"xxxxxx","attachments":[{"type":"image","image":{"media_id":"MEDIA_ID"}}]}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/update_product_album?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsProductAlbumUpdate{
		ProductID:   "xxxxxxxxxx",
		Description: "世界上最好的商品",
		Price:       30000,
		ProductSN:   "xxxxxx",
		Attachments: []*ProductAlbumAttachment{
			{
				Type: MediaImage,
				Image: &ProductAlbumImage{
					MediaID: "MEDIA_ID",
				},
			},
		},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UpdateProductAlbum(params))

	assert.Nil(t, err)
}

func TestGetProductAlbum(t *testing.T) {
	body := []byte(`{"product_id":"xxxxxxxxxx"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "product": {
        "product_id": "xxxxxxxxxx",
        "description": "世界上最好的商品",
        "price": 30000,
        "create_time": 1600000000,
        "product_sn": "xxxxxxxx",
        "attachments": [
            {
                "type": "image",
                "image": {
                    "media_id": "MEDIA_ID"
                }
            }
        ]
    }
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_product_album?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultProductAlbumGet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetProductAlbum("xxxxxxxxxx", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultProductAlbumGet{
		Product: &ProductAlbum{
			ProductID:   "xxxxxxxxxx",
			Description: "世界上最好的商品",
			Price:       30000,
			CreateTime:  1600000000,
			ProductSN:   "xxxxxxxx",
			Attachments: []*ProductAlbumAttachment{
				{
					Type: MediaImage,
					Image: &ProductAlbumImage{
						MediaID: "MEDIA_ID",
					},
				},
			},
		},
	}, result)
}

func TestListProductAlbum(t *testing.T) {
	body := []byte(`{"limit":50,"cursor":"CURSOR"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "next_cursor": "CURSOR",
    "product_list": [
        {
            "product_id": "xxxxxxxxxx",
            "description": "世界上最好的商品",
            "price": 30000,
            "product_sn": "xxxxxxxx",
            "attachments": [
                {
                    "type": "image",
                    "image": {
                        "media_id": "MEDIA_ID"
                    }
                }
            ]
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_product_album_list?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsProductAlbumList{
		Limit:  50,
		Cursor: "CURSOR",
	}

	result := new(ResultProductAlbumList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListProductAlbum(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultProductAlbumList{
		NextCursor: "CURSOR",
		ProductList: []*ProductAlbum{
			{
				ProductID:   "xxxxxxxxxx",
				Description: "世界上最好的商品",
				Price:       30000,
				ProductSN:   "xxxxxxxx",
				Attachments: []*ProductAlbumAttachment{
					{
						Type: MediaImage,
						Image: &ProductAlbumImage{
							MediaID: "MEDIA_ID",
						},
					},
				},
			},
		},
	}, result)
}

func TestDeleteProductAlbum(t *testing.T) {
	body := []byte(`{"product_id":"xxxxxxxxxx"}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/delete_product_album?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteProductAlbum("xxxxxxxxxx"))

	assert.Nil(t, err)
}
