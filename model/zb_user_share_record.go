// 自动生成模板UserShareRecord
package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 如果含有time.Time 请自行import time包
type ZbUserShareRecord struct {
      gorm.Model
      UserId  int `json:"userId" form:"userId" gorm:"column:user_id;comment:'id';type:int(10)"`
      DocumentId  int `json:"doucmentId" form:"doucmentId" gorm:"column:document_id;comment:'文书Id';type:int(10)"`
      SharedAt  time.Time `json:"sharedAt" form:"sharedAt" gorm:"column:shared_at;comment:'分享时间';type:timestamp"`
      Type  int `json:"type" form:"type" gorm:"column:type;comment:'分享类型 0:分享小程序 1:分享文书';type:smallint(1)"` 
}


func (ZbUserShareRecord) TableName() string {
  return "zb_user_share_record"
}
