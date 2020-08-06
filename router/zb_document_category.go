package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDocumentCategoryRouter(Router *gin.RouterGroup) {
	DocumentCategoryRouter := Router.Group("docCategory").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		DocumentCategoryRouter.POST("createDocumentCategory", v1.CreateDocumentCategory)   // 新建DocumentCategory
		DocumentCategoryRouter.DELETE("deleteDocumentCategory", v1.DeleteDocumentCategory) // 删除DocumentCategory
		DocumentCategoryRouter.DELETE("deleteDocumentCategoryByIds", v1.DeleteDocumentCategoryByIds) // 批量删除DocumentCategory
		DocumentCategoryRouter.PUT("updateDocumentCategory", v1.UpdateDocumentCategory)    // 更新DocumentCategory
		DocumentCategoryRouter.GET("findDocumentCategory", v1.FindDocumentCategory)        // 根据ID获取DocumentCategory
		DocumentCategoryRouter.GET("getDocumentCategoryList", v1.GetDocumentCategoryList)  // 获取DocumentCategory列表
	}
}
