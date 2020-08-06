package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateUserSearchRecord
// @description   create a ZbUserSearchRecord
// @param     searchRecord               model.ZbUserSearchRecord
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateUserSearchRecord(searchRecord model.ZbUserSearchRecord) (err error) {
	err = global.GVA_DB.Create(&searchRecord).Error
	return err
}

// @title    DeleteUserSearchRecord
// @description   delete a ZbUserSearchRecord
// @auth                     （2020/04/05  20:22）
// @param     searchRecord               model.ZbUserSearchRecord
// @return                    error

func DeleteUserSearchRecord(searchRecord model.ZbUserSearchRecord) (err error) {
	err = global.GVA_DB.Delete(searchRecord).Error
	return err
}

// @title    DeleteUserSearchRecordByIds
// @description   delete UserSearchRecords
// @auth                     （2020/04/05  20:22）
// @param     searchRecord               model.ZbUserSearchRecord
// @return                    error

func DeleteUserSearchRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.ZbUserSearchRecord{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateUserSearchRecord
// @description   update a ZbUserSearchRecord
// @param     searchRecord          *model.ZbUserSearchRecord
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateUserSearchRecord(searchRecord *model.ZbUserSearchRecord) (err error) {
	err = global.GVA_DB.Save(searchRecord).Error
	return err
}

// @title    GetUserSearchRecord
// @description   get the info of a ZbUserSearchRecord
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    ZbUserSearchRecord        ZbUserSearchRecord

func GetUserSearchRecord(id uint) (err error, searchRecord model.ZbUserSearchRecord) {
	err = global.GVA_DB.Where("id = ?", id).First(&searchRecord).Error
	return
}

// @title    GetUserSearchRecordInfoList
// @description   get ZbUserSearchRecord list by pagination, 分页获取UserSearchRecord
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetUserSearchRecordInfoList(info request.UserSearchRecordSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.ZbUserSearchRecord{})
    var searchRecords []model.ZbUserSearchRecord
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&searchRecords).Error
	return err, searchRecords, total
}