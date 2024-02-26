package request

import (
	"github.com/veteran-dev/server/model/common/request"
	"github.com/veteran-dev/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
