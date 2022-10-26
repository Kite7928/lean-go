package main

import "fmt"

//显示定义,定义变量的话在go中使用必须使用，所以匿名_常量使用广泛
const tb0 string = "abc"
const tb = "abc"

func main() {
	//const LENGTH int = 10
	//const WIDTH int = 5
	//var area int
	//const a, b, c = 1, false, "str" //多重赋值

	//area = LENGTH * WIDTH
	//fmt.Printf("面积为 : %d", area)
	//fmt.Println(a, b, c)

	//------------------------------------
	/*	iota，特殊常量，可以认为是一个可以被编译器修改的常量
		如果中断iota自增，则必须显式恢复。且后续自增值按行序递增
		自增默认是int类型，可以自行进行显示指定类型
		数字常量不会分配存储空间，无须像变量那样通过内存寻址来取值，因此无法获取地址
		使用iota能简化定义，在定义枚举时很有用。
		每次 const 出现时，都会让 iota 初始化为0.*/
	{
		const (
			a = iota          //0
			b                 //1
			c                 //2
			d = "ha"          //独立值，iota += 1
			e                 //"ha"   iota += 1
			f string = "1000" //iota +=1
			g                 //100  iota +=1
			h = iota          //7,恢复计数
			i                 //8
		)
		fmt.Println(a, b, c, d, e, f, g, h, i)
		const (
			kv = iota
			q
			p
			v
		)
		fmt.Println(kv, p, v, q)
	}

	//常量组中如不指定类型和初始化值，则与上一行非空常量右值相同
	/*	常量的注意事项：

		常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型
		不曾使用的常量，在编译的时候，是不会报错的
		显示指定类型的时候，必须确保常量左右值类型一致，需要时可做显示类型转换。这与变量就不一样了，变量是可以是不同的类型值*/
	const (
		x uint16 = 16
		y
		s = "abc"
		z
	)
	fmt.Printf("%T,%v\n", y, y)
	fmt.Printf("%T,%v\n", z, z)

}

//作为枚举，常量组
const (
	Unknown = 0
	Female  = 1
	Male    = 2
)
