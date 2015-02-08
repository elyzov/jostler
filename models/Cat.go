package models

import (
  "github.com/astaxie/beego/orm"
)

type Cat struct {
  Id        int
  Title     string    `valid:"Required;MaxSize(64)"`
  Parent    *Cat      `orm:"column(parent_id);rel(fk)"`
  DataType  *DealType `orm:"column(data_type_id);rel(one)"`
}


func init() {
  // Need to register model in init
  orm.RegisterModel(new(Cat))
}