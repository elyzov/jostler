package models

import (
  "strconv"
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/orm"
)

type UserModel struct {
  restModel
}

type User struct {
  Id      int
  Login   string
  Name    string
  Email   string
  Pass    string
}

func (u *User) TableName() string {
  return "users"
}

func init() {
  // Need to register model in init
  orm.RegisterModel(new(User))
}

func (this *UserModel) NewModel() interface{} {
  return &User{}
}

func (this *UserModel) AddOne(user interface{}) (string, error) {
  id, err := this.db.Insert(&user)
  return strconv.FormatInt(id, 10), err
}

func (this *UserModel) FindOne(oid string) (interface{}, error) {
  id, _ := strconv.Atoi(oid)
  user := User{Id: id}
  beego.Trace("db = %v", this.db)
  err := this.db.Read(&user)
  if err != nil {
    return nil, err
  }
  return user, nil
}

func (this *UserModel) MapOne(user interface{}) map[string]interface{} {
  return map[string]interface{}{"user": user}
}

func (this *UserModel) FindAll() (interface{}, error) {
  var users []*User
  _, err := this.db.QueryTable("users").All(&users)
  if err != nil {
    return nil, err
  }
  return users, nil
}

func (this *UserModel) MapAll(users interface{}) map[string]interface{} {
  return map[string]interface{}{"users": users}
}


func (this *UserModel) Update(user interface{}) error {
  _, err := this.db.Update(&user)
  return err
}

func (this *UserModel) Delete(oid string) error {
  id, _ := strconv.Atoi(oid)
   _, err := this.db.Delete(&User{Id: id})
   return err
}
