package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("========== Golang 基本数据类型示例 ==========")

	// 1. 整数类型 (Integer)
	demoInteger()

	// 2. 浮点类型 (Float)
	demoFloat()

	// 3. 布尔类型 (Boolean)
	demoBool()

	// 4. 字符串类型 (String)
	demoString()

	// 5. 字符类型 (Char - Byte & Rune)
	demoChar()

	// 6. 零值 (Default Zero Values)
	demoZeroValues()

	fmt.Println("\n========== 示例结束 ==========")
}

// demoInteger 演示整数类型
func demoInteger() {
	fmt.Println("\n>>> 1. 整数类型 (Integer)")

	// 有符号整数 (Signed)
	var i1 int = 10   // int 的大小取决于系统 (32位或64位)
	var i2 int8 = 127 // -128 到 127
	var i3 int64 = 9223372036854775807

	// 无符号整数 (Unsigned)
	var u1 uint = 20   // uint 的大小也取决于系统
	var u2 uint8 = 255 // 0 到 255 (常用于 byte)

	fmt.Printf("int: %d, size: %d bytes\n", i1, unsafe.Sizeof(i1))
	fmt.Printf("int8: %d\n", i2)
	fmt.Printf("int64: %d\n", i3)
	fmt.Printf("uint: %d\n", u1)
	fmt.Printf("uint8: %d\n", u2)
}

// demoFloat 演示浮点类型
func demoFloat() {
	fmt.Println("\n>>> 2. 浮点类型 (Float)")

	// float32: 单精度，精度约 7 位小数
	var f1 float32 = 3.1415926
	// float64: 双精度，精度约 15 位小数 (Golang 默认浮点类型)
	var f2 float64 = 3.141592653589793

	fmt.Printf("float32: %v (精度丢失风险)\n", f1)
	fmt.Printf("float64: %v (推荐使用)\n", f2)
}

// demoBool 演示布尔类型
func demoBool() {
	fmt.Println("\n>>> 3. 布尔类型 (Boolean)")

	var isOpen bool = true
	var isClosed bool = false

	fmt.Println("isOpen:", isOpen)
	fmt.Println("isClosed:", isClosed)

	// 逻辑运算
	fmt.Println("AND (true && false):", isOpen && isClosed)
	fmt.Println("OR  (true || false):", isOpen || isClosed)
	fmt.Println("NOT (!true):", !isOpen)
}

// demoString 演示字符串类型
func demoString() {
	fmt.Println("\n>>> 4. 字符串类型 (String)")

	// 普通字符串 (双引号) - 支持转义字符
	var str1 string = "Hello, Golang!\n这是新的一行"
	fmt.Println("双引号字符串:", str1)

	// 原生字符串 (反引号) - 原样输出，支持多行
	var str2 string = `
	这是一个原生字符串
	它可以包含换行
	无需转义 \n \t
	`
	var str3 string = `
	gogogo
	出发喽`
	fmt.Println(str3)
	fmt.Println("反引号字符串:", str2)

	// 字符串拼接
	fmt.Println("拼接:", "Go"+" "+"Language")

	// 长度
	fmt.Println("字符串长度 (len):", len("Hello"))
}

// demoChar 演示字符类型
func demoChar() {
	fmt.Println("\n>>> 5. 字符类型 (Char - Byte & Rune)")

	// 在 Go 中没有 char 类型
	// byte 是 uint8 的别名，用于表示 ASCII 字符
	var c1 byte = 'a'
	fmt.Printf("byte (ASCII 'a'): 值=%d, 字符=%c, 类型=%T\n", c1, c1, c1)

	// rune 是 int32 的别名，用于表示 Unicode 字符 (如中文)
	var c2 rune = '中'
	fmt.Printf("rune (Unicode '中'): 值=%d, 字符=%c, 类型=%T\n", c2, c2, c2)
}

// demoZeroValues 演示零值机制
func demoZeroValues() {
	fmt.Println("\n>>> 6. 零值 (Zero Values)")
	fmt.Println("当声明变量未赋值时，Go 会自动赋予默认值：")

	var i int
	var f float64
	var b bool
	var s string
	var ptr *int

	fmt.Printf("int 默认值: %d\n", i)
	fmt.Printf("float64 默认值: %f\n", f)
	fmt.Printf("bool 默认值: %t\n", b)
	fmt.Printf("string 默认值: %q (空字符串)\n", s)
	fmt.Printf("pointer(指针类型) 默认值: %v (nil)\n", ptr)
}
