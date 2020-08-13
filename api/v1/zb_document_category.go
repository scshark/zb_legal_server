package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
)

// @Tags ZbDocumentCategory
// @Summary 创建DocumentCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbDocumentCategory true "创建DocumentCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /docCategory/createDocumentCategory [post]
func CreateDocumentCategory(c *gin.Context) {
	var docCategory model.ZbDocumentCategory
	_ = c.ShouldBindJSON(&docCategory)
	// 参数验证
	CategoryVerify := utils.Rules{
		"Title": {utils.NotEmpty()},
	}
	CategoryVerifyErr := utils.Verify(docCategory, CategoryVerify)
	if CategoryVerifyErr != nil {
		response.FailWithMessage(CategoryVerifyErr.Error(), c)
		return
	}
	// 数据逻辑
	err := service.CreateDocumentCategory(docCategory)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ZbDocumentCategory
// @Summary 删除DocumentCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbDocumentCategory true "删除DocumentCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /docCategory/deleteDocumentCategory [delete]
func DeleteDocumentCategory(c *gin.Context) {
	var docId request.GetById
	_ = c.ShouldBindJSON(&docId)
	verifyErr := utils.Verify(docId, utils.CustomizeMap["IdVerify"])
	if verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	err := service.DeleteDocumentCategory(uint(docId.Id))
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ZbDocumentCategory
// @Summary 批量删除DocumentCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DocumentCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /docCategory/deleteDocumentCategoryByIds [delete]
func DeleteDocumentCategoryByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteDocumentCategoryByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ZbDocumentCategory
// @Summary 更新DocumentCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbDocumentCategory true "更新DocumentCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /docCategory/updateDocumentCategory [put]
func UpdateDocumentCategory(c *gin.Context) {
	var docCategory model.ZbDocumentCategory
	_ = c.ShouldBindJSON(&docCategory)

	categoryVerify := utils.Rules{
		"Title":  {utils.NotEmpty()},
		"Sort":   {utils.Ge("0")},
	}
	categoryVerifyErr := utils.Verify(docCategory, categoryVerify)
	if categoryVerifyErr != nil {
		response.FailWithMessage(categoryVerifyErr.Error(), c)
		return
	}
	err := service.UpdateDocumentCategory(&docCategory)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ZbDocumentCategory
// @Summary 用id查询DocumentCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbDocumentCategory true "用id查询DocumentCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /docCategory/findDocumentCategory [get]
func FindDocumentCategory(c *gin.Context) {

	var docId request.GetById
	_ = c.ShouldBindQuery(&docId)
	verifyErr := utils.Verify(docId, utils.CustomizeMap["IdVerify"])
	if verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	err, docCategory := service.GetDocumentCategory(uint(docId.Id))
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"docCategory": docCategory}, c)
	}
}

// @Tags ZbDocumentCategory
// @Summary 分页获取DocumentCategory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DocumentCategorySearch true "分页获取DocumentCategory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /docCategory/getDocumentCategoryList [get]
func GetDocumentCategoryList(c *gin.Context) {
	var pageInfo request.DocumentCategorySearch
	_ = c.ShouldBindQuery(&pageInfo)

	err, list, total := service.GetDocumentCategoryInfoList(pageInfo)
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
