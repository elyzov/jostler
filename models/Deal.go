package models

import (
  "time"
  "github.com/astaxie/beego/orm"
)

type Deal struct {
  Id        int
  Tick      *Tick     `orm:"column(tick_id);rel(fk)" json:"-"`     // RelForeignKey relation
  DateTime  time.Time `valid:"Required"`
  Cat       *Cat      `valid:"Required" orm:"column(cat_id);rel(one)"`
  Amount    float32   `valid:"Required"`
  Cy        *Cy       `valid:"Required" orm:"column(cy_id);rel(one)"`
  Comment   string    `valid:"MaxSize(128)"`
  aux       string    `json:"-"`
  Tags      []*Tag    `orm:"rel(m2m)"`
}


func init() {
  // Need to register model in init
  orm.RegisterModel(new(Deal))
}