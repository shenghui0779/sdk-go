package tools

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/corp"
	"github.com/shenghui0779/gochat/mock"
)

func TestCreateWedriveSpace(t *testing.T) {
	body := []byte(`{"userid":"USERID","space_name":"SPACE_NAME","auth_info":[{"type":1,"userid":"USERID","auth":2},{"type":2,"departmentid":1,"auth":1}]}`)
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "spaceid": "SPACEID"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/wedrive/space_create?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	params := &ParamsWedriveSpaceCreate{
		UserID:    "USERID",
		SpaceName: "SPACE_NAME",
		AuthInfo: []*WedriveAuthInfo{
			{
				Type:   1,
				UserID: "USERID",
				Auth:   2,
			},
			{
				Type:         2,
				DepartmentID: 1,
				Auth:         1,
			},
		},
	}
	result := new(ResultWedriveSpaceCreate)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", CreateWedriveSpace(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultWedriveSpaceCreate{
		SpaceID: "SPACEID",
	}, result)
}

func TestRenameWedriveSpace(t *testing.T) {
	body := []byte(`{"userid":"USERID","spaceid":"SPACEID","space_name":"SPACE_NAME"}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/wedrive/space_rename?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", RenameWedriveSpace("USERID", "SPACEID", "SPACE_NAME"))

	assert.Nil(t, err)
}

func TestDismissWedriveSpace(t *testing.T) {
	body := []byte(`{"userid":"USERID","spaceid":"SPACEID"}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/wedrive/space_dismiss?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DismissWedriveSpace("USERID", "SPACEID"))

	assert.Nil(t, err)
}

func TestGetWedriveSpaceInfo(t *testing.T) {
	body := []byte(`{"userid":"USERID","spaceid":"SPACEID"}`)
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "space_info": {
        "spaceid": "SPACEID",
        "space_name": "SPACE_NAME",
        "auth_list": {
            "auth_info": [
                {
                    "type": 1,
                    "userid": "USERID1",
                    "auth": 3
                },
                {
                    "type": 1,
                    "userid": "USERID2",
                    "auth": 2
                },
                {
                    "type": 2,
                    "departmentid": 1,
                    "auth": 1
                }
            ],
            "quit_userid": [
                "USERID3",
                "USERID4"
            ]
        }
    }
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/wedrive/space_info?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultWedriveSpaceInfo)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetWedriveSpaceInfo("USERID", "SPACEID", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultWedriveSpaceInfo{
		SpaceInfo: &WedriveSpaceInfo{
			SpaceID:   "SPACEID",
			SpaceName: "SPACE_NAME",
			AuthList: &WedriveSpaceAuthList{
				AuthInfo: []*WedriveAuthInfo{
					{
						Type:   1,
						UserID: "USERID1",
						Auth:   3,
					},
					{
						Type:   1,
						UserID: "USERID2",
						Auth:   2,
					},
					{
						Type:         2,
						DepartmentID: 1,
						Auth:         1,
					},
				},
				QuitUserID: []string{"USERID3", "USERID4"},
			},
		},
	}, result)
}

