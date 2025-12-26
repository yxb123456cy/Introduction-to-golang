package main

import (
	"fmt"
)

func main() {
	fmt.Println("========== Golang 选择结构 (if/switch/select) 示例 ==========")

	// 1. If-Else 基础与进阶
	demoIfElse()

	// 2. Switch 多种用法
	demoSwitch()

	fmt.Println("\n========== 示例结束 ==========")
}

// demoIfElse 演示 if-else 结构
func demoIfElse() {
	fmt.Println("\n>>> 1. If-Else 结构")

	// --- 基础用法 ---
	age := 18
	if age >= 18 {
		fmt.Println("基础用法: 您已成年")
	} else {
		fmt.Println("基础用法: 您未成年")
	}

	// --- 多重判断 (else if) ---
	score := 85
	fmt.Print("多重判断: ")
	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}

	// --- 特殊用法: If 包含初始化语句 ---
	// 变量 x 只在 if 代码块内有效
	if x := 10; x > 5 {
		fmt.Printf("特殊用法: x=%d 大于 5 (x 的作用域仅限于此 if 块)\n", x)
	}
	// fmt.Println(x) // 这里报错: undefined: x
}

// demoSwitch 演示 switch 结构
func demoSwitch() {
	fmt.Println("\n>>> 2. Switch 结构")

	// --- 基础用法 (匹配值) ---
	day := "Monday"
	fmt.Print("基础用法: ")
	switch day {
	case "Monday":
		fmt.Println("今天是周一")
	case "Tuesday":
		fmt.Println("今天是周二")
	default:
		fmt.Println("其他日子")
	}

	// --- 多条件匹配 ---
	char := 'e'
	fmt.Print("多条件匹配: ")
	switch char {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Printf("'%c' 是元音字母\n", char)
	default:
		fmt.Printf("'%c' 是辅音字母\n", char)
	}

	// --- 无表达式 switch (替代复杂的 if-else) ---
	// 类似 if-else if-else 链，更加清晰
	score := 75
	fmt.Print("无表达式: ")
	switch {
	case score >= 90:
		fmt.Println("A等级")
	case score >= 80:
		fmt.Println("B等级")
	default:
		fmt.Println("C等级或以下")
	}

	// --- Fallthrough (穿透) ---
	// 默认情况下 switch 匹配成功后会自动 break 所以开发者无需自行手动添加break，使用 fallthrough 强制执行下一个 case
	fmt.Print("Fallthrough: ")
	num := 10
	switch num {
	case 10:
		fmt.Print("匹配到10 -> ")
		fallthrough // 强制执行下一个 case，不判断条件
	case 20:
		fmt.Println("穿透执行到了这里")
	case 30:
		fmt.Println("不会执行到这里")
	}
}
