package controllers

import (
	"fmt"
	"log"
	"strings"
	//	"memory/levenshtein"
	"memory/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/mozillazg/go-pinyin"
)

type MainController struct {
	beego.Controller
}

//var enc = simplifiedchinese.GBK

func (c *MainController) Get() {
	c.Data["Website"] = "pku.edu.cn"
	c.Data["Email"] = "wemeya@163.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {
	c.Data["Website"] = "pku.edu.cn"
	c.Data["Email"] = "wemeya@163.com"
	c.TplName = "index.tpl"

	password := c.GetString("password")
	fmt.Println(password)

	var poets []models.Poetry
	o := orm.NewOrm()
	num, err := o.Raw("select * from poetries").QueryRows(&poets)
	//	num, err := o.QueryTable("user").All(&poets)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(num)
	for i := 0; i < num; i++ {

		content := poets[i].Content
		contents := strings.Split(content, "。")
		for j := 0; j < len(contents); j++ {
			contentEach := strings.Split(contents[j], "，")
			for k := 0; k < len(contentEach); k++ {
				arg := pinyin.NewArgs()
				arg.Style = pinyin.Initials
				fmt.Println(pinyin.Pinyin(contentEach[k], arg))
			}

		}

	}

	//	c.Data["Result"] = poetry
}
