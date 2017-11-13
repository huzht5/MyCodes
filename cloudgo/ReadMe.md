cloudgo
=
hello word
--
**代码**

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

**实验截图**

![](https://github.com/15331109huzehua/golang/blob/master/cloudgo/images/%E6%8D%95%E8%8E%B71.PNG) 

server
--
**代码**

*用beego的框架函数来实现*

type MainController struct {

	beego.Controller //控制器
}

func (this *MainController) Get() {

	name := this.Ctx.Input.Param(":name")                       //获得路由信息
	this.Ctx.WriteString("Welcome to this page, " + name + "!") //写入
}

func main() {

	port := flag.String("port", "", "port:default is 8080") //传端口号
	flag.Parse()
	beego.Router("/cloudgo/:name", &MainController{}) //路由设置
	beego.Run(":" + *port)                            //运行
}

**实验截图**

![](https://github.com/15331109huzehua/golang/blob/master/cloudgo/images/%E6%8D%95%E8%8E%B72.PNG) 

**压力测试**

![](https://github.com/15331109huzehua/golang/blob/master/cloudgo/images/%E6%8D%95%E8%8E%B73.PNG)

结果：

![](https://github.com/15331109huzehua/golang/blob/master/cloudgo/images/%E6%8D%95%E8%8E%B74.PNG)
