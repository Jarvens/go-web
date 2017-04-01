package controllers

import (
    "github.com/astaxie/beego"
)

type MainController struct {
    beego.Controller
}

func (c *MainController) Get() {
    c.Data["Website"] = "beego.me"
    c.Data["Email"] = "wangbinbin8907@gmail.com"
    c.TplName = "index.tpl"
}
