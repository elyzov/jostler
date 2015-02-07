package main

import (
    "fmt"
	_ "jostler/routers"
	_ "github.com/lib/pq"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
    orm.RegisterDriver("postgres", orm.DR_Postgres)
    err := orm.RegisterDataBase("default", "postgres", "user=jostler password=password dbname=jostler sslmode=disable")
    if err != nil {
        fmt.Println(err)
    }
}

func main() {
 	if beego.RunMode == "dev" {
        beego.DirectoryIndex = true
        beego.StaticDir["/swagger"] = "swagger"
    }
    beego.Run()
}

