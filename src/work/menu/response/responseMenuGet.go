package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/work/menu/request"
)

type ResponseMenuGet struct {
	response.ResponseWork

	// 企业微信官方：返回结果与请参考菜单创建接口
	// https://developer.work.weixin.qq.com/document/path/90232
	request.RequestMenuSet
}
