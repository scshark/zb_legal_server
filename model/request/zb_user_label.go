package request

import "gin-vue-admin/model"

type UserLabelSearch struct{
    model.ZbUserLabel
    PageInfo
}