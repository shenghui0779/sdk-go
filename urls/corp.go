package urls

const (
	CorpCgiBinAccessToken  = "https://qyapi.weixin.qq.com/cgi-bin/gettoken"
	CorpCgiBinAPIDomainIP  = "https://qyapi.weixin.qq.com/cgi-bin/get_api_domain_ip"
	CorpCgiBinUserInfo     = "https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo"
	CorpCgiBinUserAuthSucc = "https://qyapi.weixin.qq.com/cgi-bin/user/authsucc"
	CorpCginBinCallbackIP  = "https://qyapi.weixin.qq.com/cgi-bin/getcallbackip"
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
	CorpLinkedcorpPermList        = "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/agent/get_perm_list"
	CorpLinkedcorpUserGet         = "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/user/get"
	CorpLinkedcorpUserSimpleList  = "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/user/simplelist"
	CorpLinkedcorpUserList        = "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/user/list"
	CorpLinkedcorpDepartmentList  = "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/department/list"
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

// agent
const (
	CorpAgentGet             = "https://qyapi.weixin.qq.com/cgi-bin/agent/get"
	CorpAgentList            = "https://qyapi.weixin.qq.com/cgi-bin/agent/list"
	CorpAgentSet             = "https://qyapi.weixin.qq.com/cgi-bin/agent/set"
	CorpMenuCreate           = "https://qyapi.weixin.qq.com/cgi-bin/menu/create"
	CorpMenuGet              = "https://qyapi.weixin.qq.com/cgi-bin/menu/get"
	CorpMenuDelete           = "https://qyapi.weixin.qq.com/cgi-bin/menu/delete"
	CorpSetWorkbenchTemplate = "https://qyapi.weixin.qq.com/cgi-bin/agent/set_workbench_template"
	CorpGetWorkbenchTemplate = "https://qyapi.weixin.qq.com/cgi-bin/agent/get_workbench_template"
	CorpSetWorkbenchData     = "https://qyapi.weixin.qq.com/cgi-bin/agent/set_workbench_data"
)

// message
const (
	CorpMessageSend                = "https://qyapi.weixin.qq.com/cgi-bin/message/send"
	CorpMessageUpdateTemplateCard  = "https://qyapi.weixin.qq.com/cgi-bin/message/update_template_card"
	CorpMessageRecall              = "https://qyapi.weixin.qq.com/cgi-bin/message/recall"
	CorpAppchatCreate              = "https://qyapi.weixin.qq.com/cgi-bin/appchat/create"
	CorpAppchatUpdate              = "https://qyapi.weixin.qq.com/cgi-bin/appchat/update"
	CorpAppchatGet                 = "https://qyapi.weixin.qq.com/cgi-bin/appchat/get"
	CorpAppchatSend                = "https://qyapi.weixin.qq.com/cgi-bin/appchat/send"
	CorpLinkedcorpMessageSend      = "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/message/send"
	CorpExternalContactMessageSend = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/message/send"
	CorpMessageStaticsGet          = "https://qyapi.weixin.qq.com/cgi-bin/message/get_statistics"
)

// media
const (
	CorpMediaUpload    = "https://qyapi.weixin.qq.com/cgi-bin/media/upload"
	CorpMediaUploadimg = "https://qyapi.weixin.qq.com/cgi-bin/media/uploadimg"
	CorpMediaGet       = "https://qyapi.weixin.qq.com/cgi-bin/media/get"
	CorpMediaGetJSSDK  = "https://qyapi.weixin.qq.com/cgi-bin/media/get/jssdk"
)

// OA
const (
	CorpOAGetCorpCheckinOption      = "https://qyapi.weixin.qq.com/cgi-bin/checkin/getcorpcheckinoption"
	CorpOAGetCheckinOption          = "https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckinoption"
	CorpOAGetCheckinData            = "https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckindata"
	CorpOAGetCheckinDayData         = "https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckin_daydata"
	CorpOAGetCheckinMonthData       = "https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckin_monthdata"
	CorpOAGetCheckinScheduleList    = "https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckinschedulist"
	CorpOASetCheckinScheduleList    = "https://qyapi.weixin.qq.com/cgi-bin/checkin/setcheckinschedulist"
	CorpOAAddCheckinUserFace        = "https://qyapi.weixin.qq.com/cgi-bin/checkin/addcheckinuserface"
	CorpOAGetTemplateDetail         = "https://qyapi.weixin.qq.com/cgi-bin/oa/gettemplatedetail"
	CorpOAApplyEvent                = "https://qyapi.weixin.qq.com/cgi-bin/oa/applyevent"
	CorpOAGetApprovalInfo           = "https://qyapi.weixin.qq.com/cgi-bin/oa/getapprovalinfo"
	CorpOAGetApprovalDetail         = "https://qyapi.weixin.qq.com/cgi-bin/oa/getapprovaldetail"
	CorpOAGetVacationCorpConf       = "https://qyapi.weixin.qq.com/cgi-bin/oa/vacation/getcorpconf"
	CorpOAGetUserVacationQuota      = "https://qyapi.weixin.qq.com/cgi-bin/oa/vacation/getuservacationquota"
	CorpOASetUserVacationQuota      = "https://qyapi.weixin.qq.com/cgi-bin/oa/vacation/setoneuserquota"
	CorpOAGetJournalRecordList      = "https://qyapi.weixin.qq.com/cgi-bin/oa/journal/get_record_list"
	CorpOAGetJournalRecordDetail    = "https://qyapi.weixin.qq.com/cgi-bin/oa/journal/get_record_detail"
	CorpOAGetJournalStatList        = "https://qyapi.weixin.qq.com/cgi-bin/oa/journal/get_stat_list"
	CorpOAOpenApprovalDataGet       = "https://qyapi.weixin.qq.com/cgi-bin/corp/getopenapprovaldata"
	CorpOAMeetingRoomAdd            = "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/add"
	CorpOAMeetingRoomList           = "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/list"
	CorpOAMeetingRoomEdit           = "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/edit"
	CorpOAMeetingRoomDel            = "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/del"
	CorpOAGetMeetingRoomBookingInfo = "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/get_booking_info"
	CorpOAMeetingRoomBook           = "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/book"
	CorpOAMeetingRoomCancelBook     = "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/cancel_book"
	CorpOACallPstncc                = "https://qyapi.weixin.qq.com/cgi-bin/pstncc/call"
	CorpOAGetPstnccStates           = "https://qyapi.weixin.qq.com/cgi-bin/pstncc/getstates"
)

// tools
const (
	CorpToolsCalendarAdd              = "https://qyapi.weixin.qq.com/cgi-bin/oa/calendar/add"
	CorpToolsCalendarUpdate           = "https://qyapi.weixin.qq.com/cgi-bin/oa/calendar/update"
	CorpToolsCalendarGet              = "https://qyapi.weixin.qq.com/cgi-bin/oa/calendar/get"
	CorpToolsCalendarDel              = "https://qyapi.weixin.qq.com/cgi-bin/oa/calendar/del"
	CorpToolsScheduleAdd              = "https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/add"
	CorpToolsScheduleUpdate           = "https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/update"
	CorpToolsScheduleGet              = "https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/get"
	CorpToolsScheduleDel              = "https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/del"
	CorpToolsScheduleGetByCalendar    = "https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/get_by_calendar"
	CorpToolsMeetingCreate            = "https://qyapi.weixin.qq.com/cgi-bin/meeting/create"
	CorpToolsMeetingUpdate            = "https://qyapi.weixin.qq.com/cgi-bin/meeting/update"
	CorpToolsMeetingCancel            = "https://qyapi.weixin.qq.com/cgi-bin/meeting/cancel"
	CorpToolsMeetingGetUserMeetingID  = "https://qyapi.weixin.qq.com/cgi-bin/meeting/get_user_meetingid"
	CorpToolsMeetingGetInfo           = "https://qyapi.weixin.qq.com/cgi-bin/meeting/get_info"
	CorpToolsLivingCreate             = "https://qyapi.weixin.qq.com/cgi-bin/living/create"
	CorpToolsLivingModify             = "https://qyapi.weixin.qq.com/cgi-bin/living/modify"
	CorpToolsLivingCancel             = "https://qyapi.weixin.qq.com/cgi-bin/living/cancel"
	CorpToolsLivingDeleteReplayData   = "https://qyapi.weixin.qq.com/cgi-bin/living/delete_replay_data"
	CorpToolsLivingGetCode            = "https://qyapi.weixin.qq.com/cgi-bin/living/get_living_code"
	CorpToolsLivingGetUserAllLivingID = "https://qyapi.weixin.qq.com/cgi-bin/living/get_user_all_livingid"
	CorpToolsLivingGetInfo            = "https://qyapi.weixin.qq.com/cgi-bin/living/get_living_info"
	CorpToolsLivingGetWatchStat       = "https://qyapi.weixin.qq.com/cgi-bin/living/get_watch_stat"
	CorpToolsLivingGetShareInfo       = "https://qyapi.weixin.qq.com/cgi-bin/living/get_living_share_info"
)
