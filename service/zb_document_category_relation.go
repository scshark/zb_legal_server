package service

import (
	"fmt"
	"gin-vue-admin/global"
	"strconv"
)

func CreateDocCategoryRelation(docId string, cids []int) (err error) {

	sql := "insert into zb_document_category_relation (category_id,document_id ) values "

	for k, cid := range cids {
		if k != 0 {
			sql += ","
		}
		sql += "(" + strconv.Itoa(int(cid)) + "," + docId + ")"
	}

	err = global.GVA_DB.Exec(sql).Error
	return err
}
