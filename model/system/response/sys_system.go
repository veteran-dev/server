package response

import "github.com/veteran-dev/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
