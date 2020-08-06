package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserBrowseRecordRouter(Router *gin.RouterGroup) {
	UserBrowseRecordRouter := Router.Group("browseRecord").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		UserBrowseRecordRouter.POST("createUserBrowseRecord", v1.CreateUserBrowseRecord)   // 新建UserBrowseRecord
		UserBrowseRecordRouter.DELETE("deleteUserBrowseRecord", v1.DeleteUserBrowseRecord) // 删除UserBrowseRecord
		UserBrowseRecordRouter.DELETE("deleteUserBrowseRecordByIds", v1.DeleteUserBrowseRecordByIds) // 批量删除UserBrowseRecord
		UserBrowseRecordRouter.PUT("updateUserBrowseRecord", v1.UpdateUserBrowseRecord)    // 更新UserBrowseRecord
		UserBrowseRecordRouter.GET("findUserBrowseRecord", v1.FindUserBrowseRecord)        // 根据ID获取UserBrowseRecord
		UserBrowseRecordRouter.GET("getUserBrowseRecordList", v1.GetUserBrowseRecordList)  // 获取UserBrowseRecord列表
	}
}
