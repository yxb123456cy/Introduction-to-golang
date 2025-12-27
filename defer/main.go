package main

import "fmt"

func main() {
	fmt.Println("========== Golang defer 函数规则与示例 ==========")

	// 1. 基本用法：LIFO (后进先出)
	demoLIFO()

	// 2. 参数预计算 (重要规则)
	demoParamEvaluation()

	// 3. 在 Panic 时依然执行 (资源释放保障)
	// demoPanicRecover() // 取消注释可测试 panic 恢复
}

// demoLIFO 演示 defer 的执行顺序：栈结构，后进先出
func demoLIFO() {
	fmt.Println("\n>>> 1. Defer 执行顺序 (LIFO)")

	fmt.Println("Start")

	defer fmt.Println("Defer 1 (最后执行)")
	defer fmt.Println("Defer 2 (倒数第二执行)")
	defer fmt.Println("Defer 3 (最先执行的 defer)")

	fmt.Println("End")
}

// demoParamEvaluation 演示 defer 参数的预计算特性
func demoParamEvaluation() {
	fmt.Println("\n>>> 2. Defer 参数预计算")

	i := 0
	// 规则：defer 后面的函数参数，在 defer 语句声明时就已经确定(计算)好了
	// 这里 i 当前是 0，所以打印的是 "Defer: i = 0"，而不是函数结束时的 10
	defer fmt.Println("Defer: i =", i)

	i++
	fmt.Println("Function: i 增加到了", i)
	i = 10
	fmt.Println("Function: i 最终变成了", i)
}
