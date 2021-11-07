package tools

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type WedriveAuthInfo struct {
	Type         int    `json:"type"`
	UserID       string `json:"userid,omitempty"`
	DepartmentID string `json:"departmentid,omitempty"`
	Auth         int    `json:"auth,omitempty"`
}

type ParamsWedriveSpaceCreate struct {
	UserID    string             `json:"userid"`
	SpaceName string             `json:"space_name"`
	AuthInfo  []*WedriveAuthInfo `json:"auth_info,omitempty"`
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
	SpaceInfo *WedriveSpaceInfo `json:"space_info"`
}

type WedriveSpaceInfo struct {
	SpaceID   string                `json:"spaceid"`
	SpaceName string                `json:"space_name"`
	AuthList  *WedriveSpaceAuthList `json:"auth_list"`
}

type WedriveSpaceAuthList struct {
	AuthInfo   []*WedriveAuthInfo `json:"auth_info"`
	QuitUserID []string           `json:"quit_userid"`
}

func GetWedriveSpaceInfo(params *ParamsWedriveSpaceInfo, result *ResultWedriveSpaceInfo) wx.Action {
	return wx.NewPostAction(urls.CorpToolsWedriveSpaceInfo,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsWedriveSpaceAclAdd struct {
	UserID   string             `json:"userid"`
	SpaceID  string             `json:"spaceid"`
	AuthInfo []*WedriveAuthInfo `json:"auth_info"`
}

func AddWedriveSpaceAcl(params *ParamsWedriveSpaceAclAdd) wx.Action {
	return wx.NewPostAction(urls.CorpToolsWedriveSpaceAclAdd,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsWedriveSpaceAclDel struct {
	UserID   string             `json:"userid"`
	SpaceID  string             `json:"spaceid"`
	AuthInfo []*WedriveAuthInfo `json:"auth_info"`
}

func DelWedriveSpaceAcl(params *ParamsWedriveSpaceAclAdd) wx.Action {
	return wx.NewPostAction(urls.CorpToolsWedriveSpaceAclAdd,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsWedriveSpaceSetting struct {
	UserID                       string `json:"userid"`
	SpaceID                      string `json:"spaceid"`
	EnableWatermark              *bool  `json:"enable_watermark,omitempty"`
	AddMemberOnlyAdmin           *bool  `json:"add_member_only_admin,omitempty"`
	EnableShareURL               *bool  `json:"enable_share_url,omitempty"`
	ShareURLNoApprove            *bool  `json:"share_url_no_approve,omitempty"`
	ShareURLNoApproveDefaultAuth int    `json:"share_url_no_approve_default_auth,omitempty"`
}

func SetWedriveSpace(params *ParamsWedriveSpaceSetting) wx.Action {
	return wx.NewPostAction(urls.CorpToolsWedriveSpaceSetting,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsWedriveSpaceShare struct {
	UserID  string `json:"userid"`
	SpaceID string `json:"spaceid"`
}

type ResultWedriveSpaceShare struct {
	SpaceShareURL string `json:"space_share_url"`
}

func ShareWedriveSpace(params *ParamsWedriveSpaceShare, result *ResultWedriveSpaceShare) wx.Action {
	return wx.NewPostAction(urls.CorpToolsWedriveSpaceShare,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type WedriveFileInfo struct {
	FileID       string `json:"fileid"`
	FileName     string `json:"file_name"`
	SpaceID      string `json:"spaceid"`
	FatherID     string `json:"fatherid"`
	FileSize     int64  `json:"file_size"`
	CTime        int64  `json:"ctime"`
	MTime        int64  `json:"mtime"`
	FileType     int    `json:"file_type"`
	FileStatus   int    `json:"file_status"`
	CreateUserID string `json:"create_userid"`
	UpdateUserID string `json:"update_userid"`
	SHA          string `json:"sha"`
	MD5          string `json:"md5"`
	URL          string `json:"url"`
}

type WedriveFileList struct {
	Item []*WedriveFileInfo `json:"item"`
}

type ParamsWedriveFileList struct {
	UserID   string `json:"userid"`
	SpaceID  string `json:"spaceid"`
	FatherID string `json:"fatherid"`
	SortType int    `json:"sort_type"`
	Start    int    `json:"start"`
	Limit    int    `json:"limit"`
}

type ResultWedriveFileList struct {
	HasMore   bool             `json:"has_more"`
	NextStart int64            `json:"next_start"`
	FileList  *WedriveFileList `json:"file_list"`
}

func GetWedriveFileList(params *ParamsWedriveFileList, result *ResultWedriveFileList) wx.Action {
	return wx.NewPostAction(urls.CorpToolsWedriveFileList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsWedriveFileUpload struct {
	UserID            string `json:"userid"`
	SpaceID           string `json:"spaceid"`
	FatherID          string `json:"fatherid"`
	FileName          string `json:"file_name"`
	FileBase64Content string `json:"file_base64_content"`
}

type ResultWedriveFileUpload struct {
	FileID string `json:"fileid"`
}

func UploadWedriveFile(params *ParamsWedriveFileUpload, result *ResultWedriveFileUpload) wx.Action {
	return wx.NewPostAction(urls.CorpToolsWedriveFileUpload,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsWedriveFileDownload struct {
	UserID string `json:"userid"`
	FileID string `json:"fileid"`
}

type ResultWedriveFileDownload struct {
	DownloadURL string `json:"download_url"`
	CookieName  string `json:"cookie_name"`
	CookieValue string `json:"cookie_value"`
}

func DownloadWedriveFile(params *ParamsWedriveFileDownload, result *ResultWedriveFileDownload) wx.Action {
	return wx.NewPostAction(urls.CorpToolsWedriveFileDownload,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsWedriveFileCreate struct {
	UserID   string `json:"userid"`
	SpaceID  string `json:"spaceid"`
	FatherID string `json:"fatherid"`
	FileType string `json:"file_type"`
	FileName string `json:"file_name"`
}

type ResultWedriveFileCreate struct {
	FileID string `json:"fileid"`
	URL    string `json:"url"`
}

func CreateWedriveFile(params *ParamsWedriveFileCreate, result *ResultWedriveFileCreate) wx.Action {
	return wx.NewPostAction(urls.CorpToolsWedriveFileCreate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsWedriveFileRename struct {
	UserID  string `json:"userid"`
	FileID  string `json:"fileid"`
	NewName string `json:"new_name"`
}

type ResultWedriveFileRename struct {
	File *WedriveFileInfo `json:"file"`
}

func RenameWedriveFile(params *ParamsWedriveFileRename, result *ResultWedriveFileRename) wx.Action {
	return wx.NewPostAction(urls.CorpToolsWedriveFileRename,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsWedriveFileMove struct {
	UserID   string   `json:"userid"`
	FatherID string   `json:"fatherid"`
	Replace  bool     `json:"replace"`
	FileID   []string `json:"fileid"`
}

type ResultWedriveFileMove struct {
	FileList *WedriveFileList `json:"file_list"`
}

func MoveWedriveFile(params *ParamsWedriveFileMove, result *ResultWedriveFileMove) wx.Action {
	return wx.NewPostAction(urls.CorpToolsWedriveFileMove,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsWedriveFileDelete struct {
	UserID string   `json:"userid"`
	FileID []string `json:"fileid"`
}

func DeleteWedriveFile(params *ParamsWedriveFileDelete) wx.Action {
	return wx.NewPostAction(urls.CorpToolsWedriveFileDelete,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsWedriveFileInfo struct {
	UserID string `json:"userid"`
	FileID string `json:"fileid"`
}

type ResultWedriveFileInfo struct {
	FileInfo *WedriveFileInfo `json:"file_info"`
}

func GetWedriveFileInfo(params *ParamsWedriveFileInfo, result *ResultWedriveFileInfo) wx.Action {
	return wx.NewPostAction(urls.CorpToolsWedriveFileInfo,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsWedriveFileAclAdd struct {
	UserID   string             `json:"userid"`
	FileID   string             `json:"fileid"`
	AuthInfo []*WedriveAuthInfo `json:"auth_info"`
}

func AddWedriveFileAcl(params *ParamsWedriveFileAclAdd) wx.Action {
	return wx.NewPostAction(urls.CorpToolsWedriveFileAclAdd,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsWedriveFileAclDel struct {
	UserID   string             `json:"userid"`
	FileID   string             `json:"fileid"`
	AuthInfo []*WedriveAuthInfo `json:"auth_info"`
}

func DelWedriveFileAcl(params *ParamsWedriveFileAclAdd) wx.Action {
	return wx.NewPostAction(urls.CorpToolsWedriveFileAclAdd,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsWedriveFileSetting struct {
	UserID    string `json:"userid"`
	FileID    string `json:"fileid"`
	AuthScope int    `json:"auth_scope"`
	Auth      int    `json:"auth"`
}

func SetWedriveFile(params *ParamsWedriveFileSetting) wx.Action {
	return wx.NewPostAction(urls.CorpToolsWedriveFileSetting,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsWedriveFileShare struct {
	UserID string `json:"userid"`
	FileID string `json:"fileid"`
}

type ResultWedriveFileShare struct {
	FileShareURL string `json:"file_share_url"`
}

func ShareWedriveFile(params *ParamsWedriveFileShare, result *ResultWedriveFileShare) wx.Action {
	return wx.NewPostAction(urls.CorpToolsWedriveFileShare,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
