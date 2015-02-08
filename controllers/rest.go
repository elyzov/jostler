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
  m := this.model.New()
  err := json.Unmarshal(this.Ctx.Input.RequestBody, &m)
  if err != nil {
    this.Response(ERROR, err)
    return
  }
  _, err = this.model.Add(m)
  if err != nil {
    this.Response(ERROR, err)
  } else {
    this.Response(SUCCESS, this.model.Map(models.SINGULAR, &m))
  }
}

func (this *restController) Get() {
  objectId := this.Ctx.Input.Params[":objectId"]
  if objectId != "" {
    om, err := this.model.FindById(objectId)
    if err != nil {
      this.Response(FAIL, this.model.NotFound(objectId))
    } else {
      this.Response(SUCCESS, this.model.Map(models.SINGULAR, &om))
    }
  } else {
    oms, _ := this.model.FindAll()
    this.Response(SUCCESS, this.model.Map(models.PLURAL, &oms))
  }
}

func (this *restController) Put() {
  objectId := this.Ctx.Input.Params[":objectId"]
  m, err := this.model.FindById(objectId)
  if err != nil {
    this.Response(ERROR, err)
    return
  }
  err = json.Unmarshal(this.Ctx.Input.RequestBody, &m)
  if err != nil {
    this.Response(ERROR, err)
    return
  }

  err = this.model.Update(m)
  if err != nil {
    this.Response(ERROR, err)
  } else {
    this.Response(SUCCESS, this.model.Map(models.SINGULAR, &m))
  }
}

func (this *restController) Delete() {
  objectId := this.Ctx.Input.Params[":objectId"]
  err := this.model.Delete(objectId)
  st := SUCCESS
  if err != nil { st = ERROR }
  this.Response(st, err)
}