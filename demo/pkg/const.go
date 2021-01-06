package pkg

import (
	"fmt"
	"math"
)

var pak_name = "常量const"

const PI = 3.14
const age = 30 / 2
const sex = 139479388230.249887449801287324128782342348732783248 //用反斜杠连接多行表达式
const log2E = 1 / math.Ln2 //表达式
const hardEight = (1 << 100) >> 97

const Monday,Tuesday = 1,2
const (
	Monday1, Tuesday1 = 3, 4
)

const(
	iota_a = iota
	iota_b
	iota_c
	iota_d = 5
	iota_e //没赋值的常量等于上一个的值
)

func Doit_const() {
	fmt.Println(pak_name,"------------------------------------------------------------------------")

	fmt.Println("PI = ", PI)
	fmt.Println("age = ", age)
	fmt.Println("sex = ", sex)
	fmt.Println("log2E = ", log2E)
	fmt.Println("hardEight = ", hardEight)
	fmt.Println("Monday = ", Monday)
	fmt.Println("Tuesday = ", Tuesday)
	fmt.Println("Monday1 = ", Monday1)
	fmt.Println("Tuesday1 = ", Tuesday1)

	fmt.Println("iota作为枚举值，从0开始，每引用一次自增加一")
	fmt.Println("iota_a = ", iota_a)
	fmt.Println("iota_b = ", iota_b)
	fmt.Println("iota_c = ", iota_c)
	fmt.Println("iota_d = ", iota_d)
	fmt.Println("iota_e = ", iota_e)

	fmt.Println("------------------------------------------------------------------------")

}
