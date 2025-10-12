package main

import "fmt"

//slice本身是一个包含指针，长度，容量的结构体
//切片(slice)是对数组一个连续片段的引用，提供了一个相关数组的动态窗口，所以切片是一个引用类型
//这个切片可以是整个数组，或者是由起始和终止索引标识的一些项的子集
//需要注意的是，终止索引标识不包括在切片内
func main() {
	var intarr [6]int = [6]int{3, 6, 9, 1, 4, 7}
	//切片构建在数组之上：
	//[]动态变化的数组长度不写，从1开始切到3结束(不包括3)
	var slice []int = intarr[1:3]
	//or slice:=intarr[1:3]
	fmt.Println(intarr)
	fmt.Println(slice)
	fmt.Println(len(slice))
	//获取切片的容量，容量可以动态变化
	fmt.Println(cap(slice))

	fmt.Printf("下标为一的数组的地址:%p\n", &intarr[1])
	//&slice[0]为指针字段指向底层数组的第一个元素的地址
	fmt.Printf("下标为零的切片的地址:%p\n", &slice[0])
	//&slice为slice本身结构体的内存地址
	fmt.Printf("slice的地址为:%p\n", &slice)
	slice[1] = 16
	fmt.Println(intarr)
	fmt.Println(slice)
}

// 注意：
// 如果主函数定义了一个数组，函数传参的时候可以用切片接收吗？为什么？
// 答案：
// 不可以直接用切片参数接收数组参数。
// 但你可以把数组转为切片后传递。
// ❌ 错误示例（直接传）
// func f(s []int) {      // 想用切片接收
// 	fmt.Println(s)
// }

// func main() {
// 	arr := [5]int{1, 2, 3, 4, 5}
// 	f(arr) // ❌ 编译错误：cannot use arr (type [5]int) as type []int
// }
// 📘 解释：
// [5]int 和 []int 是完全不同的类型；
// Go 的类型系统非常严格，不会自动转换；
// 因为数组包含长度信息 [5] 是类型的一部分。
