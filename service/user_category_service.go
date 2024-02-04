package service

import (
	"qiqi-go/middleware/midAuto"
	"qiqi-go/response"
	"qiqi-go/response/BuildCategory"
)

// UserCategoryService 前端请求过来的数据
type UserCategoryService struct{}

// UserCategoryImages 给用户返回base64码的图片
func (service *UserCategoryService) UserCategoryImages() response.Response {
	codeId, base64, err := midAuto.CreateCode()
	if err != nil {
		return response.Response{Status: 501, Data: nil, Msg: "系统异常", Error: ""}
	}
	return response.Response{Status: 201, Data: BuildCategory.ResponseBuildCategoryImage(codeId, base64), Msg: "", Error: ""}
}
