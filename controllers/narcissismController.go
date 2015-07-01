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
	this.Data["xeh"] = "It's Narcissism Hex\nEmail:paincompiler@gmail.com"
	this.TplNames = "index.tpl"
}
