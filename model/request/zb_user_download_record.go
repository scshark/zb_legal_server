package request

import "gin-vue-admin/model"

type UserDownloadRecordSearch struct{
    model.ZbUserDownloadRecord
    PageInfo
}