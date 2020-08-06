package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

// @Tags ZbUserBrowseRecord
// @Summary 创建UserBrowseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbUserBrowseRecord true "创建UserBrowseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /browseRecord/createUserBrowseRecord [post]
func CreateUserBrowseRecord(c *gin.Context) {
	var browseRecord model.ZbUserBrowseRecord
	_ = c.ShouldBindJSON(&browseRecord)
	err := service.CreateUserBrowseRecord(browseRecord)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ZbUserBrowseRecord
// @Summary 删除UserBrowseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbUserBrowseRecord true "删除UserBrowseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /browseRecord/deleteUserBrowseRecord [delete]
func DeleteUserBrowseRecord(c *gin.Context) {
	var browseRecord model.ZbUserBrowseRecord
	_ = c.ShouldBindJSON(&browseRecord)
	err := service.DeleteUserBrowseRecord(browseRecord)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ZbUserBrowseRecord
// @Summary 批量删除UserBrowseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserBrowseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /browseRecord/deleteUserBrowseRecordByIds [delete]
func DeleteUserBrowseRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteUserBrowseRecordByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ZbUserBrowseRecord
// @Summary 更新UserBrowseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbUserBrowseRecord true "更新UserBrowseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /browseRecord/updateUserBrowseRecord [put]
func UpdateUserBrowseRecord(c *gin.Context) {
	var browseRecord model.ZbUserBrowseRecord
	_ = c.ShouldBindJSON(&browseRecord)
	err := service.UpdateUserBrowseRecord(&browseRecord)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ZbUserBrowseRecord
// @Summary 用id查询UserBrowseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbUserBrowseRecord true "用id查询UserBrowseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /browseRecord/findUserBrowseRecord [get]
func FindUserBrowseRecord(c *gin.Context) {
	var browseRecord model.ZbUserBrowseRecord
	_ = c.ShouldBindQuery(&browseRecord)
	err, rebrowseRecord := service.GetUserBrowseRecord(browseRecord.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"rebrowseRecord": rebrowseRecord}, c)
	}
}

// @Tags ZbUserBrowseRecord
// @Summary 分页获取UserBrowseRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UserBrowseRecordSearch true "分页获取UserBrowseRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /browseRecord/getUserBrowseRecordList [get]
func GetUserBrowseRecordList(c *gin.Context) {
	var pageInfo request.UserBrowseRecordSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetUserBrowseRecordInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
