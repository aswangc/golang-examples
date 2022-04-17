package main

import (
	"fmt"
	"unsafe"
)

/*
	切片（slice）是对数组的一个连续片段的引用，这个片段可以是整个数组，也可以是由起始和终止索引标识的一些项的子集，需要注意的是，终止索引标识的项不包括在切片内。
*/

func main() {
	//Examples1()
	//Examples2()
	//Examples3()
	//Examples4()
	Examples5()

}

// Examples1 基于数组创建一个切片
func Examples1() {
	// 创建一个数组
	arr1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 基于arr1创建一个切片(切片属于半开形式：不包括下标为5的元素)
	slice1 := arr1[1:5]

	// 打印切片
	fmt.Printf("print slice:%+v\n", slice1)
}

// Examples2 基于数组创建一个切片，然后将该切片扩容
func Examples2() {
	// 创建一个数组
	arr1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 基于arr1创建一个切片(切片属于半开形式：不包括下标为5的元素)
	slice1 := arr1[1:5]

	// 打印切片 result: print slice:[2 3 4 5]
	fmt.Printf("print slice:%v\n", slice1)

	// 打印切片的容量  result: print slice cap:9
	fmt.Printf("print slice cap:%d\n", cap(slice1))

	// 打印切片的长度 result： print slice len:4
	fmt.Printf("print slice len:%d\n", len(slice1))

	// 对切片进行扩容测试
	slice1 = append(slice1, 11, 12, 13)

	// 打印扩容后的切片的：长度、容量和值 result: print slice len:7, cap:9, value:[2 3 4 5 11 12 13]
	// 从结果来看，基于数组的切片支持扩容
	fmt.Printf("print slice len:%d, cap:%d, value:%v\n", len(slice1), cap(slice1), slice1)
}

// Examples3 基于数组创建一个切片，然后修改原始数组，看是否会影响到切片
func Examples3() {
	// 创建一个数组
	arr1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 基于arr1创建一个切片(切片属于半开形式：不包括下标为5的元素)
	slice1 := arr1[1:]

	fmt.Println("-----原始数组修改前-----")

	// 打印原始数组
	fmt.Printf("print arr:%v\n", arr1)

	// 打印切片 result: print slice:[2 3 4 5]
	fmt.Printf("print slice:%v\n", slice1)

	// 打印切片的容量  result: print slice cap:9
	fmt.Printf("print slice cap:%d\n", cap(slice1))

	// 打印切片的长度 result： print slice len:4
	fmt.Printf("print slice len:%d\n", len(slice1))

	// 修改原始的数组
	arr1[0] = 101
	arr1[1] = 102
	arr1[2] = 103
	arr1[3] = 104
	arr1[4] = 105
	arr1[5] = 106
	arr1[6] = 107
	arr1[7] = 108
	arr1[8] = 109
	arr1[9] = 110

	fmt.Println("-----原始数组修改后-----")

	// 打印原始数组
	fmt.Printf("print arr:%v\n", arr1)

	// 打印切片 result: print slice:[2 3 4 5]
	fmt.Printf("print slice:%v\n", slice1)

	// 打印切片的容量  result: print slice cap:9
	fmt.Printf("print slice cap:%d\n", cap(slice1))

	// 打印切片的长度 result： print slice len:4
	fmt.Printf("print slice len:%d\n", len(slice1))

	// 对切片进行扩容测试
	slice1 = append(slice1, 11, 12, 13)

	fmt.Println("-----对切片进行扩容后-----")

	// 打印原始数组
	fmt.Printf("print arr:%v\n", arr1)

	// 打印切片 result: print slice:[2 3 4 5]
	fmt.Printf("print slice:%v\n", slice1)

	// 打印切片的容量  result: print slice cap:9
	fmt.Printf("print slice cap:%d\n", cap(slice1))

	// 打印切片的长度 result： print slice len:4
	fmt.Printf("print slice len:%d\n", len(slice1))

	// 打印切片的内存地址
	fmt.Println("打印切片的内存地址")
	for i := range slice1 {
		fmt.Printf("value:%d, point:%p\n", slice1[i], &slice1[i])
	}

	// 打印数组内存地址
	fmt.Println("打印数组内存地址")
	for i := range arr1 {
		fmt.Printf("value:%d, point:%p\n", arr1[i], &arr1[i])
	}

	/*
		这个结果是对原始数组的完全扩容后，对切片进行扩容的结果
		可以看出：切片就是对原油数组的引用，至于扩容，就是在原始的数组上面进行扩张，并且内存地址还是连续的
			打印切片的内存地址
				value:102, point:0x140000b8000
				value:103, point:0x140000b8008
				value:104, point:0x140000b8010
				value:105, point:0x140000b8018
				value:106, point:0x140000b8020
				value:107, point:0x140000b8028
				value:108, point:0x140000b8030
				value:109, point:0x140000b8038
				value:110, point:0x140000b8040
				value:11, point:0x140000b8048
				value:12, point:0x140000b8050
				value:13, point:0x140000b8058

			打印数组内存地址
				value:101, point:0x140000ae000
				value:102, point:0x140000ae008
				value:103, point:0x140000ae010
				value:104, point:0x140000ae018
				value:105, point:0x140000ae020
				value:106, point:0x140000ae028
				value:107, point:0x140000ae030
				value:108, point:0x140000ae038
				value:109, point:0x140000ae040
				value:110, point:0x140000ae048

		这个结果是对原属数组不完全切片后，对切片进行扩容的结果
		可以看出：切片是只是对原始数组的引用，然后切片的扩容还会把原始数组的值冲掉。
			打印切片的内存地址
				value:102, point:0x140000220a8
				value:103, point:0x140000220b0
				value:104, point:0x140000220b8
				value:105, point:0x140000220c0
				value:11, point:0x140000220c8
				value:12, point:0x140000220d0
				value:13, point:0x140000220d8
			打印数组内存地址
				value:101, point:0x140000220a0
				value:102, point:0x140000220a8
				value:103, point:0x140000220b0
				value:104, point:0x140000220b8
				value:105, point:0x140000220c0
				value:11, point:0x140000220c8
				value:12, point:0x140000220d0
				value:13, point:0x140000220d8
				value:109, point:0x140000220e0
				value:110, point:0x140000220e8
	*/
}

