package main

import (
	"fmt"
	"reflect"
)

/*
	推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
*/

func main() {
	//examples1()
	//examples2()
	//examples3()
	//examples4()
	fmt.Println(examples5())
}

// examples1 go defer 执行顺序：相当于压栈，先进后出，所以先打印 "tow"， 再打印 "one"
func examples1() {
	defer println("one")
	defer println("tow")
}

// examples1 遵从先进后出原则，输出为：9，8，7，6，5，4，3，2，1
func examples2() {
	for i := 1; i < 10; i++ {
		defer println(i)
	}
}

// examples3 在defer被压入栈中时，变量i的值已经被确定（i是值传递），所以后面的i++并不会影响defer的打印
func examples3() {
	i := 0
	defer println(i)
	i = 100
}

// examples4 arr 是引用传递，所以defer后面的函数拿到的是指针，所以最终输出会受到defer后面函数的影响
func examples4() {
	var arr = new([]int)
	fmt.Printf("%s\n", reflect.TypeOf(arr))
	defer fmt.Println(arr)
	*arr = append(*arr, 1, 2, 3)
}

// examples5 需要注意一个原则：defer一定是在return后面执行，这样就不难理解，i的默认值为0，return将i赋值为1，最后执行defer将1++，座钟返回的值就为2
func examples5() (i int) {
	defer func() { i++ }()
	return 1
}
