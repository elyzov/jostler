package models

import (
  "github.com/astaxie/beego/orm"
  "github.com/astaxie/beego/validation"
)

type User struct {
  Id      int
  Login   string    `valid:"Required;MaxSize(32)"`
  Name    string    `valid:"MaxSize(64)"`
  Email   string    `valid:"Email;MaxSize(96)"`
  Pass    string    `valid:"Required;MinSize(8)" json:"-"`
  Ticks   []*Tick   `orm:"reverse(many)"` // reverse relationship of fk
}

func (u *User) TableName() string {
  return "users"
}

func (u *User) Valid(v *validation.Validation) {
  db := orm.NewOrm()
  if exist := db.QueryTable(u.TableName()).Filter("Login", u.Login).Exclude("Id", u.Id).Exist(); exist {
    // Set error messages of Name by SetError and HasErrors will return true
    v.SetError("Login", "Login already exists")
  }
  if u.Email != "" {
    if exist := db.QueryTable(u.TableName()).Filter("Email", u.Email).Exclude("Id", u.Id).Exist(); exist {
      v.SetError("Email", "Email already registered")
    }
  }
}


func init() {
  // Need to register model in init
  orm.RegisterModel(new(User))
}