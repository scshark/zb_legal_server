// 自动生成模板UserLabel
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type ZbUserLabel struct {
      gorm.Model
      UserId  int `json:"userId" form:"userId" gorm:"column:user_id;comment:'用户id';type:int(10)"`
      LabelId  int `json:"labelId" form:"labelId" gorm:"column:label_id;comment:'标签id';type:int(10)"` 
}


func (ZbUserLabel) TableName() string {
  return "zb_user_label"
}
