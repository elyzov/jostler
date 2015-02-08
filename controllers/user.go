package controllers

import (
  "fmt"
  "strconv"
  "encoding/json"
  "jostler/models"
  "github.com/astaxie/beego/validation"
)


type UserController struct {
  baseController
}

func (this *UserController) AddUser() {
  var user models.User
  err := json.Unmarshal(this.Ctx.Input.RequestBody, &user)
  if err != nil {
    this.Response(ERROR, fmt.Sprintf("Incorrect JSON: %v", err))
    return
  }
  valid := validation.Validation{}
  if ok, _ := valid.Valid(&user); !ok {
    inv := make(map[string]string)
    for _, err := range valid.Errors {
        inv[err.Key] =  err.Message
    }
    this.Response(FAIL, inv)
    return
  }
  _, err = this.db.Insert(&user)
  if err != nil {
    this.Response(ERROR, err)
  } else {
    this.Response(SUCCESS, map[string]interface{}{"user": user})
  }
}

func (this *UserController) UserInfo() {
  id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
  if err != nil {
    this.Response(FAIL, map[string]interface{}{"id": "Must be an integer"})
    return
  }
  user := models.User{Id: id}
  err = this.db.Read(&user)
  if err != nil {
    this.Response(FAIL, map[string]interface{}{"id": fmt.Sprintf("User with id %v not found", id)})
  } else {
    this.Response(SUCCESS, map[string]interface{}{"user": user})
  }
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
}

func (this *UserController) UserDelete() {
  id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
  if _, err := this.db.Delete(&models.User{Id: id}); err == nil {
    this.Response(SUCCESS, nil)
  } else {
    this.Response(ERROR, err)
  }
}