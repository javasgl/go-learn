package routers

import (
	"beego_test/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/ping", &controllers.MainController{})
}
