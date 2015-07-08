package routers

import (
	"funny/controllers"
	"github.com/astaxie/beego"
	"os"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/tucao", &controllers.MainController{})
	// 附件处理
	os.Mkdir("attachment", os.ModePerm)
	beego.Router("/attachment/:all", &controllers.AttachController{})
}
