package pkg

import (
	"fmt"
	"math"
)

const PI = 3.14
const age = 30 / 2
const sex = 139479388230.249887449801287324128782342348732783248
const log2E = 1 / math.Ln2
const hardEight = (1 << 100) >> 97

func Init_const() {
	fmt.Println("PI = ", PI)
	fmt.Println("age = ", age)
	fmt.Println("sex = ", sex)
	fmt.Println("log2E = ", log2E)
	fmt.Println("hardEight = ", hardEight)
}
