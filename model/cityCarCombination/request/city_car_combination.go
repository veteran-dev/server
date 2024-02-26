package request

import (
	"time"

	"github.com/veteran-dev/server/model/common/request"
)

type CityCarCombinationSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	request.PageInfo
}
