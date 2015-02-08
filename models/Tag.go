package models

import (
  "time"
  "github.com/astaxie/beego/orm"
)

type Tag struct {
  Id        int
  Title     string    `valid:"Required;MaxSize(48)"`
}


func init() {
  // Need to register model in init
  orm.RegisterModel(new(Tag))
}