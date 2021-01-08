package demo

import (
	"fmt"
	"strconv"
)

func Doit_strconv() {
	a := 5.12
	b := int(a)
	fmt.Println("简单转换：float -> int: ", a, b)
	fmt.Println("strconv主要是简单类型和字符串之间的转换")

	fmt.Println("string to int : a->", strconv.Itoa(32))
	re, _ := strconv.Atoi("32")
	fmt.Println("int to string : ", re)

	fmt.Println("Parse类函数转换指定类型为字符串类型")
	c1, _ := strconv.ParseBool("true")
	fmt.Println("ParseBool:", c1)
	c2, _ := strconv.ParseFloat("3.1415926", 64)
	fmt.Println("ParseFloat:", c2)
	c3, _ := strconv.ParseInt("101011110", 2, 0)
	fmt.Println("ParseInt 二进制 ", c3)
	c4, _ := strconv.ParseInt("101011110", 5, 0)
	fmt.Println("ParseInt 五进制 ", c4)

	fmt.Println("Format类函数转换字符串为指定类型")
	d1 := strconv.FormatBool(true)
	fmt.Println("FormatBool ", d1)
	d2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	fmt.Println("FormatFloat ", d2)
	d3 := strconv.FormatInt(-42, 16)
	fmt.Println("FormatInt 16进制", d3)
	d4 := strconv.FormatUint(42, 16)
	fmt.Println("FormatInt 16进制 ", d4)

	fmt.Println("Append类函数跟Format类似，只是将转换后的结果追加到一个slice")
	e1 := []byte("int (base 10):")
	e1 = strconv.AppendBool(e1, true)
	fmt.Println("AppendBool ", string(e1))
	e1 = strconv.AppendFloat(e1, 3.1415, 'E', -1, 64)
	fmt.Println("AppendFloat ", string(e1))
	e1 = strconv.AppendInt(e1, 42, 16)
	fmt.Println("AppendInt ", string(e1))
}
