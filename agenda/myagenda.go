package main

import (
	"fmt"
)

var command string

type Users struct {
	Myusername string
	Mypassword string
	Myemail    string
	Myphone    string
}

var p = make([]Users, 10)

func main() {

	command = "register"
	var a string = "huzehua"
	var b string = "hahaha"
	var c string = "gggggg"
	var d string = "ffffff"
	if command == "register" {
		register(&a, &b, &c, &d)
		register(&a, &b, &c, &d)
	} else if command == "login" {
		fmt.Println("2")
	} else if command == "logout" {
		fmt.Println("3")
	} else if command == "query" {
		fmt.Println("4")
	} else if command == "delete" {
		fmt.Println("5")
	} else {
		fmt.Println("请输入正确的命令")
	}
}

var temp = 7

func register(myusername *string, mypassword *string, myemail *string, myphone *string) {
	/*data, err := ioutil.ReadFile("curUser.txt")
	if err != nil {
		fmt.Println("read error")
		os.Exit(1)
	}
	fmt.Println(string(data))*/
	p[0].Myusername = *myusername
	p[0].Mypassword = *mypassword
	p[0].Myemail = *myemail
	/*i := 0
	for i < 10 {
		if p[i].Myusername != "" {
			temp = i
			fmt.Println(temp)
		}
	}*/
	//p[temp]
	/*p.Myusername = *myusername
	p.Mypassword = *mypassword
	p.Myemail = *myemail
	p.Myphone = *myphone
	data, _ := json.Marshal(p)
	ioutil.WriteFile("curUser.txt", data, 0644)*/

	/*q := ioutil.ReadFile("curUser.txt")
	err := json.Unmarshal(data, &q)
	if err != nil {
		fmt.Println("error:", err)
	}*/
	//fmt.Println(string(data))
	//fmt.Printf(q.Myusername)
}
