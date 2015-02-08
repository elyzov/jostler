package models

import (
    "github.com/astaxie/beego/orm"
)

type Rest interface {
	Prepare()
	NewModel() interface{}
	AddOne(model interface{}) (string, error)
	FindOne(id string) (interface{}, error)
	MapOne(m interface{}) map[string]interface{}
	FindAll() (interface{}, error)
	MapAll(m interface{}) map[string]interface{}
	Update(model interface{}) error
	Delete(id string) error
}

type restModel struct {
  db      orm.Ormer
}

func (this *restModel) Prepare() {
  this.db = orm.NewOrm()
}