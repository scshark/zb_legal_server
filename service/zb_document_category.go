package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"github.com/pkg/errors"
	"strconv"
)

// @title    CreateDocumentCategory
// @description   create a ZbDocumentCategory
// @param     docCategory               model.ZbDocumentCategory
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateDocumentCategory(docCategory model.ZbDocumentCategory) (err error) {

	var parentCategory model.ZbDocumentCategory
	if docCategory.ParentId != 0 {
		err := global.GVA_DB.Where("id = ?", docCategory.ParentId).First(&parentCategory).Error
		if err != nil {
			return errors.New("上级分类不存在")
		}
		docCategory.CategoryLevel = parentCategory.CategoryLevel + 1
		// 顶级分类ID
		if parentCategory.ParentId == 0 {
			docCategory.TopParentId = int(parentCategory.ID)
		} else {
			docCategory.TopParentId = parentCategory.TopParentId
		}

	}
	err = global.GVA_DB.Create(&docCategory).Error
	return err
}

// @title    DeleteDocumentCategory
// @description   delete a ZbDocumentCategory
// @auth                     （2020/04/05  20:22）
// @param     docCategory               model.ZbDocumentCategory
// @return                    error

func DeleteDocumentCategory(id uint) (err error) {
	var docCategory model.ZbDocumentCategory
	err = global.GVA_DB.Where("id = ?",id).Delete(&docCategory).Error
	return err
}

// @title    DeleteDocumentCategoryByIds
// @description   delete DocumentCategorys
// @auth                     （2020/04/05  20:22）
// @param     docCategory               model.ZbDocumentCategory
// @return                    error

func DeleteDocumentCategoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.ZbDocumentCategory{}, "id in (?)", ids.Ids).Error
	return err
}

// @title    UpdateDocumentCategory
// @description   update a ZbDocumentCategory
// @param     docCategory          *model.ZbDocumentCategory
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateDocumentCategory(docCategory *model.ZbDocumentCategory) (err error) {

	db := global.GVA_DB.Where("id = ?", docCategory.ID).Find(&model.ZbDocumentCategory{})

	updateCategory := make(map[string]interface{})
	updateCategory["title"] = docCategory.Title
	updateCategory["sort"] = docCategory.Sort
	updateCategory["hidden"] = docCategory.Hidden

	err = db.Update(updateCategory).Error

	global.GVA_LOG.Debug("文书分类修改err:%v", err)
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
	db := global.GVA_DB.Model(&model.ZbDocumentCategory{}).Where("parent_id = ?", 0)
	// 搜索条件
	if info.ZbDocumentCategory.Title != "" {
		db.Where("title = ?", info.ZbDocumentCategory.Title)
	}
	var docCategories []model.ZbDocumentCategory
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order("id DESC").Find(&docCategories).Error

	docAllCategories, err := getAllChildrenCategory(docCategories)

	return err, docAllCategories, total
}

func getAllChildrenCategory(c []model.ZbDocumentCategory) ([]model.ZbDocumentCategory, error) {

	var err error
	var topParentId []string
	for _, v := range c {
		topParentId = append(topParentId, strconv.Itoa(int(v.ID)))
	}
	var docChildrenCategories []model.ZbDocumentCategory

	err = global.GVA_DB.Where("top_parent_id in (?)", topParentId).Find(&docChildrenCategories).Error

	categoryTree := make(map[int][]model.ZbDocumentCategory)
	categoryTree[0] = c
	for _, v := range docChildrenCategories {
		categoryTree[v.ParentId] = append(categoryTree[v.ParentId], v)
	}

	for i := 0; i < len(c); i++ {
		err = getAllCategoryList(&c[i], categoryTree)
	}

	return c, err
}

//
func getAllCategoryList(category *model.ZbDocumentCategory, categoryTree map[int][]model.ZbDocumentCategory) (err error) {

	category.Children = categoryTree[int(category.ID)]

	for i := 0; i < len(category.Children); i++ {
		err = getAllCategoryList(&category.Children[i], categoryTree)
	}
	return err
}
