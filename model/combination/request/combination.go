package request

import (
	"time"

	"github.com/5asp/gin-vue-admin/server/model/common/request"
)

type CombinationSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	request.PageInfo
}
