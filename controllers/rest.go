package controllers

import (
  "encoding/json"
  "jostler/models"
)

type restController struct {
  baseController
  model   models.Rest
}

func (this *restController) Post() {

  // var user models.User
  // json.Unmarshal(this.Ctx.Input.RequestBody, &user)
  // this.log.Trace("RequestBody = %v", this.Ctx.Input.RequestBody)
  // _, err := this.db.Insert(&user)
  // if err != nil {
  //   this.Response(ERROR, err)
  // } else {
  //   this.Response(SUCCESS, map[string]interface{}{"user": user})
  // }
  
  // this.ServeJson()

  m := this.model.NewModel()
  json.Unmarshal(this.Ctx.Input.RequestBody, &m)
  _, err := this.model.AddOne(&m)
  if err != nil {
    this.Response(ERROR, err)
  } else {
    this.Response(SUCCESS, this.model.MapOne(&m))
  }
  this.ServeJson()
}

func (this *restController) Get() {
  objectId := this.Ctx.Input.Params[":objectId"]
  if objectId != "" {
    om, err := this.model.FindOne(objectId)
    if err != nil {
      this.Response(ERROR, err)
    } else {
      this.Response(SUCCESS, this.model.MapOne(&om))
    }
  } else {
    oms, _ := this.model.FindAll()
    this.Response(SUCCESS, this.model.MapAll(&oms))
  }
  this.ServeJson()
}

func (this *restController) Put() {
  objectId := this.Ctx.Input.Params[":objectId"]
  m, err := this.model.FindOne(objectId)
  if err != nil {
    this.Response(ERROR, err)
  }
  json.Unmarshal(this.Ctx.Input.RequestBody, &m)

  err = this.model.Update(&m)
  if err != nil {
    this.Response(ERROR, err)
  } else {
    this.Response(SUCCESS, this.model.MapOne(&m))
  }
  this.ServeJson()
}

func (this *restController) Delete() {
  objectId := this.Ctx.Input.Params[":objectId"]
  err := this.model.Delete(objectId)
  st := SUCCESS
  if err != nil { st = ERROR }
  this.Response(st, err)
  this.ServeJson()
}