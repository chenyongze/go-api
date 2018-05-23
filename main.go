package main

import (
	"models"
	_ "go-api/routers"

	"github.com/astaxie/beego"
)

func main() {
	models.Init()
	beego.Run()
}
