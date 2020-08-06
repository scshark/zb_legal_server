package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserLabelRouter(Router *gin.RouterGroup) {
	UserLabelRouter := Router.Group("userLabel").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		UserLabelRouter.POST("createUserLabel", v1.CreateUserLabel)   // 新建UserLabel
		UserLabelRouter.DELETE("deleteUserLabel", v1.DeleteUserLabel) // 删除UserLabel
		UserLabelRouter.DELETE("deleteUserLabelByIds", v1.DeleteUserLabelByIds) // 批量删除UserLabel
		UserLabelRouter.PUT("updateUserLabel", v1.UpdateUserLabel)    // 更新UserLabel
		UserLabelRouter.GET("findUserLabel", v1.FindUserLabel)        // 根据ID获取UserLabel
		UserLabelRouter.GET("getUserLabelList", v1.GetUserLabelList)  // 获取UserLabel列表
	}
}
