package response

import (
	"github.com/5asp/gin-vue-admin/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
