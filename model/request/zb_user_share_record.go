package request

import "gin-vue-admin/model"

type UserShareRecordSearch struct{
    model.ZbUserShareRecord
    PageInfo
}