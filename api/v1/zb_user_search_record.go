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

// @Tags ZbUserSearchRecord
// @Summary 创建UserSearchRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbUserSearchRecord true "创建UserSearchRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /searchRecord/createUserSearchRecord [post]
func CreateUserSearchRecord(c *gin.Context) {
	var searchRecord model.ZbUserSearchRecord
	_ = c.ShouldBindJSON(&searchRecord)
	err := service.CreateUserSearchRecord(searchRecord)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ZbUserSearchRecord
// @Summary 删除UserSearchRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbUserSearchRecord true "删除UserSearchRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /searchRecord/deleteUserSearchRecord [delete]
func DeleteUserSearchRecord(c *gin.Context) {
	var searchRecord model.ZbUserSearchRecord
	_ = c.ShouldBindJSON(&searchRecord)
	err := service.DeleteUserSearchRecord(searchRecord)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ZbUserSearchRecord
// @Summary 批量删除UserSearchRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserSearchRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /searchRecord/deleteUserSearchRecordByIds [delete]
func DeleteUserSearchRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteUserSearchRecordByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ZbUserSearchRecord
// @Summary 更新UserSearchRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbUserSearchRecord true "更新UserSearchRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /searchRecord/updateUserSearchRecord [put]
func UpdateUserSearchRecord(c *gin.Context) {
	var searchRecord model.ZbUserSearchRecord
	_ = c.ShouldBindJSON(&searchRecord)
	err := service.UpdateUserSearchRecord(&searchRecord)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ZbUserSearchRecord
// @Summary 用id查询UserSearchRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbUserSearchRecord true "用id查询UserSearchRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /searchRecord/findUserSearchRecord [get]
func FindUserSearchRecord(c *gin.Context) {
	var searchRecord model.ZbUserSearchRecord
	_ = c.ShouldBindQuery(&searchRecord)
	err, researchRecord := service.GetUserSearchRecord(searchRecord.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"researchRecord": researchRecord}, c)
	}
}

// @Tags ZbUserSearchRecord
// @Summary 分页获取UserSearchRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UserSearchRecordSearch true "分页获取UserSearchRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /searchRecord/getUserSearchRecordList [get]
func GetUserSearchRecordList(c *gin.Context) {
	var pageInfo request.UserSearchRecordSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetUserSearchRecordInfoList(pageInfo)
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
