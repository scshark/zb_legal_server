// 自动生成模板Label
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type ZbLabel struct {
	gorm.Model
	Name        string `json:"name" form:"name" gorm:"column:name;comment:'标签名称';type:varchar(50)"`
	Description string `json:"description" form:"description" gorm:"column:description;comment:'描述';type:varchar(255)"`
	Status      int    `json:"status" form:"status" gorm:"column:status;comment:'状态  0：禁用 1:启用';type:smallint(1)"`
	UserNum     int    `json:"userNum" form:"userNum" gorm:"column:user_num;comment:'标签用户数量';type:int(10)"`
}

func (ZbLabel) TableName() string {
	return "zb_label"
}
