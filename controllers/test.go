package controllers

import "github.com/chenyongze/go-api/libs"

type TestController struct {
	BaseController
}

//   /test/g?size=80434
func (this *TestController) G() {

	size_ ,_:= this.GetFloat("size")
	size := libs.SizeFormat(size_)
	this.ajaxReturn("success",MSG_OK,size)

}
