package main

import "fmt"

func test01(arr [5]int) {
	arr[0] = 7
}

func test02(arr *[5]int) {
	(*arr)[0] = 7
}

func main() {
	//定义数组并且初始化
	// var scores [5]int = [5]int{3, 6, 9, 11, 13}
	scores := [5]int{3, 6, 9, 11, 13}
	// var scores = [...]int{3, 6, 9, 11, 13}
	// var scores=[...]int{2:9,0:3,1:6,3:11,4:13}

	//数组存值
	scores[0] = 1
	scores[1] = 2
	scores[2] = 3
	scores[3] = 4
	scores[4] = 5
	//遍历数组
	//普通for循环
	for i := 0; i < len(scores); i++ {
		fmt.Printf("请录入第%d个成绩", scores[i])
		fmt.Scanln(&scores[i])
	}
	fmt.Println("-------------------------")
	//键值循环 for-range
	for key, val := range scores {
		fmt.Println("索引为：", key+1, "索引上的值为：", val)
	}

	fmt.Println(scores)
	fmt.Printf("%T\n", scores) //长度属于类型的一部分,[3]int和[4]int类型不一样
	fmt.Printf("%T\n", &scores)
	fmt.Printf("%p\n", &scores)
	fmt.Printf("%p\n", &scores[1])

	test01(scores)
	fmt.Println(scores[0]) //与c++不同
	test02(&scores)
	fmt.Println(scores[0])
}

//注意：
// 	🧠 在 C++ 里：
// int n = 5;
// int* arr = new int[n];
// 这会在堆上分配一块 连续的、长度为 n 的 int 数组内存。
// arr 是一个指针，指向这块内存的首地址。
// 这是 纯内存分配行为（没有类型信息参与）。
// C++ 的类型系统允许 “运行时动态长度的数组”。
// ⚙️ 底层机制：
// 编译器在堆上调用 operator new(n * sizeof(int))
// 不关心数组的长度是不是常量，反正是“原始内存块”。
// 🧠 在 Go 里：
// p := new([5]int)   // ✅ OK：固定长度
// p := new([n]int)   // ❌ 编译错误
// 因为 [n]int 是 类型，而在 Go 里：
// 数组的长度是类型的一部分，必须在编译期已知。
// 所以 [5]int、[10]int 是两种完全不同的类型。
// new([n]int) 只有当 n 是编译时常量时才合法。

// 如果用数组指针接收，数组指针可以定义为 *[...]int 吗？
// 答案：
// 不能用 *[...]int，必须指定长度，例如 *[5]int。
// 📘 原因：
// 在 Go 中，[...]int 只能用在 定义数组时 自动推导长度，比如：
// arr := [...]int{1, 2, 3} // 编译器推导长度 3
// 但在类型定义中，长度必须是常量表达式，不能用 ...：
// func f(a *[3]int) {}  // ✅ 正确
// func g(a *[...]int) {} // ❌ 错误：cannot use [...] in type declaration
