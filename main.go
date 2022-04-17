package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//定义一个长度为3的int8类型数组
	a := [3]int8{6, 8, 9}

	//取出数组第一个位置的地址
	a_first_point := &a[0]
	a_first_unsafe_point := unsafe.Pointer(a_first_point)
	fmt.Println("a[0]的地址为：", a_first_unsafe_point)

	//指针只能一个字节字节取，int8占一个字节，所以看到值只加了1
	fmt.Println("a[1]的地址为：", unsafe.Pointer(&a[1]))

	//把a_first_unsafe_point转成uintptr类型，就可以指针运算了
	a_uintptr_first_unsafe_point := uintptr(a_first_unsafe_point)

	//指针+1 表示到了数组的第二个位置
	a_uintptr_first_unsafe_point++
	fmt.Println("a[0]位置指针自增1后，的指针位置：", a_uintptr_first_unsafe_point)

	//打印出来可以看到跟&a[1]的地址是一样的
	a_uintptr_second_unsafe_point := unsafe.Pointer(a_uintptr_first_unsafe_point)
	fmt.Println("a[0]位置指针自增1后，的指针位置，转成unsafe_Pointer类型：", a_uintptr_second_unsafe_point)

	//将该指针转换成 *int8类型（因为它本身就是*int8类型）
	int8_point := (*int8)(a_uintptr_second_unsafe_point)

	//解引用，得到指针对应的结果，就是数组的第二个值，8
	fmt.Println(*int8_point)
}
