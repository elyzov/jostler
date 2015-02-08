// @APIVersion 1.0.0
// @Title Jostler API
// @Description jostler rest api service
// @Contact lyzov.e.r@gmail.com
// @TermsOfServiceUrl http://jostler.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"jostler/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	api := beego.NewNamespace("/api/v1/",
		beego.NSAutoRouter(&controllers.AuthController{}),

        beego.NSRouter("/users", &controllers.UserController{}, "post:AddUser"),
        beego.NSRouter("/users/:id([0-9]+)", &controllers.UserController{}, "get:UserInfo;put:UserUpdate;delete:UserDelete"),
        beego.NSRouter("/ticks", &controllers.TickController{}, "post:AddTick;get:AllTicks"),
        beego.NSRouter("/ticks/:id([0-9]+)", &controllers.TickController{}, "get:TickInfo;put:TickUpdate;delete:TickDelete"),
    )
    beego.AddNamespace(api)
}
