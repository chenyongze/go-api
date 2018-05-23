package main

import (
	"github.com/chenyongze/go-api/models"
	_ "github.com/chenyongze/go-api/routers"

	"github.com/astaxie/beego"
)

func main() {
	models.Init()
	beego.Run()
}
