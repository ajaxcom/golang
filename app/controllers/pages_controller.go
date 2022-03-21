package controllers

import (
	"fmt"
	"net/http"
)

type PageController struct {
}

func (*PageController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, 欢迎1</h1>")
}

func (*PageController) About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, 欢迎2</h1>")
}
func (*PageController) NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, 欢迎3</h1>")
}
