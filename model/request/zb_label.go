package request

import "gin-vue-admin/model"

type LabelSearch struct{
    model.ZbLabel
    PageInfo
}