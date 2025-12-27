package main

import "fmt"

// main 函数：数组 (Array) 基础与进阶
// 注意：数组是固定长度的，一旦声明，长度不可改变。
// 如果需要动态长度，请使用切片 (Slice)。
func main() {
	fmt.Println("========== Golang 数组 (Array) 运用示例 ==========")

	// 1. 数组的声明与初始化
	demoDeclaration()

	// 2. 数组的遍历
	demoTraversal()

	// 3. 数组是值类型 (Value Type)
	demoValueType()

	// 4. 多维数组
	demoMultiArray()

	fmt.Println("\n========== 示例结束 ==========")
}

// demoDeclaration 演示数组的声明与初始化
func demoDeclaration() {
	fmt.Println("\n>>> 1. 数组声明与初始化")

	// 方式一：声明指定长度，默认初始化为零值
	var arr1 [5]int
	var ages [3]string
	ages = [3]string{"18", "25", "64"}
	fmt.Println(ages)
	arr1[0] = 100
	fmt.Printf("方式一 (默认零值): %v, 长度: %d\n", arr1, len(arr1))

	// 方式二：声明并初始化
	var arr2 = [3]string{"Golang", "Java", "Python"}
	var names [5]string = [5]string{"张三", "李四", "王五", "赵六", "王二"}
	fmt.Println(names)
	fmt.Printf("方式二 (直接初始化): %v\n", arr2)

	// 方式三：自动推导长度 [...]
	arr3 := [...]int{1, 2, 3, 4, 5, 6}
	var schools = [...]string{"清华大学", "北京大学", "中国人民大学", "北京航空航天大学", "北京理工大学"}
	for index, value := range schools {
		fmt.Printf("学校[%d]: %s\n", index, value)
	}
	fmt.Printf("方式三 (自动推导长度): %v, 推导出的长度: %d\n", arr3, len(arr3))
	grades := [...]int{8, 64, 72, 64}
	for i := 0; i < len(grades); i++ {
		fmt.Printf("成绩[%d]: %d\n", i, grades[i])
	}

	// 方式四：指定索引初始化
	// 索引 1 为 10, 索引 3 为 30, 其余为 0
	arr4 := [...]int{1: 10, 3: 30}
	fmt.Printf("方式四 (指定索引): %v\n", arr4)
}

// demoTraversal 演示数组的遍历
func demoTraversal() {
	fmt.Println("\n>>> 2. 数组遍历")
	// 自动推导长度写法;
	arr := [...]string{"Apple", "Banana", "Cherry"}

	// 方式一：传统 for 循环
	fmt.Print("传统 for 循环: ")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%s ", arr[i])
	}
	fmt.Println()

	// 方式二：for-range (推荐)
	fmt.Print("for-range 循环: ")
	for index, value := range arr {
		fmt.Printf("[%d:%s] ", index, value)
	}
	fmt.Println()

	// 如果只需要值，可以忽略索引
	// for _, value := range arr { ... }
}

// demoValueType 演示数组是值类型 (重要特性)
func demoValueType() {
	fmt.Println("\n>>> 3. 数组的值类型特性")

	// 定义一个原数组
	original := [3]int{1, 2, 3}

	// 将数组赋值给新变量 (发生完全拷贝)
	copyArr := original

	// 修改新数组
	copyArr[0] = 999

	fmt.Printf("原数组: %v (未变)\n", original)
	fmt.Printf("副本数组: %v (已变)\n", copyArr)
	fmt.Println("结论: Golang 数组赋值是值拷贝，修改副本不影响原数组。")

	// 演示函数传参也是值拷贝
	changeArray(original)
	fmt.Printf("函数调用后原数组: %v (依然未变)\n", original)
}

// changeArray 尝试修改数组 (但只会修改副本)
func changeArray(arr [3]int) {
	arr[0] = 888
	// fmt.Println("函数内部修改:", arr)
}

// demoMultiArray 演示多维数组
func demoMultiArray() {
	fmt.Println("\n>>> 4. 多维数组")

	// 声明一个 2行 3列 的二维数组
	var grid [2][3]int

	// 初始化赋值
	grid[0] = [3]int{1, 2, 3}
	grid[1] = [3]int{4, 5, 6}

	// 简写初始化
	// grid := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println("二维数组内容:")
	for i, row := range grid {
		fmt.Printf("第 %d 行: ", i)
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}
