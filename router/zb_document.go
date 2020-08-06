package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDocumentRouter(Router *gin.RouterGroup) {
	DocumentRouter := Router.Group("doc").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		DocumentRouter.POST("createDocument", v1.CreateDocument)   // 新建Document
		DocumentRouter.DELETE("deleteDocument", v1.DeleteDocument) // 删除Document
		DocumentRouter.DELETE("deleteDocumentByIds", v1.DeleteDocumentByIds) // 批量删除Document
		DocumentRouter.PUT("updateDocument", v1.UpdateDocument)    // 更新Document
		DocumentRouter.GET("findDocument", v1.FindDocument)        // 根据ID获取Document
		DocumentRouter.GET("getDocumentList", v1.GetDocumentList)  // 获取Document列表
	}
}
