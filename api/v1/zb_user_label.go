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

// @Tags ZbUserLabel
// @Summary 创建UserLabel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbUserLabel true "创建UserLabel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userLabel/createUserLabel [post]
func CreateUserLabel(c *gin.Context) {
	var userLabel model.ZbUserLabel
	_ = c.ShouldBindJSON(&userLabel)
	err := service.CreateUserLabel(userLabel)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ZbUserLabel
// @Summary 删除UserLabel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbUserLabel true "删除UserLabel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userLabel/deleteUserLabel [delete]
func DeleteUserLabel(c *gin.Context) {
	var userLabel model.ZbUserLabel
	_ = c.ShouldBindJSON(&userLabel)
	err := service.DeleteUserLabel(userLabel)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ZbUserLabel
// @Summary 批量删除UserLabel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserLabel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userLabel/deleteUserLabelByIds [delete]
func DeleteUserLabelByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteUserLabelByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ZbUserLabel
// @Summary 更新UserLabel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbUserLabel true "更新UserLabel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userLabel/updateUserLabel [put]
func UpdateUserLabel(c *gin.Context) {
	var userLabel model.ZbUserLabel
	_ = c.ShouldBindJSON(&userLabel)
	err := service.UpdateUserLabel(&userLabel)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ZbUserLabel
// @Summary 用id查询UserLabel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbUserLabel true "用id查询UserLabel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userLabel/findUserLabel [get]
func FindUserLabel(c *gin.Context) {
	var userLabel model.ZbUserLabel
	_ = c.ShouldBindQuery(&userLabel)
	err, reuserLabel := service.GetUserLabel(userLabel.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"reuserLabel": reuserLabel}, c)
	}
}

// @Tags ZbUserLabel
// @Summary 分页获取UserLabel列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UserLabelSearch true "分页获取UserLabel列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userLabel/getUserLabelList [get]
func GetUserLabelList(c *gin.Context) {
	var pageInfo request.UserLabelSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetUserLabelInfoList(pageInfo)
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
