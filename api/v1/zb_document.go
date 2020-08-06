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

// @Tags ZbDocument
// @Summary 创建Document
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbDocument true "创建Document"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /doc/createDocument [post]
func CreateDocument(c *gin.Context) {
	var doc model.ZbDocument
	_ = c.ShouldBindJSON(&doc)
	err := service.CreateDocument(doc)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ZbDocument
// @Summary 删除Document
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbDocument true "删除Document"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /doc/deleteDocument [delete]
func DeleteDocument(c *gin.Context) {
	var doc model.ZbDocument
	_ = c.ShouldBindJSON(&doc)
	err := service.DeleteDocument(doc)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ZbDocument
// @Summary 批量删除Document
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Document"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /doc/deleteDocumentByIds [delete]
func DeleteDocumentByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteDocumentByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ZbDocument
// @Summary 更新Document
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbDocument true "更新Document"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /doc/updateDocument [put]
func UpdateDocument(c *gin.Context) {
	var doc model.ZbDocument
	_ = c.ShouldBindJSON(&doc)
	err := service.UpdateDocument(&doc)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ZbDocument
// @Summary 用id查询Document
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbDocument true "用id查询Document"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /doc/findDocument [get]
func FindDocument(c *gin.Context) {
	var doc model.ZbDocument
	_ = c.ShouldBindQuery(&doc)
	err, redoc := service.GetDocument(doc.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"redoc": redoc}, c)
	}
}

// @Tags ZbDocument
// @Summary 分页获取Document列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DocumentSearch true "分页获取Document列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /doc/getDocumentList [get]
func GetDocumentList(c *gin.Context) {
	var pageInfo request.DocumentSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetDocumentInfoList(pageInfo)
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
