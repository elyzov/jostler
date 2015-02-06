package main

import (
	_ "jostler/routers"

	_ "github.com/lib/pq"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
    orm.RegisterDriver("postgres", orm.DR_Postgres)
    orm.RegisterDataBase("default", "postgres", "jostler:password@/jostler?charset=utf8")
}

func main() {
 	if beego.RunMode == "dev" {
        beego.DirectoryIndex = true
        beego.StaticDir["/swagger"] = "swagger"
    }
    beego.Run()
}

