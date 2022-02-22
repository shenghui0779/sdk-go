package externalcontact

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ProductAlbumImage struct {
	MediaID string `json:"media_id,omitempty"`
}

type ProductAlbumAttachment struct {
	Type  MediaType          `json:"type"` // 附件类型，目前仅支持image
	Image *ProductAlbumImage `json:"image"`
}

type ParamsProductAlbumAdd struct {
	Description string                    `json:"description"`
	Price       int                       `json:"price"`
	ProductSN   string                    `json:"product_sn,omitempty"`
	Attachments []*ProductAlbumAttachment `json:"attachments"`
}

type ResultProductAlbumAdd struct {
	ProductID string `json:"product_id"`
}

// AddProductAlbum 创建商品图册
func AddProductAlbum(params *ParamsProductAlbumAdd, result *ResultProductAlbumAdd) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactProductAlbumAdd,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsProductAlbumUpdate struct {
	ProductID   string                    `json:"product_id"`
	Description string                    `json:"description"`
	Price       int                       `json:"price"`
	ProductSN   string                    `json:"product_sn,omitempty"`
	Attachments []*ProductAlbumAttachment `json:"attachments,omitempty"`
}

// UpdateProductAlbum 编辑商品图册
func UpdateProductAlbum(params *ParamsProductAlbumUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactProductAlbumUpdate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ProductAlbum struct {
	ProductID   string                    `json:"product_id"`
	Description string                    `json:"description"`
	Price       int                       `json:"price"`
	CreateTime  int64                     `json:"create_time"`
	ProductSN   string                    `json:"product_sn"`
	Attachments []*ProductAlbumAttachment `json:"attachments"`
}

type ParamsProductAlbumGet struct {
	ProductID string `json:"product_id"`
}

type ResultProductAlbumGet struct {
	Product *ProductAlbum `json:"product"`
}

// GetProductAlbum 获取商品图册
func GetProductAlbum(productID string, result *ResultProductAlbumGet) wx.Action {
	params := &ParamsProductAlbumGet{
		ProductID: productID,
	}

	return wx.NewPostAction(urls.CorpExternalContactProductAlbumGet,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsProductAlbumList struct {
	Limit  int    `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}

type ResultProductAlbumList struct {
	NextCursor  string          `json:"next_cursor"`
	ProductList []*ProductAlbum `json:"product_list"`
}

// ListProductAlbum 获取商品图册列表
func ListProductAlbum(cursor string, limit int, result *ResultProductAlbumList) wx.Action {
	params := &ParamsProductAlbumList{
		Limit:  limit,
		Cursor: cursor,
	}

	return wx.NewPostAction(urls.CorpExternalContactProductAlbumList,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsProductAlbumDelete struct {
	ProductID string `json:"product_id"`
}

// DeleteProductAlbum 删除商品图册
func DeleteProductAlbum(productID string) wx.Action {
	params := &ParamsProductAlbumDelete{
		ProductID: productID,
	}

	return wx.NewPostAction(urls.CorpExternalContactProductAlbumDelete,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}
