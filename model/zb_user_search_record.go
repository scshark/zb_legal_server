// 自动生成模板UserSearchRecord
package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 如果含有time.Time 请自行import time包
type ZbUserSearchRecord struct {
      gorm.Model
      UserId  int `json:"userId" form:"userId" gorm:"column:user_id;comment:'id';type:int(10)"`
      Keyword  string `json:"keyword" form:"keyword" gorm:"column:keyword;comment:'关键词';type:varchar(50)"`
      SearchedAt  time.Time `json:"searchedAt" form:"searchedAt" gorm:"column:searched_at;comment:'检索时间';type:timestamp"`
      RepeatNum  int `json:"repeatNum" form:"repeatNum" gorm:"column:repeat_num;comment:'重复次数';type:int(10)"` 
}


func (ZbUserSearchRecord) TableName() string {
  return "zb_user_search_record"
}
