package demo

import "fmt"

func Doit_variable() {
	var pak_name = "常量variable"
	fmt.Println(pak_name + "------------------------------------------------------------------------")
	var inta int
	var floata float32
	var stringa string
	var boola bool
	var pointa *int
	fmt.Println("变量默认值 int:", inta)
	fmt.Println("变量默认值 float:", floata)
	fmt.Println("变量默认值 string:", stringa)
	fmt.Println("变量默认值 bool:", boola)
	fmt.Println("变量默认值 point:", pointa)

	fmt.Println("------------------------------------------------------------------------")

}
