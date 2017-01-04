package controllers

import (
	"fmt"
	"log"
	//	"memory/levenshtein"
	"memory/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//	"github.com/mozillazg/go-pinyin"
)

type MainController struct {
	beego.Controller
}

//var enc = simplifiedchinese.GBK

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
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
	//	for i := 0; i < len(poets); i++ {
	//		fmt.Println(poets[i].Content)
	//		arg := pinyin.NewArgs()
	//		arg.Style = pinyin.Initials
	//		fmt.Println(pinyin.Pinyin(poets[i].Content, arg))
	//	}

	//fmt.Println(o.Insert(user))*/
	//	f, err := os.Open("sample.txt")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	r := transform.NewReader(f, enc.NewDecoder())
	//	sc := bufio.NewScanner(r)
	//	dis := 10000
	//	poetry := ""
	//	for sc.Scan() {
	//		fmt.Println(sc.Text())
	//		strs := pinyin.LazyPinyin(sc.Text(), pinyin.NewArgs())
	//		var str string
	//		for i := 0; i < len(strs); i++ {
	//			str += strs[i]
	//		}
	//		distance := levenshtein.DistanceForStrings([]rune(password), []rune(str), levenshtein.DefaultOptions)
	//		if distance < dis {
	//			dis = distance
	//			poetry = sc.Text()
	//		}
	//		fmt.Printf(`Distance between "%s" and "%s" computed as %d\n`, password, str, distance)
	//	}
	//	if err = sc.Err(); err != nil {
	//		log.Fatal(err)
	//	}
	//	if err = f.Close(); err != nil {
	//		log.Fatal(err)
	//	}

	//	c.Data["Result"] = poetry
}
