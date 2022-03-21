package bootstrap

import (
	"embed"
	"goblog/pkg/view"
)

func SetupTemplate(tmplFs embed.FS) {
	view.TplFS = tmplFs
}
