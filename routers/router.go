package routers

import (
    "go-web/controllers"
    "github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/first", &controllers.FirstController{})
}
