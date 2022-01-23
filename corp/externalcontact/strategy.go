package externalcontact

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsCustomerStrategyList struct {
	Cursor string `json:"cursor,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

type ResultCustomerStrategyList struct {
	Strategy   []*CustomerStrategyListData `json:"strategy"`
	NextCursor string                      `json:"next_cursor"`
}

type CustomerStrategyListData struct {
	StrategyID int64 `json:"strategy_id"`
}

func ListCustomerStrategy(params *ParamsCustomerStrategyList, result *ResultCustomerStrategyList) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactCustomerStrategyList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type CustomerStrategy struct {
	StrategyID   int64                      `json:"strategy_id"`
	ParentID     int64                      `json:"parent_id"`
	StrategyName string                     `json:"strategy_name"`
	CreateTime   int64                      `json:"create_time"`
	AdminList    []string                   `json:"admin_list"`
	Privilege    *CustomerStrategyPrivilege `json:"privilege"`
}

type CustomerStrategyPrivilege struct {
	ViewCustomerList        bool `json:"view_customer_list"`
	ViewCustomerData        bool `json:"view_customer_data"`
	ViewRoomList            bool `json:"view_room_list"`
	ContactMe               bool `json:"contact_me"`
	JoinRoom                bool `json:"join_room"`
	ShareCustomer           bool `json:"share_customer"`
	OperResignCustomer      bool `json:"oper_resign_customer"`
	OperResignGroup         bool `json:"oper_resign_group"`
	SendCustomerMsg         bool `json:"send_customer_msg"`
	EditWelcomeMsg          bool `json:"edit_welcome_msg"`
	ViewBehaviorData        bool `json:"view_behavior_data"`
	ViewRoomData            bool `json:"view_room_data"`
	SendGroupMsg            bool `json:"send_group_msg"`
	RoomDeduplication       bool `json:"room_deduplication"`
	RapidReply              bool `json:"rapid_reply"`
	OnjobCustomerTransfer   bool `json:"onjob_customer_transfer"`
	EditAntiSpamRule        bool `json:"edit_anti_spam_rule"`
	ExportCustomerList      bool `json:"export_customer_list"`
	ExportCustomerData      bool `json:"export_customer_data"`
	ExportCustomerGroupList bool `json:"export_customer_group_list"`
	ManageCustomerTag       bool `json:"manage_customer_tag"`
}

type ParamsCustomerStrategyGet struct {
	StrategyID int64 `json:"strategy_id"`
}

type ResultCustomerStrategyGet struct {
	Strategy *CustomerStrategy `json:"strategy"`
}

func GetCustomerStrategy(strategyID int64, result *ResultCustomerStrategyGet) wx.Action {
	params := &ParamsCustomerStrategyGet{
		StrategyID: strategyID,
	}

	return wx.NewPostAction(urls.CorpExternalContactCustomerStrategyGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type CustomerStrategyRange struct {
	Type    int    `json:"type"`
	UserID  string `json:"userid,omitempty"`
	PartyID int64  `json:"partyid,omitempty"`
}

type ParamsCustomerStrategyRange struct {
	StrategyID int64  `json:"strategy_id"`
	Cursor     string `json:"cursor,omitempty"`
	Limit      int    `json:"limit,omitempty"`
}

type ResultCustomerStrategyRange struct {
	Range []*CustomerStrategyRange `json:"range"`
}

func GetCustomerStrategyRange(params *ParamsCustomerStrategyRange, result *ResultCustomerStrategyRange) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactCustomerStrategyGetRange,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsCustomerStrategyCreate struct {
	ParentID     int64                      `json:"parent_id,omitempty"`
	StrategyName string                     `json:"strategy_name"`
	AdminList    []string                   `json:"admin_list"`
	Privilege    *CustomerStrategyPrivilege `json:"privilege"`
	Range        []*CustomerStrategyRange   `json:"range"`
}

type ResultCustomerStrategyCreate struct {
	StrategyID int64 `json:"strategy_id"`
}

func CreateCustomerStrategy(params *ParamsCustomerStrategyCreate, result *ResultCustomerStrategyCreate) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactCustomerStrategyCreate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsCustomerStrategyEdit struct {
	StrategyID   int64                      `json:"strategy_id"`
	StrategyName string                     `json:"strategy_name,omitempty"`
	AdminList    []string                   `json:"admin_list,omitempty"`
	Privilege    *CustomerStrategyPrivilege `json:"privilege,omitempty"`
	RangeAdd     []*CustomerStrategyRange   `json:"range_add,omitempty"`
	RangeDel     []*CustomerStrategyRange   `json:"range_del,omitempty"`
}

func EditCustomerStrategy(params *ParamsCustomerStrategyEdit) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactCustomerStrategyEdit,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsCustomerStrategyDelete struct {
	StrategyID int64 `json:"strategy_id"`
}

func DeleteCustomerStrategy(strategyID int64) wx.Action {
	params := &ParamsCustomerStrategyDelete{
		StrategyID: strategyID,
	}

	return wx.NewPostAction(urls.CorpExternalContactCustomerStrategyDelete,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsMomentStrategyList struct {
	Cursor string `json:"cursor,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

type MomentStrategyListData struct {
	StrategyID int64 `json:"strategy_id"`
}

type ResultMomentStrategyList struct {
	Strategy   []*MomentStrategyListData `json:"strategy"`
	NextCursor string                    `json:"next_cursor"`
}

func ListMomentStrategy(params *ParamsMomentStrategyList, result *ResultMomentStrategyList) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactMomentStrategyList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsMomentStrategyGet struct {
	StrategyID int64 `json:"strategy_id"`
}

type MomentStrategy struct {
	StrategyID   int64                    `json:"strategy_id"`
	ParentID     int64                    `json:"parent_id"`
	StrategyName string                   `json:"strategy_name"`
	CreateTime   int64                    `json:"create_time"`
	AdminList    []string                 `json:"admin_list"`
	Privilege    *MomentStrategyPrivilege `json:"privilege"`
}

type MomentStrategyPrivilege struct {
	SendMoment               bool `json:"send_moment"`
	ViewMomentList           bool `json:"view_moment_list"`
	ManageMomentCoverAndSign bool `json:"manage_moment_cover_and_sign"`
}

type ResultMomentStrategyGet struct {
	Strategy *MomentStrategy `json:"strategy"`
}

func GetMomentStrategy(strategyID int64, result *ResultMomentStrategyGet) wx.Action {
	params := &ParamsMomentStrategyGet{
		StrategyID: strategyID,
	}

	return wx.NewPostAction(urls.CorpExternalContactMomentStrategyGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type MomentStrategyRange struct {
	Type    int    `json:"type"`
	UserID  string `json:"userid,omitempty"`
	PartyID int64  `json:"partyid,omitempty"`
}

type ParamsMomentStrategyRange struct {
	StrategyID int64  `json:"strategy_id"`
	Cursor     string `json:"cursor,omitempty"`
	Limit      int    `json:"limit,omitempty"`
}

type ResultMomentStrategyRange struct {
	Range []*MomentStrategyRange `json:"range"`
}

func GetMomentStrategyRange(params *ParamsMomentStrategyRange, result *ResultMomentStrategyRange) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactMomentStrategyGetRange,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsMomentStrategyCreate struct {
	ParentID     int64                    `json:"parent_id,omitempty"`
	StrategyName string                   `json:"strategy_name"`
	AdminList    []string                 `json:"admin_list"`
	Privilege    *MomentStrategyPrivilege `json:"privilege"`
	Range        []*MomentStrategyRange   `json:"range"`
}

type ResultMomentStrategyCreate struct {
	StrategyID int64 `json:"strategy_id"`
}

func CreateMomentStrategy(params *ParamsMomentStrategyCreate, result *ResultMomentStrategyCreate) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactMomentStrategyCreate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsMomentStrategyEdit struct {
	StrategyID   int64                    `json:"strategy_id"`
	StrategyName string                   `json:"strategy_name,omitempty"`
	AdminList    []string                 `json:"admin_list,omitempty"`
	Privilege    *MomentStrategyPrivilege `json:"privilege,omitempty"`
	RangeAdd     []*MomentStrategyRange   `json:"range_add,omitempty"`
	RangeDel     []*MomentStrategyRange   `json:"range_del,omitempty"`
}

func EditMomentStrategy(params *ParamsMomentStrategyEdit) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactMomentStrategyEdit,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsMomentStrategyDelete struct {
	StrategyID int64 `json:"strategy_id"`
}

func DeleteMomentStrategy(strategyID int64) wx.Action {
	params := &ParamsMomentStrategyDelete{
		StrategyID: strategyID,
	}

	return wx.NewPostAction(urls.CorpExternalContactMomentStrategyDelete,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}
