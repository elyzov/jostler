package models

import (
  "time"
  "github.com/astaxie/beego/orm"
)

type DealType struct {
  Id        int
  Title     string    `valid:"Required;MaxSize(32)"`
}


func init() {
  // Need to register model in init
  orm.RegisterModel(new(DealType))
}