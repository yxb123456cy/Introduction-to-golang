package main

import (
	"fmt"
)

// 全局变量初始化 (最先执行)
var globalVar = initGlobal()

func initGlobal() string {
	fmt.Println("1. 全局变量初始化完成")
	return "Global Variable"
}

// init 函数 (在全局变量初始化后，main 函数执行前自动执行)
// 用途：常用于初始化配置、数据库连接、注册组件等
func init() {
	fmt.Println("2. init 函数执行 (init 1)")
}

// 一个包中可以包含多个 init 函数，按出现顺序执行
func init() {
	fmt.Println("3. init 函数执行 (init 2)")
}

func main() {
	fmt.Println("4. main 函数开始执行")
	fmt.Println("========== Golang init 函数规则 ==========")
	fmt.Println("执行顺序: 全局变量定义 -> init() 函数 -> main() 函数")
	fmt.Println("特点:")
	fmt.Println("  1. init 函数没有参数也没有返回值")
	fmt.Println("  2. init 函数由 Go 运行时自动调用，不能显式调用")
	fmt.Println("  3. 一个文件/包可以有多个 init 函数")
	fmt.Println("  4. 如果导入了其他包，会先执行被导入包的 init 函数")
}
