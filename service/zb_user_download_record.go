package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateUserDownloadRecord
// @description   create a ZbUserDownloadRecord
// @param     downloadRecord               model.ZbUserDownloadRecord
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateUserDownloadRecord(downloadRecord model.ZbUserDownloadRecord) (err error) {
	err = global.GVA_DB.Create(&downloadRecord).Error
	return err
}

// @title    DeleteUserDownloadRecord
// @description   delete a ZbUserDownloadRecord
// @auth                     （2020/04/05  20:22）
// @param     downloadRecord               model.ZbUserDownloadRecord
// @return                    error

func DeleteUserDownloadRecord(downloadRecord model.ZbUserDownloadRecord) (err error) {
	err = global.GVA_DB.Delete(downloadRecord).Error
	return err
}

// @title    DeleteUserDownloadRecordByIds
// @description   delete UserDownloadRecords
// @auth                     （2020/04/05  20:22）
// @param     downloadRecord               model.ZbUserDownloadRecord
// @return                    error

func DeleteUserDownloadRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.ZbUserDownloadRecord{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateUserDownloadRecord
// @description   update a ZbUserDownloadRecord
// @param     downloadRecord          *model.ZbUserDownloadRecord
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateUserDownloadRecord(downloadRecord *model.ZbUserDownloadRecord) (err error) {
	err = global.GVA_DB.Save(downloadRecord).Error
	return err
}

// @title    GetUserDownloadRecord
// @description   get the info of a ZbUserDownloadRecord
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    ZbUserDownloadRecord        ZbUserDownloadRecord

func GetUserDownloadRecord(id uint) (err error, downloadRecord model.ZbUserDownloadRecord) {
	err = global.GVA_DB.Where("id = ?", id).First(&downloadRecord).Error
	return
}

// @title    GetUserDownloadRecordInfoList
// @description   get ZbUserDownloadRecord list by pagination, 分页获取UserDownloadRecord
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetUserDownloadRecordInfoList(info request.UserDownloadRecordSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.ZbUserDownloadRecord{})
    var downloadRecords []model.ZbUserDownloadRecord
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&downloadRecords).Error
	return err, downloadRecords, total
}