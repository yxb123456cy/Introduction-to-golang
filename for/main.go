package main

import "fmt"

// main 函数：For 循环基础与进阶
// Golang 只有 for 关键字，没有 while 或 do-while，所有循环都由 for 实现。
func main() {
	fmt.Println("========== Golang For 循环运用示例 ==========")

	// 1. 标准 For 循环 (类似 C/Java)
	demoStandardFor()

	// 2. While 风格循环 (仅条件)
	demoWhileStyle()

	// 3. 无限循环与控制 (break/continue)
	demoInfiniteLoop()

	// 4. For-Range 循环 (强大且常用)
	demoForRange()

	// 5. 标签与跳转 (Label/Goto) - 了解即可
	demoLabelBreak()

	fmt.Println("\n========== 示例结束 ==========")
}

// demoStandardFor 演示标准三段式循环
func demoStandardFor() {
	fmt.Println("\n>>> 1. 标准 For 循环")

	sum := 0
	// for init; condition; post { ... }
	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Printf("1 到 5 的和为: %d\n", sum)
}

// demoWhileStyle 演示 While 风格循环
func demoWhileStyle() {
	fmt.Println("\n>>> 2. While 风格循环")

	n := 3
	// 相当于 while (n > 0)
	for n > 0 {
		fmt.Printf("倒计时: %d\n", n)
		n--
	}
	fmt.Println("发射!")
}

// demoInfiniteLoop 演示无限循环与控制语句
func demoInfiniteLoop() {
	fmt.Println("\n>>> 3. 无限循环与 break/continue")

	count := 0
	// 相当于 while (true)
	for {
		count++

		// continue: 跳过本次循环剩余代码，进入下一次
		if count%2 == 0 {
			// fmt.Printf("%d 是偶数，跳过\n", count)
			continue
		}

		fmt.Printf("处理奇数: %d\n", count)

		// break: 终止循环
		if count >= 5 {
			fmt.Println("达到停止条件，退出循环")
			break
		}
	}
}

// demoForRange 演示 For-Range 循环 (重点)
func demoForRange() {
	fmt.Println("\n>>> 4. For-Range 循环 (遍历集合)")

	// --- 1. 遍历数组/切片 ---
	nums := []int{10, 20, 30}
	fmt.Print("遍历切片: ")
	for i, v := range nums {
		fmt.Printf("[%d:%d] ", i, v)
	}
	fmt.Println()

	// --- 2. 遍历 Map ---
	m := map[string]string{"a": "Apple", "b": "Banana"}
	fmt.Print("遍历 Map: ")
	for k, v := range m {
		fmt.Printf("%s->%s ", k, v)
	}
	fmt.Println()

	// --- 3. 遍历字符串 ---
	str := "Hello, 世界"
	fmt.Println("遍历字符串 (按 rune 字符):")
	for i, char := range str {
		fmt.Printf("索引:%d 字符:%c (Unicode: %d)\n", i, char, char)
	}
}

// demoLabelBreak 演示带标签的 Break (用于跳出多层循环)
func demoLabelBreak() {
	fmt.Println("\n>>> 5. 带标签的 Break (多层跳出)")

	// 定义标签
OuterLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i=%d, j=%d\n", i, j)
			if i == 2 && j == 2 {
				fmt.Println("触发条件，直接跳出最外层循环!")
				break OuterLoop // 如果只有 break，只能跳出内层 j 循环
			}
		}
	}
	fmt.Println("已跳出 OuterLoop")
}
