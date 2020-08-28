package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"strconv"
	"time"
)

// @title    CreateDocument
// @description   create a ZbDocument
// @param     doc               model.ZbDocument
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateDocument(doc model.ZbDocument) (err error) {
	// 标题 分类 标签 上传地址
	doc.ReleasedAt = time.Now()
	err = global.GVA_DB.Create(&doc).Error

	if len(doc.ClassId) > 0 {
		err = CreateDocCategoryRelation(strconv.Itoa(int(doc.ID)),doc.ClassId)
	}

	return err
}

// @title    DeleteDocument
// @description   delete a ZbDocument
// @auth                     （2020/04/05  20:22）
// @param     doc               model.ZbDocument
// @return                    error

func DeleteDocument(doc model.ZbDocument) (err error) {
	err = global.GVA_DB.Delete(doc).Error
	return err
}

// @title    DeleteDocumentByIds
// @description   delete Documents
// @auth                     （2020/04/05  20:22）
// @param     doc               model.ZbDocument
// @return                    error

func DeleteDocumentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.ZbDocument{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateDocument
// @description   update a ZbDocument
// @param     doc          *model.ZbDocument
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateDocument(doc *model.ZbDocument) (err error) {
	err = global.GVA_DB.Save(doc).Error
	return err
}

// @title    GetDocument
// @description   get the info of a ZbDocument
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    ZbDocument        ZbDocument

func GetDocument(id uint) (err error, doc model.ZbDocument) {
	err = global.GVA_DB.Where("id = ?", id).First(&doc).Error
	return
}

// @title    GetDocumentInfoList
// @description   get ZbDocument list by pagination, 分页获取Document
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetDocumentInfoList(info request.DocumentSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.ZbDocument{})
    var docs []model.ZbDocument
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&docs).Error
	return err, docs, total
}