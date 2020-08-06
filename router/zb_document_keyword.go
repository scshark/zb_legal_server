package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDocumentKeywordRouter(Router *gin.RouterGroup) {
	DocumentKeywordRouter := Router.Group("docKeyword").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		DocumentKeywordRouter.POST("createDocumentKeyword", v1.CreateDocumentKeyword)   // 新建DocumentKeyword
		DocumentKeywordRouter.DELETE("deleteDocumentKeyword", v1.DeleteDocumentKeyword) // 删除DocumentKeyword
		DocumentKeywordRouter.DELETE("deleteDocumentKeywordByIds", v1.DeleteDocumentKeywordByIds) // 批量删除DocumentKeyword
		DocumentKeywordRouter.PUT("updateDocumentKeyword", v1.UpdateDocumentKeyword)    // 更新DocumentKeyword
		DocumentKeywordRouter.GET("findDocumentKeyword", v1.FindDocumentKeyword)        // 根据ID获取DocumentKeyword
		DocumentKeywordRouter.GET("getDocumentKeywordList", v1.GetDocumentKeywordList)  // 获取DocumentKeyword列表
	}
}
