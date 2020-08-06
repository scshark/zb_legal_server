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

// @Tags ZbDocumentKeyword
// @Summary 创建DocumentKeyword
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbDocumentKeyword true "创建DocumentKeyword"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /docKeyword/createDocumentKeyword [post]
func CreateDocumentKeyword(c *gin.Context) {
	var docKeyword model.ZbDocumentKeyword
	_ = c.ShouldBindJSON(&docKeyword)
	err := service.CreateDocumentKeyword(docKeyword)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ZbDocumentKeyword
// @Summary 删除DocumentKeyword
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbDocumentKeyword true "删除DocumentKeyword"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /docKeyword/deleteDocumentKeyword [delete]
func DeleteDocumentKeyword(c *gin.Context) {
	var docKeyword model.ZbDocumentKeyword
	_ = c.ShouldBindJSON(&docKeyword)
	err := service.DeleteDocumentKeyword(docKeyword)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ZbDocumentKeyword
// @Summary 批量删除DocumentKeyword
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DocumentKeyword"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /docKeyword/deleteDocumentKeywordByIds [delete]
func DeleteDocumentKeywordByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteDocumentKeywordByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ZbDocumentKeyword
// @Summary 更新DocumentKeyword
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbDocumentKeyword true "更新DocumentKeyword"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /docKeyword/updateDocumentKeyword [put]
func UpdateDocumentKeyword(c *gin.Context) {
	var docKeyword model.ZbDocumentKeyword
	_ = c.ShouldBindJSON(&docKeyword)
	err := service.UpdateDocumentKeyword(&docKeyword)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ZbDocumentKeyword
// @Summary 用id查询DocumentKeyword
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbDocumentKeyword true "用id查询DocumentKeyword"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /docKeyword/findDocumentKeyword [get]
func FindDocumentKeyword(c *gin.Context) {
	var docKeyword model.ZbDocumentKeyword
	_ = c.ShouldBindQuery(&docKeyword)
	err, redocKeyword := service.GetDocumentKeyword(docKeyword.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"redocKeyword": redocKeyword}, c)
	}
}

// @Tags ZbDocumentKeyword
// @Summary 分页获取DocumentKeyword列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DocumentKeywordSearch true "分页获取DocumentKeyword列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /docKeyword/getDocumentKeywordList [get]
func GetDocumentKeywordList(c *gin.Context) {
	var pageInfo request.DocumentKeywordSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetDocumentKeywordInfoList(pageInfo)
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