func TestAddWedriveSpaceAcl(t *testing.T) {
	body := []byte(`{"userid":"USERID","spaceid":"SPACEID","auth_info":[{"type":1,"userid":"USERID1","auth":2},{"type":2,"departmentid":1,"auth":2}]}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/wedrive/space_acl_add?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	acls := []*WedriveAuthInfo{
		{
			Type:   1,
			UserID: "USERID1",
			Auth:   2,
		},
		{
			Type:         2,
			DepartmentID: 1,
			Auth:         2,
		},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", AddWedriveSpaceAcl("USERID", "SPACEID", acls...))

	assert.Nil(t, err)
}

func TestDeleteWedriveSpaceAcl(t *testing.T) {
	body := []byte(`{"userid":"USERID","spaceid":"SPACEID","auth_info":[{"type":1,"userid":"USERID1"},{"type":2,"departmentid":1}]}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/wedrive/space_acl_del?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	acls := []*WedriveAuthInfo{
		{
			Type:   1,
			UserID: "USERID1",
		},
		{
			Type:         2,
			DepartmentID: 1,
		},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteWedriveSpaceAcl("USERID", "SPACEID", acls...))

	assert.Nil(t, err)
}

func TestSetWedriveSpace(t *testing.T) {
	body := []byte(`{"userid":"USERID","spaceid":"SPACEID","enable_watermark":true,"add_member_only_admin":true,"enable_share_url":false,"share_url_no_approve":true,"share_url_no_approve_default_auth":4}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/wedrive/space_setting?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	enable := new(bool)
	*enable = true

	params := &ParamsWedriveSpaceSetting{
		UserID:                       "USERID",
		SpaceID:                      "SPACEID",
		EnableWatermark:              enable,
		AddMemberOnlyAdmin:           enable,
		EnableShareURL:               new(bool),
		ShareURLNoApprove:            enable,
		ShareURLNoApproveDefaultAuth: 4,
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SetWedriveSpace(params))

	assert.Nil(t, err)
}

func TestShareWedriveSpace(t *testing.T) {
	body := []byte(`{"userid":"USERID","spaceid":"SPACEID"}`)
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "space_share_url": "SPACE_SHARE_URL"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/wedrive/space_share?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultWedriveSpaceShare)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ShareWedriveSpace("USERID", "SPACEID", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultWedriveSpaceShare{
		SpaceShareURL: "SPACE_SHARE_URL",
	}, result)
}

func TestListWedriveFile(t *testing.T) {
	body := []byte(`{"userid":"USERID","spaceid":"SPACEID","fatherid":"FATHERID","sort_type":6,"start":0,"limit":100}`)
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "has_more": true,
    "next_start": 100,
    "file_list": {
        "item": [
            {
                "fileid": "FILEID1",
                "file_name": "FILE_NAME1",
                "spaceid": "SPACEID",
                "fatherid": "FATHERID",
                "file_size": 10240,
                "ctime": 123456789,
                "mtime": 123456789,
                "file_type": 1,
                "file_status": 1,
                "create_userid": "CREATE_USERID",
                "update_userid": "UPDATE_USERID",
                "sha": "SHA",
                "md5": "MD5",
                "url": "URL"
            },
            {
                "fileid": "FILEID2",
                "file_name": "FILE_NAME2"
            }
        ]
    }
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_list?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	params := &ParamsWedriveFileList{
		UserID:   "USERID",
		SpaceID:  "SPACEID",
		FatherID: "FATHERID",
		SortType: 6,
		Start:    0,
		Limit:    100,
	}
	result := new(ResultWedriveFileList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListWedriveFile(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultWedriveFileList{
		HasMore:   true,
		NextStart: 100,
		FileList: &WedriveFileList{
			Item: []*WedriveFileInfo{
				{
					FileID:       "FILEID1",
					FileName:     "FILE_NAME1",
					SpaceID:      "SPACEID",
					FatherID:     "FATHERID",
					FileSize:     10240,
					CTime:        123456789,
					MTime:        123456789,
					FileType:     1,
					FileStatus:   1,
					CreateUserID: "CREATE_USERID",
					UpdateUserID: "UPDATE_USERID",
					SHA:          "SHA",
					MD5:          "MD5",
					URL:          "URL",
				},
				{
					FileID:   "FILEID2",
					FileName: "FILE_NAME2",
				},
			},
		},
	}, result)
}

func TestUploadWedriveFile(t *testing.T) {
	body := []byte(`{"userid":"USERID","spaceid":"SPACEID","fatherid":"FATHERID","file_name":"FILE_NAME","file_base64_content":"FILE_BASE64_CONTENT"}`)
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "fileid": "FILEID"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_upload?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	params := &ParamsWedriveFileUpload{
		UserID:            "USERID",
		SpaceID:           "SPACEID",
		FatherID:          "FATHERID",
		FileName:          "FILE_NAME",
		FileBase64Content: "FILE_BASE64_CONTENT",
	}
	result := new(ResultWedriveFileUpload)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UploadWedriveFile(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultWedriveFileUpload{
		FileID: "FILEID",
	}, result)
}

func TestDownloadWedriveFile(t *testing.T) {
	body := []byte(`{"userid":"USERID","fileid":"FILEID"}`)
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "download_url": "DOWNLOAD_URL",
    "cookie_name": "COOKIE_NAME",
    "cookie_value": "COOKIE_VALUE"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_download?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultWedriveFileDownload)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DownloadWedriveFile("USERID", "FILEID", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultWedriveFileDownload{
		DownloadURL: "DOWNLOAD_URL",
		CookieName:  "COOKIE_NAME",
		CookieValue: "COOKIE_VALUE",
	}, result)
}

func TestCreateWedriveFile(t *testing.T) {
	body := []byte(`{"userid":"USERID","spaceid":"SPACEID","fatherid":"FATHERID","file_type":"FILE_TYPE","file_name":"FILE_NAME"}`)
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "fileid": "FILEID",
    "url": "URL"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_create?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	params := &ParamsWedriveFileCreate{
		UserID:   "USERID",
		SpaceID:  "SPACEID",
		FatherID: "FATHERID",
		FileType: "FILE_TYPE",
		FileName: "FILE_NAME",
	}
	result := new(ResultWedriveFileCreate)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", CreateWedriveFile(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultWedriveFileCreate{
		FileID: "FILEID",
		URL:    "URL",
	}, result)
}

func TestRenameWedriveFile(t *testing.T) {
	body := []byte(`{"userid":"USERID","fileid":"FILEID","new_name":"NEW_NAME"}`)
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "file": {
        "fileid": "FILEID",
        "file_name": "FILE_NAME",
        "spaceid": "SPACEID",
        "fatherid": "FATHERID",
        "file_size": 10240,
        "ctime": 123456789,
        "mtime": 123456789,
        "file_type": 1,
        "file_status": 1,
        "create_userid": "CREATE_USERID",
        "update_userid": "UPDATE_USERID",
        "sha": "SHA",
        "md5": "MD5"
    }
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_rename?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultWedriveFileRename)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", RenameWedriveFile("USERID", "FILEID", "NEW_NAME", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultWedriveFileRename{
		File: &WedriveFileInfo{
			FileID:       "FILEID",
			FileName:     "FILE_NAME",
			SpaceID:      "SPACEID",
			FatherID:     "FATHERID",
			FileSize:     10240,
			CTime:        123456789,
			MTime:        123456789,
			FileType:     1,
			FileStatus:   1,
			CreateUserID: "CREATE_USERID",
			UpdateUserID: "UPDATE_USERID",
			SHA:          "SHA",
			MD5:          "MD5",
		},
	}, result)
}

func TestMoveWedriveFile(t *testing.T) {
	body := []byte(`{"userid":"USERID","fatherid":"FATHERID","replace":true,"fileid":["FILEID1","FILEID2"]}`)
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "file_list": {
        "item": [
            {
                "fileid": "FILEID",
                "file_name": "FILE_NAME",
                "spaceid": "SPACEID",
                "fatherid": "FATHERID",
                "file_size": 10240,
                "ctime": 123456789,
                "mtime": 123456789,
                "file_type": 1,
                "file_status": 1,
                "create_userid": "CREATE_USERID",
                "update_userid": "UPDATE_USERID",
                "sha": "SHA",
                "md5": "MD5"
            }
        ]
    }
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_move?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	params := &ParamsWedriveFileMove{
		UserID:   "USERID",
		FatherID: "FATHERID",
		Replace:  true,
		FileID:   []string{"FILEID1", "FILEID2"},
	}
	result := new(ResultWedriveFileMove)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", MoveWedriveFile(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultWedriveFileMove{
		FileList: &WedriveFileList{
			Item: []*WedriveFileInfo{
				{
					FileID:       "FILEID",
					FileName:     "FILE_NAME",
					SpaceID:      "SPACEID",
					FatherID:     "FATHERID",
					FileSize:     10240,
					CTime:        123456789,
					MTime:        123456789,
					FileType:     1,
					FileStatus:   1,
					CreateUserID: "CREATE_USERID",
					UpdateUserID: "UPDATE_USERID",
					SHA:          "SHA",
					MD5:          "MD5",
				},
			},
		},
	}, result)
}

func TestDeleteWedriveFile(t *testing.T) {
	body := []byte(`{"userid":"USERID","fileid":["FILEID1","FILEID2"]}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_delete?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteWedriveFile("USERID", "FILEID1", "FILEID2"))

	assert.Nil(t, err)
}

func TestGetWedriveFileInfo(t *testing.T) {
	body := []byte(`{"userid":"USERID","fileid":"FILEID"}`)
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "file_info": {
        "fileid": "FILEID",
        "file_name": "FILE_NAME",
        "spaceid": "SPACEID",
        "fatherid": "FATHERID",
        "file_size": 10240,
        "ctime": 123456789,
        "mtime": 123456789,
        "file_type": 1,
        "file_status": 1,
        "create_userid": "CREATE_USERID",
        "update_userid": "UPDATE_USERID",
        "sha": "SHA",
        "md5": "MD5",
        "url": "URL"
    }
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_info?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultWedriveFileInfo)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetWedriveFileInfo("USERID", "FILEID", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultWedriveFileInfo{
		FileInfo: &WedriveFileInfo{
			FileID:       "FILEID",
			FileName:     "FILE_NAME",
			SpaceID:      "SPACEID",
			FatherID:     "FATHERID",
			FileSize:     10240,
			CTime:        123456789,
			MTime:        123456789,
			FileType:     1,
			FileStatus:   1,
			CreateUserID: "CREATE_USERID",
			UpdateUserID: "UPDATE_USERID",
			SHA:          "SHA",
			MD5:          "MD5",
			URL:          "URL",
		},
	}, result)
}

