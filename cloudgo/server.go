package main

import (
	"flag"

	"github.com/astaxie/beego" //beego框架
)

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
