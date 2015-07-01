package routers

import (
	"github.com/astaxie/beego"
	"paincompiler.com/controllers"
)

func init() {
	beego.Router("/", &controllers.NarcissismController{})
}
