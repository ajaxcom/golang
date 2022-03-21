package view

import (
	"embed"
	"goblog/app/models/category"
	"goblog/app/models/user"
	"goblog/pkg/auth"
	"goblog/pkg/flash"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"html/template"
	"io"
	"io/fs"
	"strings"
)

type D map[string]interface{}

var TplFS embed.FS

// 简单布局 simple
func RenderSimple(w io.Writer, data D, tplFils ...string) {
	RenderTemplate(w, "simple", data, tplFils...)
}

// 通用布局 app
func Render(w io.Writer, data D, tplFils ...string) {
	RenderTemplate(w, "app", data, tplFils...)
}

// 为了解决多模板 与 指定模板的解析
// 增加状态
func RenderTemplate(w io.Writer, name string, data D, tplFils ...string) {

	data["User"], _ = user.All()
	data["isLogined"] = auth.Check()
	data["loginUser"] = auth.User()
	data["flash"] = flash.All()
	data["Users"], _ = user.All()
	data["Categories"], _ = category.All()

	// 生成模板文件
	allFiles := getTemplateFiles(tplFils...)

	tmpl, err := template.New("").Funcs(template.FuncMap{
		"RouteName2URL": route.Name2URL,
	}).ParseFS(TplFS, allFiles...)
	logger.LogError(err)

	err = tmpl.ExecuteTemplate(w, name, data)
	logger.LogError(err)
}

// 生成模板文件
func getTemplateFiles(tplFiles ...string) []string {
	viewDir := "resources/views/"

	for i, f := range tplFiles {
		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".gohtml"
	}

	file, err := fs.Glob(TplFS, viewDir+"layouts/*.gohtml")
	logger.LogError(err)

	return append(file, tplFiles...)
}
