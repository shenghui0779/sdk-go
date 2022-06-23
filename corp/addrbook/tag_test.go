package addrbook

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/chenghonour/gochat/corp"
	"github.com/chenghonour/gochat/mock"
	"github.com/chenghonour/gochat/wx"
)

func TestCreateTag(t *testing.T) {
	body := []byte(`{"tagname":"UI"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "created",
    "tagid": 12
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/tag/create?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultTagCreate)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", CreateTag("UI", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultTagCreate{
		TagID: 12,
	}, result)
}

func TestUpdateTag(t *testing.T) {
	body := []byte(`{"tagid":12,"tagname":"UI design"}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/tag/update?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UpdateTag(12, "UI design"))

	assert.Nil(t, err)
}

func TestDeleteTag(t *testing.T) {
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/tag/delete?access_token=ACCESS_TOKEN&tagid=12", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteTag(12))

	assert.Nil(t, err)
}

func TestGetTagUser(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "tagname": "乒乓球协会",
    "userlist": [
        {
            "userid": "zhangsan",
            "name": "李四"
        }
    ],
    "partylist": [
        2
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/tag/get?access_token=ACCESS_TOKEN&tagid=12", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultTagUser)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetTagUser(12, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultTagUser{
		TagName: "乒乓球协会",
		UserList: []*TagUser{
			{
				UserID: "zhangsan",
				Name:   "李四",
			},
		},
		PartyList: []int{2},
	}, result)
}

func TestAddTagUser(t *testing.T) {
	body := []byte(`{"tagid":12,"userlist":["user1","user2"],"partylist":[4]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "invalidlist": "usr1|usr2|usr",
    "invalidparty": [
        2,
        4
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/tag/addtagusers?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	userIDs := []string{"user1", "user2"}
	partyIDs := []int64{4}

	result := new(ResultTagUserAdd)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", AddTagUser(12, userIDs, partyIDs, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultTagUserAdd{
		InvalidList:  "usr1|usr2|usr",
		InvalidParty: []int64{2, 4},
	}, result)
}

func TestDeleteTagUser(t *testing.T) {
	body := []byte(`{"tagid":12,"userlist":["user1","user2"],"partylist":[4]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "deleted",
    "invalidlist": "usr1|usr2|usr",
    "invalidparty": [
        2,
        4
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/tag/deltagusers?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	userIDs := []string{"user1", "user2"}
	partyIDs := []int64{4}

	result := new(ResultTagUserDelete)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteTagUser(12, userIDs, partyIDs, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultTagUserDelete{
		InvalidList:  "usr1|usr2|usr",
		InvalidParty: []int64{2, 4},
	}, result)
}

func TestListTag(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "taglist": [
        {
            "tagid": 1,
            "tagname": "a"
        },
        {
            "tagid": 2,
            "tagname": "b"
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/tag/list?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultTagList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListTag(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultTagList{
		TagList: []*Tag{
			{
				TagID:   1,
				TagName: "a",
			},
			{
				TagID:   2,
				TagName: "b",
			},
		},
	}, result)
}
