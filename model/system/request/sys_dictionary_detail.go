package request

import (
	"github.com/5asp/gin-vue-admin/server/model/common/request"
	"github.com/5asp/gin-vue-admin/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
