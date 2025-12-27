package main

import (
	"LuFei_go_study/func/utils"
	"fmt"
)

// main 函数：Golang 函数 (Function) 基础与进阶
func main() {
	fmt.Println("========== Golang 函数运用示例 ==========")
	// 调用外部包的函数 (假设 utils 包已存在)
	fmt.Printf("今天是: %s\n", utils.GetToday())

	// 1. 基础函数调用
	res := add(10, 20)
	fmt.Printf("1. 基础函数 (10+20): %d\n", res)

	// 2. 多返回值
	q, r := div(10, 3)
	fmt.Printf("2. 多返回值 (10/3): 商=%d, 余数=%d\n", q, r)

	// 3. 可变参数
	sum := sumAll(1, 2, 3, 4, 5)
	fmt.Printf("3. 可变参数 (1+2+3+4+5): %d\n", sum)

	// 4. 匿名函数与闭包
	demoClosure()

	// 5. 函数作为参数 (高阶函数)
	demoHigherOrder()

	// 6. Defer 延迟执行
	demoDefer()

	fmt.Println("\n========== 示例结束 ==========")
}

// 1. 基础函数：参数类型简写
// 如果参数类型相同，可以只写最后一个
func add(a, b int) int {
	return a + b
}

// 2. 多返回值：Golang 特色
// 支持命名返回值 (q, r)，return 时可省略变量名
func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return // 等同于 return q, r
}

// 3. 可变参数：参数会被包装成切片
func sumAll(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// 4. 匿名函数与闭包
func demoClosure() {
	fmt.Println("\n>>> 4. 匿名函数与闭包")

	// 定义并立即调用匿名函数
	func(name string) {
		fmt.Printf("立即调用的匿名函数: Hello, %s\n", name)
	}("Golang")

	// 闭包：函数 + 其引用的外部变量
	counter := getCounter()
	fmt.Println("闭包计数器:", counter()) // 1
	fmt.Println("闭包计数器:", counter()) // 2
	fmt.Println("闭包计数器:", counter()) // 3
}

// getCounter 返回一个闭包函数
func getCounter() func() int {
	count := 0 // 这个变量会被闭包捕获并保留状态
	return func() int {
		count++
		return count
	}
}

// 5. 高阶函数：函数作为参数或返回值
func demoHigherOrder() {
	fmt.Println("\n>>> 5. 高阶函数 (回调)")

	// 定义一个处理函数
	process := func(x int) int {
		return x * x
	}

	// 将函数作为参数传递
	result := apply(5, process)
	fmt.Printf("apply(5, square) = %d\n", result)
}

// apply 接收一个整数和一个函数作为参数
func apply(val int, op func(int) int) int {
	return op(val)
}

// 6. Defer 延迟执行
func demoDefer() {
	fmt.Println("\n>>> 6. Defer 延迟执行")

	defer fmt.Println("Defer 1: 最后执行 (栈结构，后进先出)")
	defer fmt.Println("Defer 2: 倒数第二执行")

	fmt.Println("Normal: 正常逻辑执行中...")
	// 常用于资源释放，如 file.Close(), mutex.Unlock()
}
