package controllers

import (
  "strconv"
  "encoding/json"
  "jostler/models"
)


type UserController struct {
  baseController
}

func (this *UserController) NestPrepare() {
  // this.model = &models.UserModel{}
  // this.model.Prepare()
}


func (this *UserController) AddUser() {
  var user models.User
  json.Unmarshal(this.Ctx.Input.RequestBody, &user)
  _, err := this.db.Insert(&user)
  if err != nil {
    this.Response(ERROR, err)
  } else {
    this.Response(SUCCESS, map[string]interface{}{"user": user})
  }
  
  this.ServeJson()
}

func (this *UserController) UserInfo() {
  id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
  user := models.User{Id: id}
  err := this.db.Read(&user)
  if err != nil {
    this.Response(ERROR, err)
  } else {
    this.Response(SUCCESS, map[string]interface{}{"user": user})
  }
  this.ServeJson()
}

func (this *UserController) UserUpdate() {
  id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
  user := models.User{Id: id}
  err := this.db.Read(&user)
  if err != nil {
    this.Response(ERROR, err)
  }
  json.Unmarshal(this.Ctx.Input.RequestBody, &user)
  if _, err = this.db.Update(&user); err == nil {
    this.Response(SUCCESS, map[string]interface{}{"user": user})
  } else {
    this.Response(ERROR, err)
  }
  
  this.ServeJson()
}

func (this *UserController) UserDelete() {
  id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
  if _, err := this.db.Delete(&models.User{Id: id}); err == nil {
    this.Response(SUCCESS, nil)
  } else {
    this.Response(ERROR, err)
  }
  this.ServeJson()
}