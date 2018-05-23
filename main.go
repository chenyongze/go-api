package main

import (
	_ "github.com/chenyongze/go-api/routers"

	"github.com/astaxie/beego"
	"github.com/chenyongze/go-api/models"
)

func main() {
	models.Init()
	beego.Run()
}
