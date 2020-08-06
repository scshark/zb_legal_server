package request

import "gin-vue-admin/model"

type UserBrowseRecordSearch struct{
    model.ZbUserBrowseRecord
    PageInfo
}