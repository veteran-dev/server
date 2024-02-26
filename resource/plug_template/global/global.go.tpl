package global

{{- if .HasGlobal }}

import "github.com/veteran-dev/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}