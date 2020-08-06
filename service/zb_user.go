package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateZbUser
// @description   create a ZbUser
// @param     zbUser               model.ZbUser
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateZbUser(zbUser model.ZbUser) (err error) {
	err = global.GVA_DB.Create(&zbUser).Error
	return err
}

// @title    DeleteZbUser
// @description   delete a ZbUser
// @auth                     （2020/04/05  20:22）
// @param     zbUser               model.ZbUser
// @return                    error

func DeleteZbUser(zbUser model.ZbUser) (err error) {
	err = global.GVA_DB.Delete(zbUser).Error
	return err
}

// @title    DeleteZbUserByIds
// @description   delete ZbUsers
// @auth                     （2020/04/05  20:22）
// @param     zbUser               model.ZbUser
// @return                    error

func DeleteZbUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.ZbUser{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateZbUser
// @description   update a ZbUser
// @param     zbUser          *model.ZbUser
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateZbUser(zbUser *model.ZbUser) (err error) {
	err = global.GVA_DB.Save(zbUser).Error
	return err
}

// @title    GetZbUser
// @description   get the info of a ZbUser
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    ZbUser        ZbUser

func GetZbUser(id uint) (err error, zbUser model.ZbUser) {
	err = global.GVA_DB.Where("id = ?", id).First(&zbUser).Error
	return
}

// @title    GetZbUserInfoList
// @description   get ZbUser list by pagination, 分页获取ZbUser
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetZbUserInfoList(info request.ZbUserSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.ZbUser{})
    var zbUsers []model.ZbUser
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&zbUsers).Error
	return err, zbUsers, total
}