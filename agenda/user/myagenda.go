package user

import (
	"fmt"
	 "encoding/json"
	 "io/ioutil"
)

var command string

type Users struct {
	Myusername string
	Mypassword string
	Myemail    string
	Myphone    string
}
type now struct {
	Myusername string
	Mypassword string
}

var p = make([]Users, 50)
var q =& now{}

/*func main() {

	command = "register"
	var a string = "huehua"
	var b string = "hahaha"
	var c string = "gggggg"
	var d string = "ffffff"
	if command == "register" {
		register(&a, &b, &c, &d)
	} else if command == "login" {
		login(&a, &b)
	} else if command == "logout" {
login(&a, &b)
		logout()
	} else if command == "query" {
login(&a, &b)
		query()
	} else if command == "delete" {
		login(&a, &b)
		delete()
	} else {
		fmt.Println("请输入正确的命令")
	}
}
*/

func Register(myusername *string, mypassword *string, myemail *string, myphone *string) {
	/*data, err := ioutil.ReadFile("curUser.txt")
	if err != nil {
		fmt.Println("read error")
		os.Exit(1)
	}
	fmt.Println(string(data))*/
	data1, _ := ioutil.ReadFile("curUser.txt")
	err := json.Unmarshal(data1,& p)
	if err != nil {
		fmt.Println("error:", err)
	}



	boss:=0
	j:=0
for j<50 {
	if *myusername==p[j].Myusername {
		fmt.Println("该用户已存在，请重新输入")
		boss++
	}
	j++
}
if boss==0{
	i:=0
  temp:=0
	if(p[0].Myusername==""){
		p[0].Myusername = *myusername
		p[0].Mypassword = *mypassword
		p[0].Myemail = *myemail
		p[0].Myphone=*myphone

		data, _ := json.Marshal(p)
		ioutil.WriteFile("curUser.txt", data, 0644)
		fmt.Println("注册成功")
	}else {
		for i<50{
			if p[i].Mypassword!=""{
				temp=i
			}
		 i++
		}
		p[temp+1].Myusername = *myusername
		p[temp+1].Mypassword = *mypassword
		p[temp+1].Myemail = *myemail
		p[temp+1].Myphone=*myphone
	//}
	//p[temp]

	data, _ := json.Marshal(p)
	ioutil.WriteFile("curUser.txt", data, 0644)
	fmt.Println("注册成功")
}
}
}

func Login(myusername *string, mypassword *string){
	data1, _ := ioutil.ReadFile("curUser.txt")
	err1 := json.Unmarshal(data1,& p)
	if err1 != nil {
		fmt.Println("error:", err1)
	}

	data2, _ := ioutil.ReadFile("load.txt")
	err2 := json.Unmarshal(data2,& q)
	if err2 != nil {
		fmt.Println("error:", err2)
	}

  i:=0
	boss:=0
	super:=0
	superme:=1
	for i<50{
		if q.Myusername==""{
			if *myusername==p[i].Myusername {
				if(*mypassword==p[i].Mypassword){
					q.Myusername=*myusername
					q.Mypassword=*mypassword
					data, _ := json.Marshal(q)
					ioutil.WriteFile("load.txt", data, 0644)
				//	fmt.Println("注册成功")
					//fmt.Println(q.Myusername)
				//	fmt.Println(q.Mypassword)
					fmt.Println("用户登录成功")
				}else{
					superme=0
				}
			}else{
				boss++
			}
		}else{
			super++
		}
		i++
	}
	if boss==50{
		fmt.Println("该用户未找到")
	}
	if super==50{
		fmt.Println("用户登录中")
	}
	if superme==0{
		fmt.Println("密码错误")
	}
}

func Logout(){
	data2, _ := ioutil.ReadFile("load.txt")
	err2 := json.Unmarshal(data2,& q)
	if err2 != nil {
		fmt.Println("error:", err2)
	}

		if q.Myusername=="" {
			fmt.Println("用户尚未登录")
		}else{
			q.Myusername=""
			q.Mypassword=""
			data, _ := json.Marshal(q)
			ioutil.WriteFile("load.txt", data, 0644)
			//fmt.Println("a"+q.Myusername)
			//fmt.Println(q.Mypassword)
			fmt.Println("登出成功")
		}
}

func Query(){
	data2, _ := ioutil.ReadFile("load.txt")
	err2 := json.Unmarshal(data2,& q)
	if err2 != nil {
		fmt.Println("error:", err2)
	}
	if q.Myusername=="" {
		fmt.Println("用户尚未登录")
	}else{
	data1, _ := ioutil.ReadFile("curUser.txt")
	err := json.Unmarshal(data1,& p)
	if err != nil {
		fmt.Println("error:", err)
	}
  i:=0
	for i<50{
		if p[i].Myusername!=""{
			fmt.Printf("用户名"+p[i].Myusername)
			fmt.Printf("邮箱"+p[i].Myemail)
			fmt.Println("手机号码"+p[i].Myphone)
		}
		i++
	}
}
}
func Delete(){
	data2, _ := ioutil.ReadFile("load.txt")
	err2 := json.Unmarshal(data2,& q)
	if err2 != nil {
		fmt.Println("error:", err2)
	}
	if q.Myusername=="" {
		fmt.Println("用户尚未登录")
	}else{
	data1, _ := ioutil.ReadFile("curUser.txt")
	err := json.Unmarshal(data1,& p)
	if err != nil {
		fmt.Println("error:", err)
	}

 i:=0
	for i<50{
		if(q.Myusername==p[i].Myusername){
			p[i].Myusername=""
			p[i].Mypassword=""
			p[i].Myemail=""
			p[i].Myphone=""
			q.Myusername=""
			q.Mypassword=""
			data, _ := json.Marshal(q)
			ioutil.WriteFile("load.txt", data, 0644)
		}
		i++
	}

	data, _ := json.Marshal(p)
	ioutil.WriteFile("curUser.txt", data, 0644)
	fmt.Println("删除成功")
}
}
