package kf

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// InviteStatus 客服邀请状态
type InviteStatus string

// 微信支持的客服邀请状态
const (
	InviteWaiting  InviteStatus = "waiting"  // 待确认
	InviteRejected InviteStatus = "rejected" // 被拒绝
	InviteExpired  InviteStatus = "expired"  // 已过期
)

// Account 客服账号
type Account struct {
	KFID             string       `json:"kf_id"`              // 客服编号
	KFAccount        string       `json:"kf_account"`         // 完整客服帐号，格式为：帐号前缀@公众号微信号
	KFNick           string       `json:"kf_nick"`            // 客服昵称
	KFHeadImgURL     string       `json:"kf_headimgurl"`      // 客服头像
	KFWX             string       `json:"kf_wx"`              // 如果客服帐号已绑定了客服人员微信号， 则此处显示微信号
	InviteWeixin     string       `json:"invite_wx"`          // 如果客服帐号尚未绑定微信号，但是已经发起了一个绑定邀请， 则此处显示绑定邀请的微信号
	InviteExpireTime int64        `json:"invite_expire_time"` // 如果客服帐号尚未绑定微信号，但是已经发起过一个绑定邀请， 邀请的过期时间，为unix 时间戳
	InviteStatus     InviteStatus `json:"invite_status"`      // 邀请的状态，有等待确认“waiting”，被拒绝“rejected”， 过期“expired”
}

type ResultAccountList struct {
	KFList []*Account `json:"kf_list"`
}

// GetAccountList 获取客服列表
func GetAccountList(result *ResultAccountList) wx.Action {
	return wx.NewGetAction(urls.OffiaKFAccountList,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// Online 在线客服
type Online struct {
	KFID         string `json:"kf_id"`         // 客服编号
	KFAccount    string `json:"kf_account"`    // 完整客服帐号，格式为：帐号前缀@公众号微信号
	Status       int    `json:"status"`        // 客服在线状态，目前为：1-web在线
	AcceptedCase int    `json:"accepted_case"` // 客服当前正在接待的会话数
}

type ResultOnlineList struct {
	KFOnlineList []*Online `json:"kf_online_list"`
}

// GetOnlineList 获取客服在线列表
func GetOnlineList(result *ResultOnlineList) wx.Action {
	return wx.NewGetAction(urls.OffiaKFOnlineList,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsAccountAdd struct {
	KFAccount string `json:"kf_account"` // 完整客服帐号，格式为：帐号前缀@公众号微信号，帐号前缀最多10个字符，必须是英文、数字字符或者下划线，后缀为公众号微信号，长度不超过30个字符
	Nickname  string `json:"nickname"`   // 客服昵称，最长16个字
}

// AddAccount 添加客服账号
func AddAccount(account, nickname string) wx.Action {
	params := &ParamsAccountAdd{
		KFAccount: account,
		Nickname:  nickname,
	}

	return wx.NewPostAction(urls.OffiaKFAccountAdd,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ParamsAccountUpdate struct {
	KFAccount string `json:"kf_account"` // 完整客服帐号，格式为：帐号前缀@公众号微信号，帐号前缀最多10个字符，必须是英文、数字字符或者下划线，后缀为公众号微信号，长度不超过30个字符
	Nickname  string `json:"nickname"`   // 客服昵称，最长16个字
}

// UpdateAccount 设置客服信息
func UpdateAccount(account, nickname string) wx.Action {
	params := &ParamsAccountUpdate{
		KFAccount: account,
		Nickname:  nickname,
	}

	return wx.NewPostAction(urls.OffiaKFAccountUpdate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ParamsWorkerInvite struct {
	KFAccount string `json:"kf_account"` // 完整客服帐号，格式为：帐号前缀@公众号微信号
	InviteWX  string `json:"invite_wx"`  // 接收绑定邀请的客服微信号
}

// InviteWorker 邀请绑定客服帐号
// 新添加的客服帐号是不能直接使用的，只有客服人员用微信号绑定了客服账号后，方可登录Web客服进行操作。
// 发起一个绑定邀请到客服人员微信号，客服人员需要在微信客户端上用该微信号确认后帐号才可用。
// 尚未绑定微信号的帐号可以进行绑定邀请操作，邀请未失效时不能对该帐号进行再次绑定微信号邀请。
func InviteWorker(account, inviteWX string) wx.Action {
	params := &ParamsWorkerInvite{
		KFAccount: account,
		InviteWX:  inviteWX,
	}

	return wx.NewPostAction(urls.OffiaKFInvite,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

// UploadAvatar 上传客服头像（文件大小为5M以内）
func UploadAvatar(account, imgPath string) wx.Action {
	_, filename := filepath.Split(imgPath)

	return wx.NewPostAction(urls.OffiaKFAvatarUpload,
		wx.WithQuery("kf_account", account),
		wx.WithUpload(func() (wx.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(imgPath))

			if err != nil {
				return nil, err
			}

			return wx.NewUploadForm(
				wx.WithFormFile("media", filename, func(w io.Writer) error {
					f, err := os.Open(path)

					if err != nil {
						return err
					}

					defer f.Close()

					if _, err = io.Copy(w, f); err != nil {
						return err
					}

					return nil
				}),
			), nil
		}),
	)
}

// DeleteAccount 删除客服帐号
// 完整客服帐号，格式为：帐号前缀@公众号微信号
func DeleteAccount(account string) wx.Action {
	return wx.NewGetAction(urls.OffiaKFDelete,
		wx.WithQuery("kf_account", account),
	)
}