// Examples4 切片的生长（copy and append 函数）
func Examples4() {
	// 要增加切片的容量必须创建一个新的、更大容量的切片，然后将原有切片的内容复制到新的切片。 整个技术是一些支持动态数组语言的常见实现。

	// 创建一个切片
	slice1 := []int{1, 2, 3}

	// 打印slice1的容量
	fmt.Printf("slice1的容量:%d\n", cap(slice1))

	// 打印slice1的长度
	fmt.Printf("slice1的长度:%d\n", len(slice1))

	// 将切片容量翻倍
	// 步骤一： 重新创建一个容量是slice1一倍的切片slice2
	slice2 := make([]int, len(slice1), cap(slice1)*2)

	// 步骤二： 将slice1的值拷贝到slice2
	for i := range slice1 {
		slice2[i] = slice1[i]
	}

	// 将slice2重新赋值给slice1
	slice1 = slice2

	// 打印slice1的容量
	fmt.Printf("slice1的容量:%d\n", cap(slice1))

	// 打印slice1的长度
	fmt.Printf("slice1的长度:%d\n", len(slice1))
}

// Examples5 测试数组什么时候会被GC回收掉
func Examples5() {

	test1 := func() []int {
		// 创建一个数组
		arr := [...]int{1, 2, 3, 4, 5}

		// 打印源数组的地址
		fmt.Println("打印源数组的地址和值")
		for i := range arr {
			fmt.Printf("value：%v，pointer：%p\n", arr[i], &arr[i])
		}

		// 基于arr创建一个切片
		slice1 := arr[:4]

		// 创建一个新切片
		slice2 := make([]int, len(slice1))

		// 将slice1拷贝给slice2
		for i := range slice1 {
			slice2[i] = slice1[i] // 这里发生的是值拷贝
		}

		return slice2
	}

	result := test1()

	// 打印切片的地址
	fmt.Println("打印切片的地址")
	for i := range result {
		fmt.Printf("value：%v，pointer：%p\n", result[i], &result[i])
	}

	resultPointer := unsafe.Pointer(&result[3])
	int8Point := (*int)(unsafe.Pointer(uintptr(resultPointer) + 8))
	fmt.Println(*int8Point)

	/*
		如果函数返回的是数组切片拷贝，那么当函数执行结束后，GC就会回收源数组，所以通过指针偏移方法是无法拿到切片之外源数组的值 -> 0
			➜  slice go run main.go
			打印源数组的地址和值
			value：1，pointer：0x1400001a1e0
			value：2，pointer：0x1400001a1e8
			value：3，pointer：0x1400001a1f0
			value：4，pointer：0x1400001a1f8
			value：5，pointer：0x1400001a200
			打印切片的地址
			value：1，pointer：0x1400001e0a0
			value：2，pointer：0x1400001e0a8
			value：3，pointer：0x1400001e0b0
			value：4，pointer：0x1400001e0b8
			0

		如果函数返回的是数组的切片，那么当函数执行结束后，GC就不会回收源数组，所以通过指针偏移方法是可以拿到切片之外源数组的值 -> 5
			➜  slice go run main.go
			打印源数组的地址和值
			value：1，pointer：0x14000120060
			value：2，pointer：0x14000120068
			value：3，pointer：0x14000120070
			value：4，pointer：0x14000120078
			value：5，pointer：0x14000120080
			打印切片的地址
			value：1，pointer：0x14000120060
			value：2，pointer：0x14000120068
			value：3，pointer：0x14000120070
			value：4，pointer：0x14000120078
			5
	*/
}
