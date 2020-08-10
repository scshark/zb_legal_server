// 自动生成模板ZbUser
package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

// 如果含有time.Time 请自行import time包
type ZbUser struct {
	gorm.Model
	UUID         uuid.UUID `json:"uuid" form:"uuid" gorm:"column:uuid;comment:'uuid';type:varchar(255)"`
	NickName     string    `json:"nickName" form:"nickName" gorm:"column:nickname;comment:'用户昵称';type:varchar(255)"`
	HeaderImg    string    `json:"headerImg" form:"headerImg" gorm:"column:headerImg;comment:'微信用户头像';type:varchar(255)"`
	Mobile       string    `json:"mobile" form:"mobile" gorm:"column:mobile;comment:'用户手机号码';type:varchar(20)"`
	RegisteredAt time.Time `json:"registered_at" form:"registered_at" gorm:"column:registered_at;comment:'注册日期';type:timestamp"`
	Province     string    `json:"province" form:"province" gorm:"column:province;comment:'省';type:varchar(20)"`
	City         string    `json:"city" form:"city" gorm:"column:city;comment:'市';type:varchar(20)"`
	District     string    `json:"district" form:"district" gorm:"column:district;comment:'区';type:varchar(20)"`
	BrowseNum    int       `json:"browseNum" form:"browseNum" gorm:"column:browse_num;comment:'用户浏览次数';type:int(10)"`
	DownloadNum  int       `json:"downloadNum" form:"downloadNum" gorm:"column:download_num;comment:'下载次数';type:int(10)"`
	Status       int       `json:"status" form:"status" gorm:"column:status;comment:'用户状态 0：禁用 1：正常';type:smallint(1)"`
}

func (ZbUser) TableName() string {
	return "zb_user"
}
