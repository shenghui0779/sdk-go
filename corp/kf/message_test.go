package kf

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/corp"
	"github.com/shenghui0779/gochat/event"
	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/wx"
)

func TestSyncMsg(t *testing.T) {
	body := []byte(`{"cursor":"4gw7MepFLfgF2VC5npN","token":"ENCApHxnGDNAVNY4AaSJKj4Tb5mwsEMzxhFmHVGcra996NR","limit":1000,"voice_format":0}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "next_cursor": "4gw7MepFLfgF2VC5npN",
    "has_more": 1,
    "msg_list": [
        {
            "msgid": "from_msgid_4622416642169452483",
            "open_kfid": "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
            "external_userid": "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
            "send_time": 1615478585,
            "origin": 3,
            "servicer_userid": "Zhangsan",
            "msgtype": "text",
            "text": {
                "content": "hello world",
                "menu_id": "101"
            }
        },
        {
            "msgid": "from_msgid_4622416642169452483",
            "open_kfid": "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
            "external_userid": "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
            "send_time": 1615478585,
            "origin": 3,
            "servicer_userid": "Zhangsan",
            "msgtype": "image",
            "image": {
                "media_id": "2iSLeVyqzk4eX0IB5kTi9Ljfa2rt9dwfq5WKRQ4Nvvgw"
            }
        },
        {
            "msgid": "from_msgid_4622416642169452483",
            "open_kfid": "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
            "external_userid": "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
            "send_time": 1615478585,
            "origin": 3,
            "servicer_userid": "Zhangsan",
            "msgtype": "voice",
            "voice": {
                "media_id": "2iSLeVyqzk4eX0IB5kTi9Ljfa2rt9dwfq5WKRQ4Nvvgw"
            }
        },
        {
            "msgid": "from_msgid_4622416642169452483",
            "open_kfid": "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
            "external_userid": "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
            "send_time": 1615478585,
            "origin": 3,
            "servicer_userid": "Zhangsan",
            "msgtype": "video",
            "video": {
                "media_id": "2iSLeVyqzk4eX0IB5kTi9Ljfa2rt9dwfq5WKRQ4Nvvgw"
            }
        },
        {
            "msgid": "from_msgid_4622416642169452483",
            "open_kfid": "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
            "external_userid": "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
            "send_time": 1615478585,
            "origin": 3,
            "servicer_userid": "Zhangsan",
            "msgtype": "file",
            "file": {
                "media_id": "2iSLeVyqzk4eX0IB5kTi9Ljfa2rt9dwfq5WKRQ4Nvvgw"
            }
        },
        {
            "msgid": "from_msgid_4622416642169452483",
            "open_kfid": "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
            "external_userid": "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
            "send_time": 1615478585,
            "origin": 3,
            "servicer_userid": "Zhangsan",
            "msgtype": "location",
            "location": {
                "latitude": 23.106021881103501,
                "longitude": 113.320503234863,
                "name": "广州国际媒体港(广州市海珠区)",
                "address": "广东省广州市海珠区滨江东路"
            }
        },
        {
            "msgid": "from_msgid_4622416642169452483",
            "open_kfid": "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
            "external_userid": "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
            "send_time": 1615478585,
            "origin": 3,
            "servicer_userid": "Zhangsan",
            "msgtype": "link",
            "link": {
                "title": "TITLE",
                "desc": "DESC",
                "url": "URL",
                "pic_url": "PIC_URL"
            }
        },
        {
            "msgid": "from_msgid_4622416642169452483",
            "open_kfid": "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
            "external_userid": "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
            "send_time": 1615478585,
            "origin": 3,
            "servicer_userid": "Zhangsan",
            "msgtype": "business_card",
            "business_card": {
                "userid": "USERID"
            }
        },
        {
            "msgid": "from_msgid_4622416642169452483",
            "open_kfid": "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
            "external_userid": "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
            "send_time": 1615478585,
            "origin": 3,
            "servicer_userid": "Zhangsan",
            "msgtype": "miniprogram",
            "miniprogram": {
                "title": "TITLE",
                "appid": "APPID",
                "pagepath": "PAGE_PATH",
                "thumb_media_id": "THUMB_MEDIA_ID"
            }
        },
        {
            "msgid": "from_msgid_4622416642169452483",
            "open_kfid": "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
            "external_userid": "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
            "send_time": 1615478585,
            "origin": 3,
            "servicer_userid": "Zhangsan",
            "msgtype": "msgmenu",
            "msgmenu": {
                "head_content": "您对本次服务是否满意呢? ",
                "list": [
                    {
                        "type": "click",
                        "click": {
                            "id": "101",
                            "content": "满意"
                        }
                    },
                    {
                        "type": "click",
                        "click": {
                            "id": "102",
                            "content": "不满意"
                        }
                    },
                    {
                        "type": "view",
                        "view": {
                            "url": "https://work.weixin.qq.com",
                            "content": "点击跳转到自助查询页面"
                        }
                    },
                    {
                        "type": "miniprogram",
                        "miniprogram": {
                            "appid": "wx123123123123123",
                            "pagepath": "pages/index?userid=zhangsan&orderid=123123123",
                            "content": "点击打开小程序查询更多"
                        }
                    }
                ],
                "tail_content": "欢迎再次光临"
            }
        },
        {
            "msgid": "from_msgid_4622416642169452483",
            "open_kfid": "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
            "external_userid": "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
            "send_time": 1615478585,
            "origin": 3,
            "servicer_userid": "Zhangsan",
            "msgtype": "event",
            "event": {
                "event_type": "enter_session",
                "open_kfid": "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
                "external_userid": "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
                "scene": "123",
                "scene_param": "abc",
                "welcome_code": "aaaaaa"
            }
        },
        {
            "msgid": "from_msgid_4622416642169452483",
            "open_kfid": "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
            "external_userid": "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
            "send_time": 1615478585,
            "origin": 3,
            "servicer_userid": "Zhangsan",
            "msgtype": "event",
            "event": {
                "event_type": "msg_send_fail",
                "open_kfid": "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
                "external_userid": "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
                "fail_msgid": "FAIL_MSGID",
                "fail_type": 4
            }
        },
        {
            "msgid": "from_msgid_4622416642169452483",
            "open_kfid": "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
            "external_userid": "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
            "send_time": 1615478585,
            "origin": 3,
            "servicer_userid": "Zhangsan",
            "msgtype": "event",
            "event": {
                "event_type": "servicer_status_change",
                "servicer_userid": "SERVICER_USERID",
                "status": 1,
                "open_kfid": "OPEN_KFID"
            }
        },
        {
            "msgid": "from_msgid_4622416642169452483",
            "open_kfid": "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
            "external_userid": "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
            "send_time": 1615478585,
            "origin": 3,
            "servicer_userid": "Zhangsan",
            "msgtype": "event",
            "event": {
                "event_type": "session_status_change",
                "open_kfid": "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
                "external_userid": "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
                "change_type": 1,
                "old_servicer_userid": "OLD_SERVICER_USERID",
                "new_servicer_userid": "NEW_SERVICER_USERID",
                "msg_code": "MSG_CODE"
            }
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/kf/sync_msg?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsMsgSync{
		Cursor: "4gw7MepFLfgF2VC5npN",
		Token:  "ENCApHxnGDNAVNY4AaSJKj4Tb5mwsEMzxhFmHVGcra996NR",
		Limit:  1000,
	}

	result := new(ResultMsgSync)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SyncMsg(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMsgSync{
		NextCursor: "4gw7MepFLfgF2VC5npN",
		HasMore:    1,
		MsgList: []*MsgListData{
			{
				MsgID:          "from_msgid_4622416642169452483",
				OpenKFID:       "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
				ExternalUserID: "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
				SendTime:       1615478585,
				Origin:         3,
				ServicerUserID: "Zhangsan",
				MsgType:        event.MsgText,
				Text: &Text{
					Content: "hello world",
					MenuID:  "101",
				},
			},
			{
				MsgID:          "from_msgid_4622416642169452483",
				OpenKFID:       "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
				ExternalUserID: "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
				SendTime:       1615478585,
				Origin:         3,
				ServicerUserID: "Zhangsan",
				MsgType:        event.MsgImage,
				Image: &Image{
					MediaID: "2iSLeVyqzk4eX0IB5kTi9Ljfa2rt9dwfq5WKRQ4Nvvgw",
				},
			},
			{
				MsgID:          "from_msgid_4622416642169452483",
				OpenKFID:       "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
				ExternalUserID: "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
				SendTime:       1615478585,
				Origin:         3,
				ServicerUserID: "Zhangsan",
				MsgType:        event.MsgVoice,
				Voice: &Voice{
					MediaID: "2iSLeVyqzk4eX0IB5kTi9Ljfa2rt9dwfq5WKRQ4Nvvgw",
				},
			},
			{
				MsgID:          "from_msgid_4622416642169452483",
				OpenKFID:       "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
				ExternalUserID: "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
				SendTime:       1615478585,
				Origin:         3,
				ServicerUserID: "Zhangsan",
				MsgType:        event.MsgVideo,
				Video: &Video{
					MediaID: "2iSLeVyqzk4eX0IB5kTi9Ljfa2rt9dwfq5WKRQ4Nvvgw",
				},
			},
			{
				MsgID:          "from_msgid_4622416642169452483",
				OpenKFID:       "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
				ExternalUserID: "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
				SendTime:       1615478585,
				Origin:         3,
				ServicerUserID: "Zhangsan",
				MsgType:        event.MsgFile,
				File: &File{
					MediaID: "2iSLeVyqzk4eX0IB5kTi9Ljfa2rt9dwfq5WKRQ4Nvvgw",
				},
			},
			{
				MsgID:          "from_msgid_4622416642169452483",
				OpenKFID:       "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
				ExternalUserID: "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
				SendTime:       1615478585,
				Origin:         3,
				ServicerUserID: "Zhangsan",
				MsgType:        event.MsgLocation,
				Location: &Location{
					Latitude:  23.106021881103501,
					Longitude: 113.320503234863,
					Name:      "广州国际媒体港(广州市海珠区)",
					Address:   "广东省广州市海珠区滨江东路",
				},
			},
			{
				MsgID:          "from_msgid_4622416642169452483",
				OpenKFID:       "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
				ExternalUserID: "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
				SendTime:       1615478585,
				Origin:         3,
				ServicerUserID: "Zhangsan",
				MsgType:        event.MsgLink,
				Link: &Link{
					Title:  "TITLE",
					Desc:   "DESC",
					URL:    "URL",
					PicURL: "PIC_URL",
				},
			},
			{
				MsgID:          "from_msgid_4622416642169452483",
				OpenKFID:       "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
				ExternalUserID: "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
				SendTime:       1615478585,
				Origin:         3,
				ServicerUserID: "Zhangsan",
				MsgType:        event.MsgBussinessCard,
				BussinessCard: &BusinessCard{
					UserID: "USERID",
				},
			},
			{
				MsgID:          "from_msgid_4622416642169452483",
				OpenKFID:       "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
				ExternalUserID: "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
				SendTime:       1615478585,
				Origin:         3,
				ServicerUserID: "Zhangsan",
				MsgType:        event.MsgMinip,
				Minip: &Minip{
					Title:        "TITLE",
					AppID:        "APPID",
					PagePath:     "PAGE_PATH",
					ThumbMediaID: "THUMB_MEDIA_ID",
				},
			},
			{
				MsgID:          "from_msgid_4622416642169452483",
				OpenKFID:       "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
				ExternalUserID: "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
				SendTime:       1615478585,
				Origin:         3,
				ServicerUserID: "Zhangsan",
				MsgType:        event.MsgMsgMenu,
				Menu: &Menu{
					HeadContent: "您对本次服务是否满意呢? ",
					TailContent: "欢迎再次光临",
					List: []*MenuItem{
						{
							Type: MenuClick,
							Click: &ClickMenu{
								ID:      "101",
								Content: "满意",
							},
						},
						{
							Type: MenuClick,
							Click: &ClickMenu{
								ID:      "102",
								Content: "不满意",
							},
						},
						{
							Type: MenuView,
							View: &ViewMenu{
								URL:     "https://work.weixin.qq.com",
								Content: "点击跳转到自助查询页面",
							},
						},
						{
							Type: MenuMinip,
							Minip: &MinipMenu{
								AppID:    "wx123123123123123",
								PagePath: "pages/index?userid=zhangsan&orderid=123123123",
								Content:  "点击打开小程序查询更多",
							},
						},
					},
				},
			},
			{
				MsgID:          "from_msgid_4622416642169452483",
				OpenKFID:       "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
				ExternalUserID: "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
				SendTime:       1615478585,
				Origin:         3,
				ServicerUserID: "Zhangsan",
				MsgType:        event.MsgEvent,
				Event: &Event{
					EventType:      event.EventEnterSession,
					OpenKFID:       "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
					ExternalUserID: "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
					Scene:          "123",
					SceneParam:     "abc",
					WelcomeCode:    "aaaaaa",
				},
			},
			{
				MsgID:          "from_msgid_4622416642169452483",
				OpenKFID:       "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
				ExternalUserID: "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
				SendTime:       1615478585,
				Origin:         3,
				ServicerUserID: "Zhangsan",
				MsgType:        event.MsgEvent,
				Event: &Event{
					EventType:      event.EventMsgSendFail,
					OpenKFID:       "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
					ExternalUserID: "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
					FailMsgID:      "FAIL_MSGID",
					FailType:       4,
				},
			},
			{
				MsgID:          "from_msgid_4622416642169452483",
				OpenKFID:       "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
				ExternalUserID: "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
				SendTime:       1615478585,
				Origin:         3,
				ServicerUserID: "Zhangsan",
				MsgType:        event.MsgEvent,
				Event: &Event{
					EventType:      event.EventServicerStatusChange,
					OpenKFID:       "OPEN_KFID",
					ServicerUserID: "SERVICER_USERID",
					Status:         1,
				},
			},
			{
				MsgID:          "from_msgid_4622416642169452483",
				OpenKFID:       "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
				ExternalUserID: "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
				SendTime:       1615478585,
				Origin:         3,
				ServicerUserID: "Zhangsan",
				MsgType:        event.MsgEvent,
				Event: &Event{
					EventType:         event.EventSessionStatusChange,
					OpenKFID:          "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
					ExternalUserID:    "wmAJ2GCAAAme1XQRC-NI-q0_ZM9ukoAw",
					OldServicerUserID: "OLD_SERVICER_USERID",
					NewServicerUserID: "NEW_SERVICER_USERID",
					MsgCode:           "MSG_CODE",
				},
			},
		},
	}, result)
}

func TestSendTextMsg(t *testing.T) {
	body := []byte(`{"touser":"EXTERNAL_USERID","open_kfid":"OPEN_KFID","msgid":"MSGID","msgtype":"text","text":{"content":"你购买的物品已发货，可点击链接查看物流状态http://work.weixin.qq.com/xxxxxx"}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "msgid": "MSG_ID"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/kf/send_msg?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	text := &Text{
		Content: "",
		MenuID:  "",
	}

	result := new(ResultMsgSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN")

	assert.Nil(t, err)
	assert.Equal(t, &ResultMsgSend{
		MsgID: "MSG_ID",
	}, result)
}

func TestSendImageMsg(t *testing.T) {
	body := []byte(`{"touser":"EXTERNAL_USERID","open_kfid":"OPEN_KFID","msgid":"MSGID","msgtype":"image","image":{"media_id":"MEDIA_ID"}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "msgid": "MSG_ID"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/kf/send_msg?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMsgSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN")

	assert.Nil(t, err)
	assert.Equal(t, &ResultMsgSend{
		MsgID: "MSG_ID",
	}, result)
}

func TestSendVoiceMsg(t *testing.T) {
	body := []byte(`{"touser":"EXTERNAL_USERID","open_kfid":"OPEN_KFID","msgtype":"voice","voice":{"media_id":"MEDIA_ID"}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "msgid": "MSG_ID"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/kf/send_msg?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMsgSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN")

	assert.Nil(t, err)
	assert.Equal(t, &ResultMsgSend{
		MsgID: "MSG_ID",
	}, result)
}

func TestSendVideoMsg(t *testing.T) {
	body := []byte(`{"touser":"EXTERNAL_USERID","open_kfid":"OPEN_KFID","msgid":"MSGID","msgtype":"video","video":{"media_id":"MEDIA_ID"}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "msgid": "MSG_ID"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/kf/send_msg?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMsgSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN")

	assert.Nil(t, err)
	assert.Equal(t, &ResultMsgSend{
		MsgID: "MSG_ID",
	}, result)
}

func TestSendFileMsg(t *testing.T) {
	body := []byte(`{"touser":"EXTERNAL_USERID","open_kfid":"OPEN_KFID","msgid":"MSGID","msgtype":"file","file":{"media_id":"1Yv-zXfHjSjU-7LH-GwtYqDGS-zz6w22KmWAT5COgP7o"}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "msgid": "MSG_ID"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/kf/send_msg?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMsgSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN")

	assert.Nil(t, err)
	assert.Equal(t, &ResultMsgSend{
		MsgID: "MSG_ID",
	}, result)
}

func TestSendLinkMsg(t *testing.T) {
	body := []byte(`{"touser":"EXTERNAL_USERID","open_kfid":"OPEN_KFID","msgid":"MSGID","msgtype":"link","link":{"title":"企业如何增长？企业微信给出3个答案","desc":"今年中秋节公司有豪礼相送","url":"URL","thumb_media_id":"MEDIA_ID"}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "msgid": "MSG_ID"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/kf/send_msg?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMsgSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN")

	assert.Nil(t, err)
	assert.Equal(t, &ResultMsgSend{
		MsgID: "MSG_ID",
	}, result)
}

func TestSendMinipMsg(t *testing.T) {
	body := []byte(`{"touser":"EXTERNAL_USERID","open_kfid":"OPEN_KFID","msgid":"MSGID","msgtype":"miniprogram","miniprogram":{"appid":"APPID","title":"欢迎报名夏令营","thumb_media_id":"MEDIA_ID","pagepath":"PAGE_PATH"}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "msgid": "MSG_ID"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/kf/send_msg?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMsgSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN")

	assert.Nil(t, err)
	assert.Equal(t, &ResultMsgSend{
		MsgID: "MSG_ID",
	}, result)
}

func TestSendMenuMsg(t *testing.T) {
	body := []byte(`{"touser":"EXTERNAL_USERID","open_kfid":"OPEN_KFID","msgid":"MSGID","msgtype":"msgmenu","msgmenu":{"head_content":"您对本次服务是否满意呢? ","tail_content":"欢迎再次光临","list":[{"type":"click","click":{"id":"101","content":"满意"}},{"type":"click","click":{"id":"102","content":"不满意"}},{"type":"view","view":{"url":"https://work.weixin.qq.com","content":"点击跳转到自助查询页面"}},{"type":"miniprogram","miniprogram":{"appid":"wx123123123123123","pagepath":"pages/index?userid=zhangsan&orderid=123123123","content":"点击打开小程序查询更多"}}]}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "msgid": "MSG_ID"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/kf/send_msg?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMsgSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN")

	assert.Nil(t, err)
	assert.Equal(t, &ResultMsgSend{
		MsgID: "MSG_ID",
	}, result)
}

func TestSendLocationMsg(t *testing.T) {
	body := []byte(`{"touser":"EXTERNAL_USERID","open_kfid":"OPEN_KFID","msgid":"MSGID","msgtype":"location","location":{"name":"测试小区","address":"实例小区，不真实存在，经纬度无意义","latitude":0,"longitude":0}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "msgid": "MSG_ID"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/kf/send_msg?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMsgSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN")

	assert.Nil(t, err)
	assert.Equal(t, &ResultMsgSend{
		MsgID: "MSG_ID",
	}, result)
}
