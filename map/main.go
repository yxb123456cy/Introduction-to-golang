package main

import "fmt"

// main 函数：Map (映射) 基础与进阶
// Map 是无序的键值对集合，Key 必须是支持 == 比较的类型 (slice, map, function 不可作为 Key)
func main() {
	fmt.Println("========== Golang Map (映射) 运用示例 ==========")

	// 1. Map 的声明与初始化
	demoDeclaration()

	// 2. Map 的基本操作 (增删改查)
	demoOperations()

	// 3. Map 的遍历
	demoTraversal()

	// 4. Map 的特性 (引用类型)
	demoReference()

	fmt.Println("\n========== 示例结束 ==========")
}

// demoDeclaration 演示 Map 的声明与初始化
func demoDeclaration() {
	fmt.Println("\n>>> 1. Map 声明与初始化")

	// 方式一：make(map[KeyType]ValueType)
	// 建议指定容量，避免频繁扩容
	m1 := make(map[string]int, 10)
	m1["age"] = 18
	fmt.Printf("方式一 (make): %v, len=%d\n", m1, len(m1))

	// 方式二：字面量初始化
	m2 := map[string]string{
		"name":   "Golang",
		"type":   "Language",
		"origin": "Google",
	}
	fmt.Printf("方式二 (字面量): %v, len=%d\n", m2, len(m2))

	// 注意：只声明不初始化，Map 为 nil，不能直接赋值！
	var m3 map[string]int
	fmt.Printf("未初始化 Map: %v, isNil=%v\n", m3, m3 == nil)
	// m3["test"] = 1 // 报错: assignment to entry in nil map
}

// demoOperations 演示 Map 的基本操作 (CURD)
func demoOperations() {
	fmt.Println("\n>>> 2. Map 基本操作 (增删改查)")

	scores := map[string]int{
		"Alice": 90,
		"Bob":   85,
	}
	fmt.Printf("初始状态: %v\n", scores)

	// --- 增加 / 修改 ---
	scores["Charlie"] = 95 // 增加
	scores["Alice"] = 99   // 修改 (Key 存在则更新)
	fmt.Printf("增加/修改后: %v\n", scores)

	// --- 删除 ---
	delete(scores, "Bob") // Key 不存在也不会报错
	fmt.Printf("删除 Bob 后: %v\n", scores)

	// --- 查询 (Comma OK idiom) ---
	// value, ok := map[key]
	val, ok := scores["Alice"]
	if ok {
		fmt.Printf("查询 Alice: 存在, 分数=%d\n", val)
	} else {
		fmt.Println("查询 Alice: 不存在")
	}

	val2, ok2 := scores["David"]
	if ok2 {
		fmt.Printf("查询 David: 存在, 分数=%d\n", val2)
	} else {
		fmt.Println("查询 David: 不存在 (返回零值)", val2)
	}
}

// demoTraversal 演示 Map 的遍历
func demoTraversal() {
	fmt.Println("\n>>> 3. Map 遍历")

	m := map[string]string{
		"CN": "China",
		"US": "United States",
		"JP": "Japan",
	}

	// 注意：Map 是无序的，每次遍历顺序可能不同
	for key, value := range m {
		fmt.Printf("Key: %-3s | Value: %s\n", key, value)
	}
}

// demoReference 演示 Map 的引用特性
func demoReference() {
	fmt.Println("\n>>> 4. Map 引用特性")

	original := map[string]int{"one": 1, "two": 2}

	// Map 是引用类型，赋值给新变量指向同一底层数据结构
	refMap := original

	refMap["one"] = 111

	fmt.Printf("原 Map: %v (被修改了)\n", original)
	fmt.Printf("引用 Map: %v\n", refMap)
}
