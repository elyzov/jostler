package models

import (
  "fmt"
  "github.com/astaxie/beego/orm"
)

const (
  SINGULAR = iota
  PLURAL
)

type ResponseData map[string]interface{}

type Rest interface {
	Prepare()
	New() interface{}
	Add(model interface{}) (string, error)  // Has default implementation
	Update(model interface{}) error
	Delete(id string) error
	FindById(id string) (interface{}, error)
	FindAll() (interface{}, error)
	Map(int, interface{}) map[string]interface{}
  NotFound(id string) map[string]interface{}
}

type restModel struct {
  db      orm.Ormer
}

func (this *restModel) Prepare() {
  this.db = orm.NewOrm()
}

func (this *restModel) NotFound(id string) map[string]interface{} {
  return map[string]interface{}{ "id" : fmt.Sprintf("Id %s is not found", id) }
}