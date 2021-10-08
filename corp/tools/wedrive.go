package tools

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type SpaceAuthInfo struct {
	Type   int    `json:"type"`
	UserID string `json:"userid"`
	Auth   int    `json:"auth"`
}

type ParamsWedriveSpaceCreate struct {
	UserID    string           `json:"userid"`
	SpaceName string           `json:"space_name"`
	AuthInfo  []*SpaceAuthInfo `json:"auth_info"`
}

type ResultWedriveSpaceCreate struct {
	SpaceID string `json:"spaceid"`
}

func CreateWedriveSpace(params *ParamsWedriveSpaceCreate, result *ResultWedriveSpaceCreate) wx.Action {
	return wx.NewPostAction(urls.CorpToolsWedriveSpaceCreate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsWedriveSpaceRename struct {
	UserID    string `json:"userid"`
	SpaceID   string `json:"spaceid"`
	SpaceName string `json:"space_name"`
}

func RenameWedriveSpace(params *ParamsWedriveSpaceRename) wx.Action {
	return wx.NewPostAction(urls.CorpToolsWedriveSpaceRename,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsWedriveSpaceDismiss struct {
	UserID  string `json:"userid"`
	SpaceID string `json:"spaceid"`
}

func DismissWedriveSpace(params *ParamsWedriveSpaceRename) wx.Action {
	return wx.NewPostAction(urls.CorpToolsWedriveSpaceDismiss,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsWedriveSpaceInfo struct {
	UserID  string `json:"userid"`
	SpaceID string `json:"spaceid"`
}

type ResultWedriveSpaceInfo struct {
	SpaceInfo *SpaceInfo `json:"space_info"`
}

type SpaceInfo struct {
	SpaceID   string         `json:"spaceid"`
	SpaceName string         `json:"space_name"`
	AuthList  *SpaceAuthList `json:"auth_list"`
}

type SpaceAuthList struct {
	AuthInfo   []*SpaceAuthInfo `json:"auth_info"`
	QuitUserID []string         `json:"quit_userid"`
}

func WedriveSpaceInfo(params *ParamsWedriveSpaceInfo, result *ResultWedriveSpaceInfo) wx.Action {
	return wx.NewPostAction(urls.CorpToolsWedriveSpaceInfo,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
