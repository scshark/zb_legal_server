package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserShareRecordRouter(Router *gin.RouterGroup) {
	UserShareRecordRouter := Router.Group("shareRecord").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		UserShareRecordRouter.POST("createUserShareRecord", v1.CreateUserShareRecord)   // 新建UserShareRecord
		UserShareRecordRouter.DELETE("deleteUserShareRecord", v1.DeleteUserShareRecord) // 删除UserShareRecord
		UserShareRecordRouter.DELETE("deleteUserShareRecordByIds", v1.DeleteUserShareRecordByIds) // 批量删除UserShareRecord
		UserShareRecordRouter.PUT("updateUserShareRecord", v1.UpdateUserShareRecord)    // 更新UserShareRecord
		UserShareRecordRouter.GET("findUserShareRecord", v1.FindUserShareRecord)        // 根据ID获取UserShareRecord
		UserShareRecordRouter.GET("getUserShareRecordList", v1.GetUserShareRecordList)  // 获取UserShareRecord列表
	}
}
