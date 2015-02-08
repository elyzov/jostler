package models

import (
  "time"
  "github.com/astaxie/beego/orm"
)

type Tick struct {
  Id        int
  User      *User     `orm:"column(user_id);rel(fk)" json:"-"`     // RelForeignKey relation
  Title     string    `valid:"Required;MaxSize(64)"`
  Rest      float32
  Created   time.Time
  Modified  time.Time
}


func init() {
  // Need to register model in init
  orm.RegisterModel(new(Tick))
}