package main

import (
	"encoding/json"
	"fmt"
)

// 1. 定义结构体
// 结构体是用户自定义的类型，用于将不同类型的数据组合在一起
type Person struct {
	Name string // 字段名大写表示公开 (Public)，可被其他包访问
	Age  int    // 字段名小写表示私有 (Private)，仅当前包可见
	city string // 私有字段，外部无法直接访问
}

// 4. 结构体嵌套 (组合 Composition)
// Golang 没有继承，通过匿名结构体字段实现组合
type Student struct {
	Person // 匿名字段，Student 直接拥有 Person 的所有字段和方法
	School string
	Grade  int
}

// 5. 结构体 Tag (元数据)
// 常用于 JSON 序列化、ORM 映射等
type User struct {
	ID       int    `json:"id"`                 // JSON 输出时 key 为 "id"
	Username string `json:"username,omitempty"` // 如果为空则忽略
	Password string `json:"-"`                  // 永远不输出到 JSON
}

func main() {
	fmt.Println("========== Golang 结构体 (struct) 规则与示例 ==========")

	// 2. 结构体实例化
	demoInstantiation()

	// 3. 结构体方法 (Method)
	demoMethods()

	// 4. 结构体嵌套与组合
	demoEmbedding()

	// 5. 结构体 Tag 与 JSON
	demoTags()

	fmt.Println("\n========== 示例结束 ==========")
}

// demoInstantiation 演示结构体的创建方式
func demoInstantiation() {
	fmt.Println("\n>>> 2. 结构体实例化")

	// 方式一：键值对初始化 (推荐)
	p1 := Person{
		Name: "Alice",
		Age:  25,
		city: "New York", // 私有字段在同包内可以访问
	}
	fmt.Printf("方式一: %+v\n", p1)

	// 方式二：顺序初始化 (不推荐，字段变动易出错)
	p2 := Person{"Bob", 30, "London"}
	fmt.Printf("方式二: %+v\n", p2)

	// 方式三：new 关键字 (返回指针)
	p3 := new(Person)
	p3.Name = "Charlie" // Go 语法糖：自动解引用，等同于 (*p3).Name
	fmt.Printf("方式三 (指针): %+v, 类型: %T\n", p3, p3)

	// 方式四：取地址符 & (返回指针，最常用)
	p4 := &Person{Name: "David", Age: 20}
	fmt.Printf("方式四 (指针): %+v\n", p4)
}

// 3. 定义方法 (Receiver)
// (p Person) 是值接收者，修改不会影响原对象
func (p Person) SayHello() {
	fmt.Printf("Hello, my name is %s, I am %d years old.\n", p.Name, p.Age)
}

// (p *Person) 是指针接收者，修改会影响原对象 (推荐用于修改状态或大对象)
func (p *Person) SetCity(newCity string) {
	p.city = newCity
}

func demoMethods() {
	fmt.Println("\n>>> 3. 结构体方法 (Receiver)")
	p := Person{Name: "Eve", Age: 18, city: "Paris"}

	p.SayHello() // 调用值方法

	fmt.Println("修改前城市:", p.city)
	p.SetCity("Tokyo") // 调用指针方法 (Go 自动取地址)
	fmt.Println("修改后城市:", p.city)
}

func demoEmbedding() {
	fmt.Println("\n>>> 4. 结构体嵌套 (组合)")
	// 初始化嵌套结构体
	s := Student{
		Person: Person{
			Name: "Frank",
			Age:  15,
		},
		School: "High School",
		Grade:  10,
	}

	// 可以直接访问匿名字段的属性 (提升字段)
	fmt.Printf("学生姓名: %s (直接访问)\n", s.Name)
	fmt.Printf("学生年龄: %d\n", s.Age)
	fmt.Printf("学校: %s\n", s.School)

	// 也可以调用 Person 的方法
	s.SayHello()
}

func demoTags() {
	fmt.Println("\n>>> 5. 结构体 Tag (JSON 序列化)")

	u := User{
		ID:       101,
		Username: "golang_fan",
		Password: "secret_password",
	}

	// 序列化为 JSON
	jsonData, err := json.Marshal(u)
	if err != nil {
		fmt.Println("JSON 错误:", err)
		return
	}

	fmt.Printf("原始对象: %+v\n", u)
	fmt.Printf("JSON 结果: %s\n", string(jsonData))
	fmt.Println("(注意: Password 字段被忽略，ID 变成了小写 id)")
}
