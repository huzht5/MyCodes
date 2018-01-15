package service

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/bilibiliChangKai/GoLang/HM5/entities"

	"github.com/unrolled/render"
)

func initMydb(args []string) {
	if len(args) != 5 && len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Please input the database information!")
		fmt.Fprintln(os.Stderr, "\t./app username password port databasename")
		fmt.Fprintln(os.Stderr, "Or use: \n\t./app\nwe will use (root) (root) (2048) (test)")
		os.Exit(1)
	}

	// 声明四个变量
	name := "root"
	password := "root"
	port := "2048"
	dname := "test"

	if len(args) != 1 {
		name = args[1]
		password = args[2]
		port = args[3]
		dname = args[4]
	}

	// 创建数据库
	entities.InitMydb(name, password, port, dname)
}

func postUserInfoHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if len(req.Form["username"][0]) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Bad Input!"})
			return
		}
		u := entities.NewUserInfo(entities.UserInfo{UserName: req.Form["username"][0]})
		u.DepartName = req.Form["departname"][0]
		entities.UserInfoService.Save(u)
		formatter.JSON(w, http.StatusOK, u)
	}
}

func getUserInfoHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if len(req.Form["userid"][0]) != 0 {
			i, _ := strconv.ParseInt(req.Form["userid"][0], 10, 32)

			u := entities.UserInfoService.FindByID(int(i))
			formatter.JSON(w, http.StatusBadRequest, u)
			return
		}
		ulist := entities.UserInfoService.FindAll()
		formatter.JSON(w, http.StatusOK, ulist)
	}
}

func deleteUserInfoHandle(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if len(req.Form["userid"][0]) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Bad Input!"})
			return
		}

		i, _ := strconv.ParseInt(req.Form["userid"][0], 10, 32)
		u := entities.UserInfoService.DeleteByID(int(i))
		formatter.JSON(w, http.StatusOK, u)
	}
}
