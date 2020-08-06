// 自动生成模板UserDownloadRecord
package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 如果含有time.Time 请自行import time包
type ZbUserDownloadRecord struct {
      gorm.Model
      UserId  int `json:"userId" form:"userId" gorm:"column:user_id;comment:'id';type:int(10)"`
      DocumentId  int `json:"doucmentId" form:"doucmentId" gorm:"column:document_id;comment:'文书Id';type:int(10)"`
      DownloadedAt  time.Time `json:"downloadedAt" form:"downloadedAt" gorm:"column:downloaded_at;comment:'下载时间';type:timestamp"`
      RepeatNum  int `json:"repeatNum" form:"repeatNum" gorm:"column:repeat_num;comment:'重复次数';type:int(10)"` 
}


func (ZbUserDownloadRecord) TableName() string {
  return "zb_user_download_record"
}
