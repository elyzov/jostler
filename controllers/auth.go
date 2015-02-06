package controllers

import (
  "encoding/json"
  "jostler/models"
//  "github.com/astaxie/beego"
)


type AuthController struct {
  baseController
}

func (this *AuthController) Login() {

  var user models.User
  json.Unmarshal(this.Ctx.Input.RequestBody, &user)
  userid := "12"
  this.Data["json"] = map[string]string{"UserId": userid}
  this.ServeJson()
}

func (this *AuthController) Logout() {

}