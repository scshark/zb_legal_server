package request

import "gin-vue-admin/model"

type DocumentKeywordSearch struct{
    model.ZbDocumentKeyword
    PageInfo
}