package request

import (
	"github.com/veteran-dev/server/model/common/request"
	"github.com/veteran-dev/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
