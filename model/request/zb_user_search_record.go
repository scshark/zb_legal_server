package request

import "gin-vue-admin/model"

type UserSearchRecordSearch struct{
    model.ZbUserSearchRecord
    PageInfo
}