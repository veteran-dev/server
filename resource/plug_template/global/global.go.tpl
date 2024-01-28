package global

{{- if .HasGlobal }}

import "github.com/5asp/gin-vue-admin/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}