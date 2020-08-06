package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateDocumentKeyword
// @description   create a ZbDocumentKeyword
// @param     docKeyword               model.ZbDocumentKeyword
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateDocumentKeyword(docKeyword model.ZbDocumentKeyword) (err error) {
	err = global.GVA_DB.Create(&docKeyword).Error
	return err
}

// @title    DeleteDocumentKeyword
// @description   delete a ZbDocumentKeyword
// @auth                     （2020/04/05  20:22）
// @param     docKeyword               model.ZbDocumentKeyword
// @return                    error

func DeleteDocumentKeyword(docKeyword model.ZbDocumentKeyword) (err error) {
	err = global.GVA_DB.Delete(docKeyword).Error
	return err
}

// @title    DeleteDocumentKeywordByIds
// @description   delete DocumentKeywords
// @auth                     （2020/04/05  20:22）
// @param     docKeyword               model.ZbDocumentKeyword
// @return                    error

func DeleteDocumentKeywordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.ZbDocumentKeyword{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateDocumentKeyword
// @description   update a ZbDocumentKeyword
// @param     docKeyword          *model.ZbDocumentKeyword
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateDocumentKeyword(docKeyword *model.ZbDocumentKeyword) (err error) {
	err = global.GVA_DB.Save(docKeyword).Error
	return err
}

// @title    GetDocumentKeyword
// @description   get the info of a ZbDocumentKeyword
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    ZbDocumentKeyword        ZbDocumentKeyword

func GetDocumentKeyword(id uint) (err error, docKeyword model.ZbDocumentKeyword) {
	err = global.GVA_DB.Where("id = ?", id).First(&docKeyword).Error
	return
}

// @title    GetDocumentKeywordInfoList
// @description   get ZbDocumentKeyword list by pagination, 分页获取DocumentKeyword
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetDocumentKeywordInfoList(info request.DocumentKeywordSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.ZbDocumentKeyword{})
    var docKeywords []model.ZbDocumentKeyword
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&docKeywords).Error
	return err, docKeywords, total
}