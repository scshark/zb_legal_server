package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateDocumentCategory
// @description   create a ZbDocumentCategory
// @param     docCategory               model.ZbDocumentCategory
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateDocumentCategory(docCategory model.ZbDocumentCategory) (err error) {
	err = global.GVA_DB.Create(&docCategory).Error
	return err
}

// @title    DeleteDocumentCategory
// @description   delete a ZbDocumentCategory
// @auth                     （2020/04/05  20:22）
// @param     docCategory               model.ZbDocumentCategory
// @return                    error

func DeleteDocumentCategory(docCategory model.ZbDocumentCategory) (err error) {
	err = global.GVA_DB.Delete(docCategory).Error
	return err
}

// @title    DeleteDocumentCategoryByIds
// @description   delete DocumentCategorys
// @auth                     （2020/04/05  20:22）
// @param     docCategory               model.ZbDocumentCategory
// @return                    error

func DeleteDocumentCategoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.ZbDocumentCategory{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateDocumentCategory
// @description   update a ZbDocumentCategory
// @param     docCategory          *model.ZbDocumentCategory
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateDocumentCategory(docCategory *model.ZbDocumentCategory) (err error) {
	err = global.GVA_DB.Save(docCategory).Error
	return err
}

// @title    GetDocumentCategory
// @description   get the info of a ZbDocumentCategory
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    ZbDocumentCategory        ZbDocumentCategory

func GetDocumentCategory(id uint) (err error, docCategory model.ZbDocumentCategory) {
	err = global.GVA_DB.Where("id = ?", id).First(&docCategory).Error
	return
}

// @title    GetDocumentCategoryInfoList
// @description   get ZbDocumentCategory list by pagination, 分页获取DocumentCategory
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetDocumentCategoryInfoList(info request.DocumentCategorySearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.ZbDocumentCategory{})
    var docCategorys []model.ZbDocumentCategory
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&docCategorys).Error
	return err, docCategorys, total
}