package routers

import (
	"github.com/astaxie/beego"
	"github.com/qshuai/gcaptcha/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
