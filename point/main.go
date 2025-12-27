package main

import "fmt"

// main 函数：Golang 指针 (Pointer) 基础与进阶
// 指针存储的是变量的内存地址，使用 & 取地址，使用 * 解引用
func main() {
	fmt.Println("========== Golang 指针 (Pointer) 运用示例 ==========")

	// 1. 指针基础：声明、取地址、解引用
	demoPointerBasic()

	// 2. 重点：值传递 vs 指针传递 (引用传递)
	demoValueVsPointer()

	// 3. 数组指针 vs 指针数组
	demoArrayPointer()

	fmt.Println("\n========== 示例结束 ==========")
}

// demoPointerBasic 演示指针基础
func demoPointerBasic() {
	fmt.Println("\n>>> 1. 指针基础")

	// 定义一个普通变量
	var num int = 10
	fmt.Printf("变量 num 的值: %d\n", num)
	fmt.Printf("变量 num 的内存地址: %p\n", &num)

	// 定义一个指针变量，指向 num
	// &num 表示取出 num 的地址
	var ptr *int = &num
	fmt.Printf("指针 ptr 存储的值 (即 num 的地址): %p\n", ptr)
	fmt.Printf("指针 ptr 自己的地址: %p\n", &ptr)

	// *ptr 表示解引用：取出指针指向地址中存储的值
	fmt.Printf("通过指针 *ptr 获取的值: %d\n", *ptr)

	// 通过指针修改值
	*ptr = 20
	fmt.Printf("修改 *ptr 后，num 的值变为: %d\n", num)
}

// demoValueVsPointer 演示值类型与引用类型(指针传递)的区别
func demoValueVsPointer() {
	fmt.Println("\n>>> 2. 值传递 vs 指针传递")

	x := 100
	y := 100

	fmt.Printf("初始状态: x=%d, y=%d\n", x, y)

	// 调用值传递函数
	modifyByValue(x)
	fmt.Printf("值传递调用后 (x): %d (未改变)\n", x)

	// 调用指针传递函数
	modifyByPointer(&y)
	fmt.Printf("指针传递调用后 (y): %d (已改变)\n", y)
}

// modifyByValue 值传递：只拷贝副本，不影响原变量
func modifyByValue(num int) {
	num = 999
	fmt.Printf("  -> [函数内] 修改副本 num = %d\n", num)
}

// modifyByPointer 指针传递：传递地址，修改该地址的内容会影响原变量
func modifyByPointer(ptr *int) {
	*ptr = 999
	fmt.Printf("  -> [函数内] 修改 *ptr = %d\n", *ptr)
}

// demoArrayPointer 演示数组指针与指针数组
func demoArrayPointer() {
	fmt.Println("\n>>> 3. 数组指针 vs 指针数组")

	// --- 数组指针 (Pointer to Array) ---
	// 它是一个指针，指向一个数组
	arr := [3]int{1, 2, 3}
	var ptrToArr *[3]int = &arr

	fmt.Printf("数组指针类型: %T\n", ptrToArr)
	fmt.Printf("通过数组指针访问元素: %d\n", ptrToArr[0]) // Go 语言允许直接通过指针访问下标，等同于 (*ptrToArr)[0]

	// --- 指针数组 (Array of Pointers) ---
	// 它是一个数组，里面存的都是指针
	a, b, c := 10, 20, 30
	var arrOfPtr [3]*int = [3]*int{&a, &b, &c}

	fmt.Printf("指针数组类型: %T\n", arrOfPtr)
	fmt.Printf("指针数组第0个元素存的地址: %p\n", arrOfPtr[0])
	fmt.Printf("指针数组第0个元素解引用的值: %d\n", *arrOfPtr[0])
}
