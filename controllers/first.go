package controllers

import (
    "github.com/astaxie/beego"
)

type FirstController struct {
    beego.Controller
}

func (c *FirstController) Get() {
    c.Data["first"] = "Hello"
    c.Data["second"] = "World"
    c.TplName = "first.tpl"
}

