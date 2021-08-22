package urls

const (
	CorpCgiBinAccessToken  = "https://qyapi.weixin.qq.com/cgi-bin/gettoken"
	CorpCgiBinAPIDomainIP  = "https://qyapi.weixin.qq.com/cgi-bin/get_api_domain_ip"
	CorpCgiBinUserInfo     = "https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo"
	CorpCgiBinUserAuthSucc = "https://qyapi.weixin.qq.com/cgi-bin/user/authsucc"
)

// addr_book
const (
	CorpAddrBookUserCreate        = "https://qyapi.weixin.qq.com/cgi-bin/user/create"
	CorpAddrBookUserGet           = "https://qyapi.weixin.qq.com/cgi-bin/user/get"
	CorpAddrBookUserUpdate        = "https://qyapi.weixin.qq.com/cgi-bin/user/update"
	CorpAddrBookUserDelete        = "https://qyapi.weixin.qq.com/cgi-bin/user/delete"
	CorpAddrBookUserBatchDelete   = "https://qyapi.weixin.qq.com/cgi-bin/user/batchdelete"
	CorpAddrBookUserSimpleList    = "https://qyapi.weixin.qq.com/cgi-bin/user/simplelist"
	CorpAddrBookUserList          = "https://qyapi.weixin.qq.com/cgi-bin/user/list"
	CorpAddrBookConvert2OpenID    = "https://qyapi.weixin.qq.com/cgi-bin/user/convert_to_openid"
	CorpAddrBookConvert2UserID    = "https://qyapi.weixin.qq.com/cgi-bin/user/convert_to_userid"
	CorpAddrBookBatchInvite       = "https://qyapi.weixin.qq.com/cgi-bin/batch/invite"
	CorpAddrBookJoinQRCode        = "https://qyapi.weixin.qq.com/cgi-bin/corp/get_join_qrcode"
	CorpAddrBookActiveStat        = "https://qyapi.weixin.qq.com/cgi-bin/user/get_active_stat"
	CorpAddrBookDepartmentCreate  = "https://qyapi.weixin.qq.com/cgi-bin/department/create"
	CorpAddrBookDepartmentUpdate  = "https://qyapi.weixin.qq.com/cgi-bin/department/update"
	CorpAddrBookDepartmentDelete  = "https://qyapi.weixin.qq.com/cgi-bin/department/delete"
	CorpAddrBookDepartmentList    = "https://qyapi.weixin.qq.com/cgi-bin/department/list"
	CorpAddrBookTagCreate         = "https://qyapi.weixin.qq.com/cgi-bin/tag/create"
	CorpAddrBookTagUpdate         = "https://qyapi.weixin.qq.com/cgi-bin/tag/update"
	CorpAddrBookTagDelete         = "https://qyapi.weixin.qq.com/cgi-bin/tag/delete"
	CorpAddrBookTagGet            = "https://qyapi.weixin.qq.com/cgi-bin/tag/get"
	CorpAddrBookTagList           = "https://qyapi.weixin.qq.com/cgi-bin/tag/list"
	CorpAddrBookTagUserAdd        = "https://qyapi.weixin.qq.com/cgi-bin/tag/addtagusers"
	CorpAddrBookTagUserDel        = "https://qyapi.weixin.qq.com/cgi-bin/tag/deltagusers"
	CorpAddrBookBatchSyncUser     = "https://qyapi.weixin.qq.com/cgi-bin/batch/syncuser"
	CorpAddrBookBatchReplaceUser  = "https://qyapi.weixin.qq.com/cgi-bin/batch/replaceuser"
	CorpAddrBookBatchReplaceParty = "https://qyapi.weixin.qq.com/cgi-bin/batch/replaceparty"
	CorpAddrBookBatchResult       = "https://qyapi.weixin.qq.com/cgi-bin/batch/getresult"
	CorpLinkedCorpPermList        = "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/agent/get_perm_list"
	CorpLinkedCorpUserGet         = "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/user/get"
	CorpLinkedCorpUserSimpleList  = "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/user/simplelist"
	CorpLinkedCorpUserList        = "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/user/list"
	CorpLinkedCorpDepartmentList  = "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/department/list"
)

// external_contact
const (
	CorpExternalContactFollowUserList = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_follow_user_list"
	CorpExternalContactWayAdd         = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_contact_way"
	CorpExternalContactWayGet         = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_contact_way"
	CorpExternalContactWayUpdate      = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/update_contact_way"
	CorpExternalContactWayDelete      = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_contact_way"
	CorpExternalContactCloseTempChat  = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/close_temp_chat"
	CorpExternalContactList           = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list"
)

// kf
const (
	CorpKFAccountAdd              = "https://qyapi.weixin.qq.com/cgi-bin/kf/account/add"
	CorpKFAccountDel              = "https://qyapi.weixin.qq.com/cgi-bin/kf/account/del"
	CorpKFAccountUpdate           = "https://qyapi.weixin.qq.com/cgi-bin/kf/account/update"
	CorpKFAccountList             = "https://qyapi.weixin.qq.com/cgi-bin/kf/account/list"
	CorpKFAddContactWay           = "https://qyapi.weixin.qq.com/cgi-bin/kf/add_contact_way"
	CorpKFServicerAdd             = "https://qyapi.weixin.qq.com/cgi-bin/kf/servicer/add"
	CorpKFServicerDel             = "https://qyapi.weixin.qq.com/cgi-bin/kf/servicer/del"
	CorpKFServicerList            = "https://qyapi.weixin.qq.com/cgi-bin/kf/servicer/list"
	CorpKFSyncMsg                 = "https://qyapi.weixin.qq.com/cgi-bin/kf/sync_msg"
	CorpKFSendMsg                 = "https://qyapi.weixin.qq.com/cgi-bin/kf/send_msg"
	CorpKFCustomerBatchGet        = "https://qyapi.weixin.qq.com/cgi-bin/kf/customer/batchget"
	CorpKFGetUpgradeServiceConfig = "https://qyapi.weixin.qq.com/cgi-bin/kf/customer/get_upgrade_service_config"
	CorpKFUpgradeService          = "https://qyapi.weixin.qq.com/cgi-bin/kf/customer/upgrade_service"
	CorpKFCancelUpgradeService    = "https://qyapi.weixin.qq.com/cgi-bin/kf/customer/cancel_upgrade_service"
)
