package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateUserLabel
// @description   create a ZbUserLabel
// @param     userLabel               model.ZbUserLabel
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateUserLabel(userLabel model.ZbUserLabel) (err error) {
	err = global.GVA_DB.Create(&userLabel).Error
	return err
}

// @title    DeleteUserLabel
// @description   delete a ZbUserLabel
// @auth                     （2020/04/05  20:22）
// @param     userLabel               model.ZbUserLabel
// @return                    error

func DeleteUserLabel(userLabel model.ZbUserLabel) (err error) {
	err = global.GVA_DB.Delete(userLabel).Error
	return err
}

// @title    DeleteUserLabelByIds
// @description   delete UserLabels
// @auth                     （2020/04/05  20:22）
// @param     userLabel               model.ZbUserLabel
// @return                    error

func DeleteUserLabelByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.ZbUserLabel{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateUserLabel
// @description   update a ZbUserLabel
// @param     userLabel          *model.ZbUserLabel
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateUserLabel(userLabel *model.ZbUserLabel) (err error) {
	err = global.GVA_DB.Save(userLabel).Error
	return err
}

// @title    GetUserLabel
// @description   get the info of a ZbUserLabel
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    ZbUserLabel        ZbUserLabel

func GetUserLabel(id uint) (err error, userLabel model.ZbUserLabel) {
	err = global.GVA_DB.Where("id = ?", id).First(&userLabel).Error
	return
}

// @title    GetUserLabelInfoList
// @description   get ZbUserLabel list by pagination, 分页获取UserLabel
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetUserLabelInfoList(info request.UserLabelSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.ZbUserLabel{})
    var userLabels []model.ZbUserLabel
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&userLabels).Error
	return err, userLabels, total
}