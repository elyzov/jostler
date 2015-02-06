package controllers

import (
  "jostler/models"
  "github.com/astaxie/beego"
)

type NestPreparer interface {
    NestPrepare()
}

type baseController struct {
  beego.Controller
//  i18n.Locale
  user    models.User
  isLogin bool
}

func (this *baseController) Prepare() {
  // page start time
//  this.Data["PageStartTime"] = time.Now()

  if app, ok := this.AppController.(NestPreparer); ok {
    app.NestPrepare()
  }
}