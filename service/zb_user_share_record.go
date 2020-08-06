package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateUserShareRecord
// @description   create a ZbUserShareRecord
// @param     shareRecord               model.ZbUserShareRecord
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateUserShareRecord(shareRecord model.ZbUserShareRecord) (err error) {
	err = global.GVA_DB.Create(&shareRecord).Error
	return err
}

// @title    DeleteUserShareRecord
// @description   delete a ZbUserShareRecord
// @auth                     （2020/04/05  20:22）
// @param     shareRecord               model.ZbUserShareRecord
// @return                    error

func DeleteUserShareRecord(shareRecord model.ZbUserShareRecord) (err error) {
	err = global.GVA_DB.Delete(shareRecord).Error
	return err
}

// @title    DeleteUserShareRecordByIds
// @description   delete UserShareRecords
// @auth                     （2020/04/05  20:22）
// @param     shareRecord               model.ZbUserShareRecord
// @return                    error

func DeleteUserShareRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.ZbUserShareRecord{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateUserShareRecord
// @description   update a ZbUserShareRecord
// @param     shareRecord          *model.ZbUserShareRecord
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateUserShareRecord(shareRecord *model.ZbUserShareRecord) (err error) {
	err = global.GVA_DB.Save(shareRecord).Error
	return err
}

// @title    GetUserShareRecord
// @description   get the info of a ZbUserShareRecord
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    ZbUserShareRecord        ZbUserShareRecord

func GetUserShareRecord(id uint) (err error, shareRecord model.ZbUserShareRecord) {
	err = global.GVA_DB.Where("id = ?", id).First(&shareRecord).Error
	return
}

// @title    GetUserShareRecordInfoList
// @description   get ZbUserShareRecord list by pagination, 分页获取UserShareRecord
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetUserShareRecordInfoList(info request.UserShareRecordSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.ZbUserShareRecord{})
    var shareRecords []model.ZbUserShareRecord
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&shareRecords).Error
	return err, shareRecords, total
}