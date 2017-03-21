package controllers

/*
import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
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
	radio := c.GetString("radiobutton")
	fmt.Println(radio)

	if radio == "initial" {
		arg := pinyin.NewArgs()
		arg.Style = pinyin.FirstLetter
		c.Data["Result"] = pinyin.Pinyin(password, arg)
		fmt.Println(c.Data["Result"])
	}

	var poets []models.Poetry
	o := orm.NewOrm()
	num, err := o.Raw("select * from poetries").QueryRows(&poets)
	//	num, err := o.QueryTable("user").All(&poets)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(num)
	//0-o
	//1-l
	//2-abc 3-def 4-ghi 5-jkl 6-mno 7-pqrs 8-tuv 9-wxyz
	m := InitMap()
	fmt.Println(m)
	rand.Seed(time.Now().UnixNano())
	arg := pinyin.NewArgs()
	arg.Style = pinyin.FirstLetter
	iniLetter := 0

	for v := 0; v < 100; v++ {

		var arr [6]int
		for i1 := 0; i1 < 6; i1++ {
			x := rand.Intn(10)
			//排除0和1
			if x == 0 || x == 1 {
				for {
					x = rand.Intn(10)
					if x != 0 && x != 1 {
						break
					}
				}
			}
			arr[i1] = x
		}
		//		fmt.Println(arr)
		containPass := false
		var letters string
		for j1 := 0; j1 < len(m[arr[0]]); j1++ {
			if containPass {
				break
			}
			str1 := m[arr[0]][j1]
			for k1 := 0; k1 < len(m[arr[1]]); k1++ {
				if containPass {
					break
				}
				str2 := m[arr[1]][k1]
				for l1 := 0; l1 < len(m[arr[2]]); l1++ {
					if containPass {
						break
					}
					str3 := m[arr[2]][l1]
					for n1 := 0; n1 < len(m[arr[3]]); n1++ {
						if containPass {
							break
						}
						str4 := m[arr[3]][n1]
						for n2 := 0; n2 < len(m[arr[4]]); n2++ {
							if containPass {
								break
							}
							str5 := m[arr[4]][n2]
							for n3 := 0; n3 < len(m[arr[5]]); n3++ {
								if containPass {
									break
								}
								str6 := m[arr[5]][n3]
								letters = str1 + str2 + str3 + str4 + str5 + str6
								//								fmt.Println(letters)

								for i := 0; i < int(num); i++ {
									if containPass {
										break
									}
									content := poets[i].Content
									contents := strings.Split(content, "。")
									for j := 0; j < len(contents); j++ {
										if containPass {
											break
										}
										strArr := strings.Split(contents[j], "，")
										//								str := ""

										for k := 0; k < len(strArr); k++ {
											str := strArr[k]
											//							fmt.Println(str)
											//							fmt.Println(pinyin.Pinyin(str, arg))
											pyArr := pinyin.Pinyin(str, arg)
											pyStr := ""
											for l := 0; l < len(pyArr); l++ {
												pyStr += pyArr[l][0]
											}
											//							fmt.Println(pyStr)
											if strings.Contains(pyStr, letters) {
												iniLetter += 1
												fmt.Println("contains", iniLetter)
												fmt.Println("contains", pyStr)
												fmt.Println("contains", str)
												containPass = true
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Println("xxx", iniLetter)
	fmt.Printf("首字母匹配的概率是：%0.8f\n", float64(iniLetter)/10)

	//	c.Data["Result"] = poetry
}

func InitMap() map[int][]string {
	m := make(map[int][]string)
	m[0] = []string{"o"}
	m[1] = []string{"l"}
	m[2] = []string{"a", "b", "c"}
	m[3] = []string{"d", "e", "f"}
	m[4] = []string{"g", "h", "i"}
	m[5] = []string{"j", "k", "l"}
	m[6] = []string{"m", "n", "o"}
	m[7] = []string{"p", "q", "r", "s"}
	m[8] = []string{"t", "u", "v"}
	m[9] = []string{"w", "x", "y", "z"}
	return m
}
*/
