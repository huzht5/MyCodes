package main

import (
    "fmt"
    "net/http"
    "strings"
    "log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()  //解析参数
    fmt.Println(r.Form)  //打印信息
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello world!") //写入到w
}

func main() {
    http.HandleFunc("/", sayhelloName)       //访问路由
    err := http.ListenAndServe(":9090", nil) //监听端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

