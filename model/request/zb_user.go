package request

import "gin-vue-admin/model"

type ZbUserSearch struct{
    model.ZbUser
    PageInfo
}