package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateUserBrowseRecord
// @description   create a ZbUserBrowseRecord
// @param     browseRecord               model.ZbUserBrowseRecord
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateUserBrowseRecord(browseRecord model.ZbUserBrowseRecord) (err error) {
	err = global.GVA_DB.Create(&browseRecord).Error
	return err
}

// @title    DeleteUserBrowseRecord
// @description   delete a ZbUserBrowseRecord
// @auth                     （2020/04/05  20:22）
// @param     browseRecord               model.ZbUserBrowseRecord
// @return                    error

func DeleteUserBrowseRecord(browseRecord model.ZbUserBrowseRecord) (err error) {
	err = global.GVA_DB.Delete(browseRecord).Error
	return err
}

// @title    DeleteUserBrowseRecordByIds
// @description   delete UserBrowseRecords
// @auth                     （2020/04/05  20:22）
// @param     browseRecord               model.ZbUserBrowseRecord
// @return                    error

func DeleteUserBrowseRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.ZbUserBrowseRecord{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateUserBrowseRecord
// @description   update a ZbUserBrowseRecord
// @param     browseRecord          *model.ZbUserBrowseRecord
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateUserBrowseRecord(browseRecord *model.ZbUserBrowseRecord) (err error) {
	err = global.GVA_DB.Save(browseRecord).Error
	return err
}

// @title    GetUserBrowseRecord
// @description   get the info of a ZbUserBrowseRecord
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    ZbUserBrowseRecord        ZbUserBrowseRecord

func GetUserBrowseRecord(id uint) (err error, browseRecord model.ZbUserBrowseRecord) {
	err = global.GVA_DB.Where("id = ?", id).First(&browseRecord).Error
	return
}

// @title    GetUserBrowseRecordInfoList
// @description   get ZbUserBrowseRecord list by pagination, 分页获取UserBrowseRecord
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetUserBrowseRecordInfoList(info request.UserBrowseRecordSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.ZbUserBrowseRecord{})
    var browseRecords []model.ZbUserBrowseRecord
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&browseRecords).Error
	return err, browseRecords, total
}