package models

import (
  "github.com/astaxie/beego/orm"
)

type Cy struct {
  Id        int
  Abbr      string    `valid:"Required;MaxSize(3)"`
  Logo      string    `valid:"Required;MaxSize(64)"`
}


func init() {
  // Need to register model in init
  orm.RegisterModel(new(Cy))
}