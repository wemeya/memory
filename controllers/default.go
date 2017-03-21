package controllers

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"strings"

	"memory/levenshtein"
	"memory/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/mozillazg/go-pinyin"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Model() {
	c.Data["Website"] = "pku.edu.cn"
	c.Data["Email"] = "wemeya@163.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Get() {
	c.Model()
	Test()
}

func (c *MainController) Post() {
	c.Model()

	pin := c.GetString("pin")
	fmt.Println(pin)
	pattern := c.GetString("pattern")
	fmt.Println(pattern)
	strategy := c.GetString("strategy")
	fmt.Println(strategy)

	fmt.Println(len(pin))
	//错误检查
	if !(len(pin) == 4 || len(pin) == 6) {
		fmt.Println("PIN长度不合法，请重新输入！")
		return
	}
	var arr [6]int
	pinArray := []byte(pin)
	for i := 0; i < len(pinArray); i++ {
		arr[i] = int(pinArray[i] - 48)
		if arr[i] < 0 || arr[i] > 9 {
			fmt.Println("PIN中存在非法字符，请重新输入！")
			return
		}
	}
	fmt.Println(arr)

	//黑名单检测
	if BlackList(pin) {
		fmt.Println("PIN为弱口令，请重新输入！")
		return
	}

	var mapMode int
	if pattern == "nomapped" {
		mapMode = 0
	} else {
		mapMode = 1
	}
	m := InitMap(mapMode)
	fmt.Println(m)

	distance := 0
	arg := pinyin.NewArgs()

	if strategy == "initial" {
		arg.Style = pinyin.FirstLetter
	} else if strategy == "all" {
		arg.Style = pinyin.Normal
	} else {
		//编辑距离先设定为1
		arg.Style = pinyin.FirstLetter
		distance = 1
	}

	var poets []models.Poetry
	o := orm.NewOrm()
	num, err := o.Raw("select * from poetries").QueryRows(&poets)
	//	num, err := o.QueryTable("user").All(&poets)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(num)

	containPass := false
	var poetry, letters string

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
							fmt.Println(letters)

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
									for k := 0; k < len(strArr); k++ {
										str := strArr[k]
										pyArr := pinyin.Pinyin(str, arg)
										pyStr := ""
										for l := 0; l < len(pyArr); l++ {
											pyStr += pyArr[l][0]
										}
										if distance == 0 {
											if strings.Contains(pyStr, letters) {
												containPass = true
												poetry = str
											}
										} else {
											dis := levenshtein.DistanceForStrings([]rune(letters), []rune(pyStr), levenshtein.DefaultOptions)
											if dis == distance {
												poetry = str + "\n"
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
	fmt.Println("result", poetry)
	c.Data["Result"] = poetry
}

func InitMap(mode int) map[int][]string {
	m := make(map[int][]string)
	if mode == 1 {
		m[0] = []string{"o"}
		m[1] = []string{"l"}
	}
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

func BlackList(pin string) bool {
	var blackList = []string{"0000", "1111", "1234", "000000", "111111", "123456"}
	for _, v := range blackList {
		if v == pin {
			return true
		}
	}
	return false
}

func Test() {
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
	fmt.Println(m)
	rand.Seed(time.Now().UnixNano())
	arg := pinyin.NewArgs()
	arg.Style = pinyin.FirstLetter
	iniLetter := 0

	for v := 0; v < 10; v++ {

		var arr [6]int
		for i1 := 0; i1 < 6; i1++ {
			x := rand.Intn(10)
			//排除0和1
			//			if x == 0 || x == 1 {
			//				for {
			//					x = rand.Intn(10)
			//					if x != 0 && x != 1 {
			//						break
			//					}
			//				}
			//			}
			arr[i1] = x
		}
		fmt.Println(arr)
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
								fmt.Println(letters)

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
										str := ""

										for k := 0; k < len(strArr); k++ {
											str := strArr[k]
											fmt.Println(str)
											fmt.Println(pinyin.Pinyin(str, arg))
											pyArr := pinyin.Pinyin(str, arg)
											pyStr := ""
											for l := 0; l < len(pyArr); l++ {
												pyStr += pyArr[l][0]
											}
											fmt.Println(pyStr)
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
}