func TestAddWedriveFileAcl(t *testing.T) {
	body := []byte(`{"userid":"USERID","fileid":"FILEID","auth_info":[{"type":1,"userid":"USERID1","auth":1},{"type":2,"departmentid":1,"auth":1}]}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_acl_add?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	acls := []*WedriveAuthInfo{
		{
			Type:   1,
			UserID: "USERID1",
			Auth:   1,
		},
		{
			Type:         2,
			DepartmentID: 1,
			Auth:         1,
		},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", AddWedriveFileAcl("USERID", "FILEID", acls...))

	assert.Nil(t, err)
}

func TestDeleteWedriveFileAcl(t *testing.T) {
	body := []byte(`{"userid":"USERID","fileid":"FILEID","auth_info":[{"type":1,"userid":"USERID1"},{"type":2,"departmentid":1}]}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_acl_del?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	acls := []*WedriveAuthInfo{
		{
			Type:   1,
			UserID: "USERID1",
		},
		{
			Type:         2,
			DepartmentID: 1,
		},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteWedriveFileAcl("USERID", "FILEID", acls...))

	assert.Nil(t, err)
}

func TestSetWedriveFile(t *testing.T) {
	body := []byte(`{"userid":"USERID","fileid":"FILDID","auth_scope":1,"auth":1}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_setting?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	params := &ParamsWedriveFileSetting{
		UserID:    "USERID",
		FileID:    "FILDID",
		AuthScope: 1,
		Auth:      1,
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SetWedriveFile(params))

	assert.Nil(t, err)
}

func TestShareWedriveFile(t *testing.T) {
	body := []byte(`{"userid":"USERID","fileid":"FILEID"}`)
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "share_url": "SHARE_URL"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_share?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultWedriveFileShare)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ShareWedriveFile("USERID", "FILEID", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultWedriveFileShare{
		ShareURL: "SHARE_URL",
	}, result)
}
