package controllers

import (
	"github.com/chanyipiaomiao/hltool"
	"github.com/astaxie/beego"
	"fmt"
	redisobj "github.com/garyburd/redigo/redis"
	ypclnt "github.com/yunpian/yunpian-go-sdk/sdk"
	"time"
	"github.com/astaxie/beego/logs"
)

type DingController struct {
	BaseController
}


type People struct {
	Name string
	Age int

}


func (self *DingController) Send()  {

	beego.SetLogger("file", `{"filename":"logs/test.log"}`)

	beego.Emergency("this is emergency")
	beego.Alert("this is alert")
	beego.Critical("this is critical")
	beego.Error("this is error")
	beego.Warning("this is warning")
	beego.Notice("this is notice")
	beego.Informational("this is informational")
	beego.Debug("this is debug")


	dingtalk := hltool.NewDingTalkClient(beego.AppConfig.String("ding.conf"), "##"+beego.AppConfig.String("testx"), "markdown")
	ok, err := hltool.SendMessage(dingtalk)
	if err != nil {
		self.ajaxReturn(ok,MSG_ERR,"")
	} else {
		self.ajaxReturn(err,MSG_OK,"")
	}
}

//@router /sh-redis
func (self *DingController) Redis()  {
	//self.Ctx.ResponseWriter.WriteHeader(403)
	//account := self.GetString("account")
	c, err := redisobj.Dial("tcp", beego.AppConfig.String("redis.link"))
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	c.Do("SELECT",0)
	c.Do("Auth",beego.AppConfig.String("redis.auth"))
	defer c.Close()

	//_, err = c.Do("SET", "mykey", "superWangx", "Ex", "50000000")
	//if err != nil {
	//	fmt.Println("redis set failed:", err)
	//}

	username, err := redisobj.String(c.Do("GET", "mykey"))
	if err != nil {
		beego.Info("redis get failed:", err)
	} else {
		beego.Info("Get mykey: %v \n", username)
	}

	u:=new(People)
	u.Name = username
	u.Age = 78

	filters := make([]interface{}, 0)
	//filters = append(filters, "status", "xxxxx")
	//filters = append(filters, "status__inb", []int{0, 1, 2, 3, 4})
	//filters = append(filters, "api_name__icontains","dddd")
	filters = append(filters, time.Now())
	filters = append(filters, u)

	fmt.Println(u)

	//self.Ctx.WriteString(username)
	self.ajaxReturn(u.Name,MSG_OK,filters)

}

//@router /redis_h [get]
func (self *DingController) Reids_hash_xinxi()  {

	xinxi_id := self.GetString("xinxi_id")
	client := createClient()

	val, err := client.HGetAll("xinxi:hash:" + xinxi_id ).Result()
	if err != nil {
		panic(err)
		self.ajaxReturn("error",MSG_ERR,val)
	}
	fmt.Println("key", val)
	logs.Info("M","logs::aaaaaaaaaaabbbb")
	beego.Info("ssssssss","beego::aaaaaaaaaaabbbb")
	self.ajaxReturn("success",MSG_OK,val)
}

//云片短信发送
func (self *DingController) Send_sms()  {

	clnt := ypclnt.New("bc90ddb0a06212de3eb22100d2f993a6")
	//clnt.Sign("【胜乐典藏】")
	param := ypclnt.NewParam(2)
	param[ypclnt.MOBILE] = "18514823010"
	//param[ypclnt.SIGN] = "【胜乐典藏】"
	param[ypclnt.TEXT] = "【胜乐典藏】您正在操作修改支付密码，验证码是345679。"
	r := clnt.Sms().SingleSend(param)

	filters := make([]interface{}, 0)
	filters = append(filters, r)

	beego.Info(r);
	self.ajaxReturn("success",MSG_OK,filters)

	
}