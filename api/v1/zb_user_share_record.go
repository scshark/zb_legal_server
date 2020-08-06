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

// @Tags ZbUserShareRecord
// @Summary 创建UserShareRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbUserShareRecord true "创建UserShareRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shareRecord/createUserShareRecord [post]
func CreateUserShareRecord(c *gin.Context) {
	var shareRecord model.ZbUserShareRecord
	_ = c.ShouldBindJSON(&shareRecord)
	err := service.CreateUserShareRecord(shareRecord)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ZbUserShareRecord
// @Summary 删除UserShareRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbUserShareRecord true "删除UserShareRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shareRecord/deleteUserShareRecord [delete]
func DeleteUserShareRecord(c *gin.Context) {
	var shareRecord model.ZbUserShareRecord
	_ = c.ShouldBindJSON(&shareRecord)
	err := service.DeleteUserShareRecord(shareRecord)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ZbUserShareRecord
// @Summary 批量删除UserShareRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserShareRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shareRecord/deleteUserShareRecordByIds [delete]
func DeleteUserShareRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteUserShareRecordByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ZbUserShareRecord
// @Summary 更新UserShareRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbUserShareRecord true "更新UserShareRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shareRecord/updateUserShareRecord [put]
func UpdateUserShareRecord(c *gin.Context) {
	var shareRecord model.ZbUserShareRecord
	_ = c.ShouldBindJSON(&shareRecord)
	err := service.UpdateUserShareRecord(&shareRecord)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ZbUserShareRecord
// @Summary 用id查询UserShareRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbUserShareRecord true "用id查询UserShareRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shareRecord/findUserShareRecord [get]
func FindUserShareRecord(c *gin.Context) {
	var shareRecord model.ZbUserShareRecord
	_ = c.ShouldBindQuery(&shareRecord)
	err, reshareRecord := service.GetUserShareRecord(shareRecord.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"reshareRecord": reshareRecord}, c)
	}
}

// @Tags ZbUserShareRecord
// @Summary 分页获取UserShareRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UserShareRecordSearch true "分页获取UserShareRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shareRecord/getUserShareRecordList [get]
func GetUserShareRecordList(c *gin.Context) {
	var pageInfo request.UserShareRecordSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetUserShareRecordInfoList(pageInfo)
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
