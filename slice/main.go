package main

import "fmt"

// main 函数：切片 (Slice) 基础与进阶
// 切片是对数组的抽象，提供动态大小的、灵活的视图
func main() {
	fmt.Println("========== Golang 切片 (Slice) 运用示例 ==========")

	// 1. 切片的创建与初始化
	demoCreate()

	// 2. 切片的基本操作 (Append, Len, Cap)
	demoOperations()

	// 3. 切片是引用类型 (Reference Type)
	demoReference()

	// 4. 切片的拷贝 (Copy)
	demoCopy()

	fmt.Println("\n========== 示例结束 ==========")
}

// demoCreate 演示切片的创建方式
func demoCreate() {
	fmt.Println("\n>>> 1. 切片创建与初始化")

	// 方式一：直接字面量 (常用)
	// 注意：[]int 中括号内为空，表示切片；如果有数字则是数组
	s1 := []int{1, 2, 3}
	fmt.Printf("方式一 (字面量): %v, len=%d, cap=%d\n", s1, len(s1), cap(s1))

	// 方式二：使用 make 函数
	// make([]Type, len, cap)
	s2 := make([]int, 3, 5) // 长度为3，容量为5
	s2[0] = 10
	s2[1] = 20
	// s2[3] = 40 // 报错：index out of range (只能访问到 len-1)
	fmt.Printf("方式二 (make): %v, len=%d, cap=%d\n", s2, len(s2), cap(s2))

	// 方式三：从数组截取
	arr := [5]int{10, 20, 30, 40, 50}
	s3 := arr[1:4] // 索引 [1, 4) -> 包含索引 1, 2, 3
	fmt.Printf("方式三 (从数组截取): %v, len=%d, cap=%d\n", s3, len(s3), cap(s3))
}

// demoOperations 演示切片的基本操作
func demoOperations() {
	fmt.Println("\n>>> 2. 切片基本操作")

	var nums []int // 声明一个 nil 切片
	fmt.Printf("初始状态: %v, len=%d, cap=%d, isNil=%v\n", nums, len(nums), cap(nums), nums == nil)

	// --- Append 追加元素 ---
	fmt.Println("--- 开始追加元素 ---")
	for i := 1; i <= 5; i++ {
		nums = append(nums, i)
		// 观察容量 cap 的变化 (Go 会自动扩容，通常是成倍增加)
		fmt.Printf("追加 %d -> len=%d, cap=%d\n", i, len(nums), cap(nums))
	}

	// --- 截取 (Slicing) ---
	subSlice := nums[1:3] // 取索引 1 到 2
	fmt.Printf("截取 [1:3]: %v\n", subSlice)
}

// demoReference 演示切片的引用特性 (重要!)
func demoReference() {
	fmt.Println("\n>>> 3. 切片引用特性 (Reference Type)")

	// 定义一个切片
	original := []int{1, 2, 3}

	// 赋值给新变量 (切片底层共享同一个数组)
	refSlice := original

	// 修改新切片
	refSlice[0] = 999

	fmt.Printf("原切片: %v (被修改了!)\n", original)
	fmt.Printf("引用切片: %v\n", refSlice)
	fmt.Println("结论: 切片是引用传递，修改副本会影响原切片(因为它们指向底层同一个数组)。")
}

// demoCopy 演示切片的深拷贝
func demoCopy() {
	fmt.Println("\n>>> 4. 切片拷贝 (copy)")

	src := []int{1, 2, 3}
	// 目标切片必须预先分配好足够的长度 (len)
	dst := make([]int, len(src))

	// 使用 copy 函数进行深拷贝 (值复制)
	count := copy(dst, src)

	// 修改副本
	dst[0] = 888

	fmt.Printf("拷贝数量: %d\n", count)
	fmt.Printf("源切片: %v (未受影响)\n", src)
	fmt.Printf("目标切片: %v (修改生效)\n", dst)
	fmt.Println("结论: 使用 copy 函数可以实现独立副本，互不影响。")
}
