package models

import (
  "strconv"
  //"errors"
  // "reflect"
  // "github.com/astaxie/beego"
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

/* Implement Rest Interface */
func (this *UserModel) New() interface{} {
  return &User{}
}

func (this *UserModel) Add(ob interface{}) (string, error) {
  id, err := this.db.Insert(ob)
  return strconv.FormatInt(id, 10), err
}

func (this *UserModel) Update(user interface{}) error {
  _, err := this.db.Update(user)
  return err
}

func (this *UserModel) Delete(oid string) error {
   _, err := this.db.Delete(this.GetModel(oid))
   return err
}

func (this *UserModel) FindById(oid string) (interface{}, error) {
  ob := this.GetModel(oid)
  err := this.db.Read(ob)
  if err != nil {
    return nil, err
  }
  return ob, nil
}

func (this *UserModel) FindAll() (interface{}, error) {
  var users []*User
  _, err := this.db.QueryTable("users").All(&users)
  if err != nil {
    return nil, err
  }
  return users, nil
}

func (this *UserModel) Map(mt int, ob interface{}) map[string]interface{} {
  return map[string]interface{}{this.Tag(mt): ob}
}


/* Helper functions */
func (this *UserModel) GetModel(oid string) interface{} {
  id, _ := strconv.Atoi(oid)
  return &User{Id: id}
}

func (this *UserModel) Tag(mt int) string {
  switch mt {
    case SINGULAR: return "user"
    case PLURAL  : return "users"
  }
  return "unknown"
}












