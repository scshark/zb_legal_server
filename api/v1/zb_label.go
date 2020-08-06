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

// @Tags ZbLabel
// @Summary 创建Label
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbLabel true "创建Label"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /label/createLabel [post]
func CreateLabel(c *gin.Context) {
	var label model.ZbLabel
	_ = c.ShouldBindJSON(&label)
	err := service.CreateLabel(label)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ZbLabel
// @Summary 删除Label
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbLabel true "删除Label"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /label/deleteLabel [delete]
func DeleteLabel(c *gin.Context) {
	var label model.ZbLabel
	_ = c.ShouldBindJSON(&label)
	err := service.DeleteLabel(label)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ZbLabel
// @Summary 批量删除Label
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Label"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /label/deleteLabelByIds [delete]
func DeleteLabelByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteLabelByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ZbLabel
// @Summary 更新Label
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbLabel true "更新Label"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /label/updateLabel [put]
func UpdateLabel(c *gin.Context) {
	var label model.ZbLabel
	_ = c.ShouldBindJSON(&label)
	err := service.UpdateLabel(&label)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ZbLabel
// @Summary 用id查询Label
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbLabel true "用id查询Label"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /label/findLabel [get]
func FindLabel(c *gin.Context) {
	var label model.ZbLabel
	_ = c.ShouldBindQuery(&label)
	err, relabel := service.GetLabel(label.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"relabel": relabel}, c)
	}
}

// @Tags ZbLabel
// @Summary 分页获取Label列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.LabelSearch true "分页获取Label列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /label/getLabelList [get]
func GetLabelList(c *gin.Context) {
	var pageInfo request.LabelSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetLabelInfoList(pageInfo)
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
