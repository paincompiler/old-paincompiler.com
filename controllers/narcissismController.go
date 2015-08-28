package controllers

import (
	"github.com/astaxie/beego"
	"paincompiler.com/models/narc"
)

type NarcissismController struct {
	beego.Controller
}

func (this *NarcissismController) Get() {
	output := narc.GetHexArray()
	this.Data["hex"] = string(output)
	this.Data["xeh"] = "It's the Narcissism Hex\nhttp://blog.paincompiler.us"
	this.TplNames = "index.tpl"
}
