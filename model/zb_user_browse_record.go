// 自动生成模板UserBrowseRecord
package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 如果含有time.Time 请自行import time包
type ZbUserBrowseRecord struct {
      gorm.Model
      UserId  int `json:"userId" form:"userId" gorm:"column:user_id;comment:'id';type:int(10)"`
      DocumentId  int `json:"doucmentId" form:"doucmentId" gorm:"column:document_id;comment:'文书Id';type:int(10)"`
      BrowsedAt  time.Time `json:"browsedAt" form:"browsedAt" gorm:"column:browsed_at;comment:'浏览时间';type:timestamp"`
      RepeatNum  int `json:"repeatNum" form:"repeatNum" gorm:"column:repeat_num;comment:'重复次数';type:int(10)"` 
}


func (ZbUserBrowseRecord) TableName() string {
  return "zb_user_browse_record"
}
