package bootstrap

import (
	"embed"
	"goblog/pkg/route"
	"goblog/pkg/routes"
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoute(staticFS embed.FS) *mux.Router {

	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)

	// 注册封装的一些路由方法
	route.SetRoute(router)

	// 加载静态资源
	sub, _ := fs.Sub(staticFS, "public")
	router.PathPrefix("/").Handler(http.FileServer(http.FS(sub)))

	return router
}
