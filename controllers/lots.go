package controllers

import (
	"fmt"
	"github.com/chenyongze/go-api/models"
)

type LotsController struct {
	BaseController
}

//测试接口
func (self *LotsController) Todo()  {
	data := map[string]int{"key1": 5, "key2": 4}

	//delete(data ,"key1")
	//delete(data ,"key2")

	for k,v := range data{
		fmt.Println(k,v)
		if k == "key2"{
			data[k] = 7000
		}
		fmt.Print("backend::")
		fmt.Println(k,data[k])
	}
	//todo
	//data := [...]byte{'A', 'B', '3','b',5}
	//self.Ctx.WriteString("todo ...")

	//models.LotsGetById(15)

	self.ajaxReturn("xxx", MSG_OK,data)
}

//获取拍品信息
func (self *LotsController) Getlots()  {
	lot_id, _ := self.GetInt("lot_id", 0)
	fmt.Println(lot_id)
	data := make(map[string]interface{})
	lots , err := models.LotsGetById(lot_id)
	fmt.Println(lots)
	code := 0
	msg := "success"
	if lots != nil{
		fmt.Println(len(lots))
		list := make([]map[string]interface{}, len(lots))
		for k,lotinfo :=range lots{
			row := make(map[string]interface{})
			row["lot_sn"] = lotinfo.Lot_sn
			row["lot_name"] = lotinfo.Lot_name
			list[k] = row
		}


		StatusText := make(map[int]string)
		StatusText[0] = "<font color='red'>禁用</font>"
		StatusText[1] = "正常"

		data["list"] = list
		data["Y"] = lots
		data["error_no"] = err
		data["StatusText"] = StatusText
		data["array_test"] = [...]int{3,4,5}
		code = 0
	}else{
		//msg := "error"
		code = MSG_ERR
	}

	self.ajaxReturn(msg,code,data)
}

//http://sldc-go.io/lots/lotinfo?lot_id=1509
func (this *LotsController) LotInfo()  {
	lot_id := this.GetString("lot_id")
	redis :=createClient();
	val, err := redis.HGetAll("personage_auction:hash:325:lotid:" + lot_id ).Result()
	if err != nil {
		//panic(err)
		this.ajaxReturn("error",MSG_ERR,val)
	}
	fmt.Println("key", val)
	this.ajaxReturn("success",MSG_OK,val)
}


func (this *LotsController) Th() {
	redis :=createClient();
	val, err := redis.Keys("auction:hash:*" ).Result()
	if err != nil {
		//panic(err)
		this.ajaxReturn("error",MSG_ERR,val)
	}
	//fmt.Println("key", val)
	this.ajaxReturn("success",MSG_OK,val)
}

