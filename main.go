package main

import (
	"memory/models"
	_ "memory/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/tang_poetry?charset=utf8")
	orm.RegisterModel(new(models.Poetry))

}

func main() {
	beego.Run()
}
