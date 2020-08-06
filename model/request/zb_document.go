package request

import "gin-vue-admin/model"

type DocumentSearch struct{
    model.ZbDocument
    PageInfo
}