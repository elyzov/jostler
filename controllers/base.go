package controllers

import (
  "jostler/models"
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/orm"
  "github.com/astaxie/beego/validation"
)

const (
  SUCCESS = "success"
  FAIL    = "fail"
  ERROR   = "error"
)

type NestPreparer interface {
    NestPrepare()
}

type baseController struct {
  beego.Controller
//  i18n.Locale
  user    *models.User
  isLogin bool
  db      orm.Ormer

}

type JsonMap map[string]interface{}

func (this *baseController) Prepare() {
  // page start time
//  this.Data["PageStartTime"] = time.Now()
  this.db = orm.NewOrm()
  this.user = &models.User{Id : 1}
  this.db.Read(this.user)

  if app, ok := this.AppController.(NestPreparer); ok {
    app.NestPrepare()
  }
}

func (this *baseController) Validate(ob interface{}) map[string]string {
  valid := validation.Validation{}
  if ok, err := valid.Valid(ob); !ok {
    if err != nil {
      beego.Trace("Validation error: ", err)
    }
    inv := make(map[string]string)
    for _, err := range valid.Errors {
      inv[err.Key] =  err.Message
    }
    return inv
  }
  return nil
}

func (this *baseController) Response(status string, data interface{}) {
  if status == ERROR {
    if err, ok := data.(error); ok {
      data = err.Error()
    }
    this.Data["json"] = map[string]interface{}{
      "status"  : status,
      "message" : data,
    }
  } else {
    this.Data["json"] = map[string]interface{}{
      "status" : status,
      "data"   : data,
    }
  }
  this.ServeJson()
}