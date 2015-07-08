package controllers

import (
	// "io"
	// "net/url"
	// "os"

	"github.com/astaxie/beego"
)

type AttachController struct {
	beego.Controller
}

func (this *AttachController) Get() {
	this.Ctx.Output.Download("attachment/tucao.zip")
}
