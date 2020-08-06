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

// @Tags ZbUserDownloadRecord
// @Summary 创建UserDownloadRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbUserDownloadRecord true "创建UserDownloadRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /downloadRecord/createUserDownloadRecord [post]
func CreateUserDownloadRecord(c *gin.Context) {
	var downloadRecord model.ZbUserDownloadRecord
	_ = c.ShouldBindJSON(&downloadRecord)
	err := service.CreateUserDownloadRecord(downloadRecord)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ZbUserDownloadRecord
// @Summary 删除UserDownloadRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbUserDownloadRecord true "删除UserDownloadRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /downloadRecord/deleteUserDownloadRecord [delete]
func DeleteUserDownloadRecord(c *gin.Context) {
	var downloadRecord model.ZbUserDownloadRecord
	_ = c.ShouldBindJSON(&downloadRecord)
	err := service.DeleteUserDownloadRecord(downloadRecord)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ZbUserDownloadRecord
// @Summary 批量删除UserDownloadRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserDownloadRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /downloadRecord/deleteUserDownloadRecordByIds [delete]
func DeleteUserDownloadRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteUserDownloadRecordByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ZbUserDownloadRecord
// @Summary 更新UserDownloadRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbUserDownloadRecord true "更新UserDownloadRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /downloadRecord/updateUserDownloadRecord [put]
func UpdateUserDownloadRecord(c *gin.Context) {
	var downloadRecord model.ZbUserDownloadRecord
	_ = c.ShouldBindJSON(&downloadRecord)
	err := service.UpdateUserDownloadRecord(&downloadRecord)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ZbUserDownloadRecord
// @Summary 用id查询UserDownloadRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbUserDownloadRecord true "用id查询UserDownloadRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /downloadRecord/findUserDownloadRecord [get]
func FindUserDownloadRecord(c *gin.Context) {
	var downloadRecord model.ZbUserDownloadRecord
	_ = c.ShouldBindQuery(&downloadRecord)
	err, redownloadRecord := service.GetUserDownloadRecord(downloadRecord.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"redownloadRecord": redownloadRecord}, c)
	}
}

// @Tags ZbUserDownloadRecord
// @Summary 分页获取UserDownloadRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UserDownloadRecordSearch true "分页获取UserDownloadRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /downloadRecord/getUserDownloadRecordList [get]
func GetUserDownloadRecordList(c *gin.Context) {
	var pageInfo request.UserDownloadRecordSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetUserDownloadRecordInfoList(pageInfo)
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
