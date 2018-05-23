package controllers

import (
	"github.com/chanyipiaomiao/hltool"
	"github.com/astaxie/beego"
	"fmt"
	redis2 "github.com/garyburd/redigo/redis"
	"time"
)

type DingController struct {
	BaseController
}


type People struct {
	Name string
	Age int

}


func (self *DingController) Send()  {
	dingtalk := hltool.NewDingTalkClient(beego.AppConfig.String("ding.conf"), " 消息内容", "text")
	ok, err := hltool.SendMessage(dingtalk)
	if err != nil {
		self.ajaxReturn(ok,MSG_ERR,"")
	} else {
		self.ajaxReturn(err,MSG_OK,"")
	}
}

//@router /sh-redis
func (self *DingController) Redis()  {
	self.Ctx.ResponseWriter.WriteHeader(403)
	//account := self.GetString("account")
	c, err := redis2.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	//_, err = c.Do("SET", "mykey", "superWangx", "Ex", "5")
	//if err != nil {
	//	fmt.Println("redis set failed:", err)
	//}

	username, err := redis2.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}


	filters := make([]interface{}, 0)
	filters = append(filters, "status", "xxxxx")
	filters = append(filters, "status__inb", []int{0, 1, 2, 3, 4})
	filters = append(filters, "api_name__icontains","dddd")
	filters = append(filters, time.Now())

	u:=new(People)
	u.Name = "张三"
	u.Age = 78
	fmt.Println(u)

	//self.Ctx.WriteString(username)
	self.ajaxReturn(u.Name,MSG_OK,filters)

}