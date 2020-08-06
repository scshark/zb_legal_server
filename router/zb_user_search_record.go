package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserSearchRecordRouter(Router *gin.RouterGroup) {
	UserSearchRecordRouter := Router.Group("searchRecord").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		UserSearchRecordRouter.POST("createUserSearchRecord", v1.CreateUserSearchRecord)   // 新建UserSearchRecord
		UserSearchRecordRouter.DELETE("deleteUserSearchRecord", v1.DeleteUserSearchRecord) // 删除UserSearchRecord
		UserSearchRecordRouter.DELETE("deleteUserSearchRecordByIds", v1.DeleteUserSearchRecordByIds) // 批量删除UserSearchRecord
		UserSearchRecordRouter.PUT("updateUserSearchRecord", v1.UpdateUserSearchRecord)    // 更新UserSearchRecord
		UserSearchRecordRouter.GET("findUserSearchRecord", v1.FindUserSearchRecord)        // 根据ID获取UserSearchRecord
		UserSearchRecordRouter.GET("getUserSearchRecordList", v1.GetUserSearchRecordList)  // 获取UserSearchRecord列表
	}
}
