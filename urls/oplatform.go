/*
@Time : 2021/8/16 5:09 下午
@Author : 21
@File : consts
@Software: GoLand
*/
package urls

const  (
	BaseUrl = "https://mp.weixin.qq.com"
)

const (
	// 获取令牌
	ComponentApiComponentTokenUrl = "https://api.weixin.qq.com/cgi-bin/component/api_component_token"
	// 获取预授权码
	ComponentApiCreatePreAuthCode = "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode?component_access_token=%s"
	// 使用授权码获取授权信息
	ComponentApiQueryAuthUrl = "https://api.weixin.qq.com/cgi-bin/component/api_query_auth?component_access_token=%s"
	// 获取授权方的帐号基本信息
	ComponentApiGetAuthorizerInfoUrl = "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info?component_access_token=%s"
	)

const   (
	//关联小程序
	WxopenWxamplinkUrl = "https://api.weixin.qq.com/cgi-bin/wxopen/wxamplink?access_token=%s"
	//获取公众号关联的小程序
	WxopenWxamplinkGetUrl = "https://api.weixin.qq.com/cgi-bin/wxopen/wxamplinkget?access_token=%s"
)
