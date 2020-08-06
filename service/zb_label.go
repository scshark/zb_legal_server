package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateLabel
// @description   create a ZbLabel
// @param     label               model.ZbLabel
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateLabel(label model.ZbLabel) (err error) {
	err = global.GVA_DB.Create(&label).Error
	return err
}

// @title    DeleteLabel
// @description   delete a ZbLabel
// @auth                     （2020/04/05  20:22）
// @param     label               model.ZbLabel
// @return                    error

func DeleteLabel(label model.ZbLabel) (err error) {
	err = global.GVA_DB.Delete(label).Error
	return err
}

// @title    DeleteLabelByIds
// @description   delete Labels
// @auth                     （2020/04/05  20:22）
// @param     label               model.ZbLabel
// @return                    error

func DeleteLabelByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.ZbLabel{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateLabel
// @description   update a ZbLabel
// @param     label          *model.ZbLabel
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateLabel(label *model.ZbLabel) (err error) {
	err = global.GVA_DB.Save(label).Error
	return err
}

// @title    GetLabel
// @description   get the info of a ZbLabel
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    ZbLabel        ZbLabel

func GetLabel(id uint) (err error, label model.ZbLabel) {
	err = global.GVA_DB.Where("id = ?", id).First(&label).Error
	return
}

// @title    GetLabelInfoList
// @description   get ZbLabel list by pagination, 分页获取Label
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetLabelInfoList(info request.LabelSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.ZbLabel{})
    var labels []model.ZbLabel
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&labels).Error
	return err, labels, total
}