package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitLabelRouter(Router *gin.RouterGroup) {
	LabelRouter := Router.Group("label").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		LabelRouter.POST("createLabel", v1.CreateLabel)   // 新建Label
		LabelRouter.DELETE("deleteLabel", v1.DeleteLabel) // 删除Label
		LabelRouter.DELETE("deleteLabelByIds", v1.DeleteLabelByIds) // 批量删除Label
		LabelRouter.PUT("updateLabel", v1.UpdateLabel)    // 更新Label
		LabelRouter.GET("findLabel", v1.FindLabel)        // 根据ID获取Label
		LabelRouter.GET("getLabelList", v1.GetLabelList)  // 获取Label列表
	}
}
