package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer 创建一个新的服务器
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		Directory:  "assets",
		Extensions: []string{".html"},
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
			fmt.Println("Root:", root)
		}
	}
	// 功能1实现
	mx.Handle("/Function1", http.StripPrefix("/Function1", http.FileServer(http.Dir(webRoot+"/assets"))))
	// 功能2实现
	mx.HandleFunc("/Function2/test", apiTestHandler(formatter)).Methods("GET")
	// 功能3实现
	mx.HandleFunc("/Function3/login", postHandle(formatter))
	// 功能4实现
	mx.HandleFunc("/unknown", notImplementedHandler())
	// 默认界面
	mx.HandleFunc("/", indexHandler(formatter))
	mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/assets")))
}

func indexHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 渲染主页面
		formatter.HTML(w, http.StatusOK, "myindex", struct{}{})
	}
}

// 简单js访问功能
func apiTestHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			ID      string `json:"id"`
			Content string `json:"content"`
		}{ID: "FXXK", Content: "Welcome To Golang!"})
	}
}

// 根据表单显示表格
func postHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 处理表单
		r.ParseForm()
		un := r.FormValue("username")
		pw := r.FormValue("password")

		// 渲染表格
		formatter.HTML(w, http.StatusOK, "table", struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{Username: un, Password: pw})
	}
}

// 未实现页面
func notImplementedHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "501 not implement!\nNext Step: Look up gzip filter", http.StatusNotImplemented)
	}
}

func main() {
	m := NewServer()
	m.Run(":8000")
}
