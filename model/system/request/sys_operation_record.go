package request

import (
	"github.com/5asp/gin-vue-admin/server/model/common/request"
	"github.com/5asp/gin-vue-admin/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
