package controllers

import (
	"encoding/json"
	"jostler/models"
//	"github.com/astaxie/beego"
)


type UserController struct {
	baseController
}

// @Title create
// @Description create user
// @Param   body        body    models.User   true        "The user options"
// @Success 200 {string} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (this *UserController) AddUser() {
    var user models.User
    json.Unmarshal(this.Ctx.Input.RequestBody, &user)
    userid := "12"
    this.Data["json"] = map[string]string{"UserId": userid}
    this.ServeJson()
}

func (this *UserController) UserInfo() {
	id := this.Ctx.Input.Param(":id")
    this.Data["json"] = map[string]string{"UserId": id}
    this.ServeJson()
}

func (this *UserController) UserUpdate() {
//	id := this.Ctx.Input.Param(":id")
}

func (this *UserController) UserDelete() {
//	id := this.Ctx.Input.Param(":id")
}