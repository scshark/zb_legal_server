package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"math/rand"
	"time"
)

// @Tags ZbUser
// @Summary 创建ZbUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbUser true "创建ZbUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /zbUser/createZbUser [post]
func CreateZbUser(c *gin.Context) {
	// var zbUser model.ZbUser
	zbUser := model.ZbUser{
		UUID:         uuid.NewV4(),
		NickName:     fmt.Sprintf("用户%d", rand.Intn(999999)),
		Mobile:       fmt.Sprintf("17722%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)),
		HeaderImg:    "http://thirdwx.qlogo.cn/mmopen/vi_32/DYAIOgq83er359PXMfat0K8jwfKZgQNk71ticmeddnLWOtCxibqa6Bpibvia24I2yLZADqNCciacePBNWTzdCUOJJRg/132",
		Province:     "广东",
		City:         "广州",
		District:     "天河",
		Status:       1,
		RegisteredAt: time.Now(),
	}
	// _ = c.ShouldBindJSON(&zbUser)
	err := service.CreateZbUser(zbUser)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ZbUser
// @Summary 删除ZbUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbUser true "删除ZbUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /zbUser/deleteZbUser [delete]
func DeleteZbUser(c *gin.Context) {
	var zbUser model.ZbUser
	_ = c.ShouldBindJSON(&zbUser)
	err := service.DeleteZbUser(zbUser)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ZbUser
// @Summary 批量删除ZbUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ZbUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /zbUser/deleteZbUserByIds [delete]
func DeleteZbUserByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteZbUserByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ZbUser
// @Summary 更新ZbUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbUser true "更新ZbUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /zbUser/updateZbUser [put]
func UpdateZbUser(c *gin.Context) {
	var zbUser model.ZbUser
	_ = c.ShouldBindJSON(&zbUser)
	err := service.UpdateZbUser(&zbUser)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ZbUser
// @Summary 用id查询ZbUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZbUser true "用id查询ZbUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /zbUser/findZbUser [get]
func FindZbUser(c *gin.Context) {
	var zbUser model.ZbUser
	_ = c.ShouldBindQuery(&zbUser)
	err, rezbUser := service.GetZbUser(zbUser.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"rezbUser": rezbUser}, c)
	}
}

// @Tags ZbUser
// @Summary 分页获取ZbUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ZbUserSearch true "分页获取ZbUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /zbUser/getZbUserList [get]
func GetZbUserList(c *gin.Context) {
	var pageInfo request.ZbUserSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetZbUserInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
