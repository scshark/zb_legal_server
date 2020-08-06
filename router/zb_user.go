package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitZbUserRouter(Router *gin.RouterGroup) {
	ZbUserRouter := Router.Group("zbUser").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		ZbUserRouter.POST("createZbUser", v1.CreateZbUser)   // 新建ZbUser
		ZbUserRouter.DELETE("deleteZbUser", v1.DeleteZbUser) // 删除ZbUser
		ZbUserRouter.DELETE("deleteZbUserByIds", v1.DeleteZbUserByIds) // 批量删除ZbUser
		ZbUserRouter.PUT("updateZbUser", v1.UpdateZbUser)    // 更新ZbUser
		ZbUserRouter.GET("findZbUser", v1.FindZbUser)        // 根据ID获取ZbUser
		ZbUserRouter.GET("getZbUserList", v1.GetZbUserList)  // 获取ZbUser列表
	}
}
