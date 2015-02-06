package models

import (
  "github.com/astaxie/beego/orm"
)


type User struct {
  Id      int
  Login   string
  Name    string
  Email   string
  Pass    string
}


func init() {
  // Need to register model in init
  orm.RegisterModel(new(User))
}

