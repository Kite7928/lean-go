package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var a bool = true
	var b bool = false
	if a && b {
		fmt.Printf("第一行 - 条件为 true\n")
	}
	if a || b {
		fmt.Printf("第二行 - 条件为 true\n")
	}
	/* 修改 a 和 b 的值 */
	a = false
	b = true
	if a && b {
		fmt.Printf("第三行 - 条件为 true\n")
	} else {
		fmt.Printf("第三行 - 条件为 false\n")
	}
	if !(a && b) {
		fmt.Printf("第四行 - 条件为 true\n")
	}
	arithmetic()
	BinaryFormat(100)

	fmt.Println(BinaryFormat(0))
	fmt.Println(BinaryFormat(-1))
	fmt.Println(BinaryFormat(1))
	fmt.Println(BinaryFormat(260))
	fmt.Println(BinaryFormat(-260))
}

func arithmetic() {
	var a float32 = 8
	var b float32 = 3
	var c float32 = a + b
	var d float32 = a - b
	var e float32 = a * b
	var f float32 = a / b
	fmt.Println("a=%.3f,b=%.3f,c=%.3f,d=%.3f,e=%.3f,f=%.3f \n", a, b, c, d, e, f)
}

func BinaryFormat(n int32) string {
	a := uint32(n)
	sb := strings.Builder{}
	c := uint32(math.Pow(2, 31))
	for i := 0; i < 32; i++ {
		if a&c > 0 {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
		c >>= 1
	}
	return sb.String()
}
