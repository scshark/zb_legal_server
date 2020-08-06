package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserDownloadRecordRouter(Router *gin.RouterGroup) {
	UserDownloadRecordRouter := Router.Group("downloadRecord").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		UserDownloadRecordRouter.POST("createUserDownloadRecord", v1.CreateUserDownloadRecord)   // 新建UserDownloadRecord
		UserDownloadRecordRouter.DELETE("deleteUserDownloadRecord", v1.DeleteUserDownloadRecord) // 删除UserDownloadRecord
		UserDownloadRecordRouter.DELETE("deleteUserDownloadRecordByIds", v1.DeleteUserDownloadRecordByIds) // 批量删除UserDownloadRecord
		UserDownloadRecordRouter.PUT("updateUserDownloadRecord", v1.UpdateUserDownloadRecord)    // 更新UserDownloadRecord
		UserDownloadRecordRouter.GET("findUserDownloadRecord", v1.FindUserDownloadRecord)        // 根据ID获取UserDownloadRecord
		UserDownloadRecordRouter.GET("getUserDownloadRecordList", v1.GetUserDownloadRecordList)  // 获取UserDownloadRecord列表
	}
}
