package controllers

import (
	_ "fmt"
	"funny/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	var err error
	c.Data["TuCaos"], err = models.GetAll()
	c.Data["Num"] = models.Count()
	if err != nil {
		beego.Error(nil)
	}
	c.TplNames = "index.html"
}
func (c *MainController) Post() {
	tucao := c.Input().Get("tucaostr")
	err := models.AddTuCao(tucao)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/", 301)
	return

}
