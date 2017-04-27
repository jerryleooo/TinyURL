package routers

import (
	"github.com/astaxie/beego"
	"tinyUrl/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
