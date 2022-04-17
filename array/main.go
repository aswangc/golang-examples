package main

import "fmt"

/*
	数组类型定义了长度和元素类型
	数组是线性表结构，从而内存空间是连续的且有着相同类型的数据。
	正是由于是线性结构的特性才支持随机访问，时间复杂度为o(1)。
	通过公式a[i]_address = base_address + i * data_type_size计算得到数组元素地址。
	但是其插入、删除操作比较低效，因为会涉及到数据的迁移问题。以上就是数组的基本特性。
*/

func main() {
	//Examples1()
	//Examples2()
	Examples3()

}

// Examples1 声明及初始化
func Examples1() {
	// 数组不需要显式的初始化；数组的零值是可以直接使用的，数组元素会自动初始化为其对应类型的零值：
	var arr1 [4]int // 类型 [4]int 对应内存中四个连续的整数

	// 数组的初始化值
	fmt.Println(arr1) // [0 0 0 0]

	//如果数组长度不确定，可以使用 ... 代替数组的长度，编译器会根据元素个数自行推断数组的长度：
	var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance)
}

// Examples2 一个数组变量被赋值或者被传递的时候，实际上会复制整个数组
func Examples2() {
	getArrId := func(arr [5]int) {
		fmt.Printf("传递后的数组内存地址：%p\n", &arr)
	}

	// 申明一个数组
	var arr1 [5]int

	// 查看数组的内存地址
	fmt.Printf("传递前的数组内存地址：%p\n", &arr1)

	// 打印传递后的数组内存地址
	getArrId(arr1) // 通过下面的输出结果不难发现，数组在传递的过程中是值传递，而不是引用传递

	// 运行输出
	// 传递前的数组内存地址：0x14000120060
	// 传递后的数组内存地址：0x140001200c0
}

// Examples3 指针数组(引用)
func Examples3() {
	// 可以声明一个指针类型的数组，这样数组中就可以存放指针。注意，指针的默认初始化值为nil。

	type students struct {
		sex  string
		name string
	}

	// 创建一个指针数组
	arr1 := [5]*students{{sex: "男", name: "01"}}

	getArrId := func(arr [5]*students) [5]*students {
		// 对值传递进来的指针数组的值修改
		arr[0].sex = "女"
		return arr
	}

	fmt.Printf("传递前的数组的值:%+v：%p\n", arr1[0], arr1[0])
	fmt.Printf("传递前的数组内存地址：%p\n", &arr1)

	arr2 := getArrId(arr1)

	fmt.Printf("传递后的指针数组的值:%+v：%p\n", arr2[0], arr2[0])
	fmt.Printf("传递后的指针数组内存地址：%p\n", &arr2)

	// 运行输出
	// 传递前的数组的值:&{sex:男 name:01}：0x1400008e020
	// 传递前的数组内存地址：0x140000a4030
	// 传递后的指针数组的值:&{sex:女 name:01}：0x1400008e020
	// 传递后的指针数组内存地址：0x140000a4060

	// 指针数组在值传递时依旧是值传递，但是数组的值是指针，所以依旧是指针传递

}
