package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// main 函数是程序的入口
// 这里我们将演示 Golang 中常用的几种输入输出方式
func main() {
	fmt.Println("========== Golang 输入输出常用示例 ==========")

	// 1. 标准控制台输入输出 (fmt 包)
	// 适合简单的打印和格式化读取
	demoStdIO()

	// 2. 缓冲输入输出 (bufio 包)
	// 适合处理带空格的字符串输入、大块数据读写，效率更高
	demoBufio()

	// 3. 文件读写操作 (os 包)
	// 演示文件的创建、写入、读取
	demoFileIO()

	// 4. IO 工具接口 (io 包)
	// 演示 io.Copy 等常用工具
	demoIOTools()

	fmt.Println("\n========== 所有示例演示完毕 ==========")
}

// demoStdIO 演示 fmt 包的标准输入输出
func demoStdIO() {
	fmt.Println("\n>>> 1. 标准输入输出 (fmt)")

	// --- 输出 ---
	fmt.Print("fmt.Print: 不换行输出 ")
	fmt.Println("fmt.Println: 换行输出")
	name := "Golang"
	age := 14
	// %s 字符串, %d 整数, %v 自动推断值, %T 类型
	fmt.Printf("fmt.Printf: 格式化输出 -> Name: %s, Age: %d, Type: %T\n", name, age, name)

	// --- 输入 ---
	// 注意：fmt.Scan 系列以空格或换行作为分隔符
	/*
		var inputName string
		var inputAge int
		fmt.Println("\n[交互测试] 请输入姓名和年龄 (用空格隔开):")
		// fmt.Scan(&inputName, &inputAge)
		// fmt.Printf("你输入了: 姓名=%s, 年龄=%d\n", inputName, inputAge)
	*/
	fmt.Println("(注：为了演示流畅，fmt.Scan 交互部分已注释，取消注释可测试)")
}

// demoBufio 演示 bufio 包的缓冲输入输出
func demoBufio() {
	fmt.Println("\n>>> 2. 缓冲输入输出 (bufio)")

	// --- Reader: 读取整行 (包含空格) ---
	// 模拟一个输入流 (这里用 strings.NewReader 模拟键盘输入，方便演示)
	inputSource := "Hello Gopher World\nThis is line 2"
	reader := bufio.NewReader(strings.NewReader(inputSource))

	// ReadString 读取直到遇到指定分隔符
	line1, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取错误:", err)
	}
	// strings.TrimSpace 常用于去除末尾的换行符
	fmt.Printf("读取到的内容: %q\n", strings.TrimSpace(line1))

	// --- Writer: 缓冲写入 ---
	// 默认输出到标准输出 (屏幕)
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("这是通过 bufio.Writer 写入的内容 (在 Flush 前暂存内存中)...")
	// 必须调用 Flush() 才能将缓冲区内容真正写入目标
	writer.Flush()
	fmt.Println() // 补个换行
}

// demoFileIO 演示 os 包的文件操作
func demoFileIO() {
	fmt.Println("\n>>> 3. 文件读写操作 (os)")
	fileName := "test_io.txt"

	// --- 方式一：快速读写 (适合小文件) ---
	content := []byte("Hello, this is a test file.\nGolang IO is simple.")
	// 0644 是文件权限: rw-r--r--
	err := os.WriteFile(fileName, content, 0644)
	if err != nil {
		fmt.Println("写入文件失败:", err)
		return
	}
	fmt.Println("文件写入成功:", fileName)

	// 读取文件
	readContent, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("读取文件失败:", err)
		return
	}
	fmt.Printf("文件内容读取成功:\n---\n%s\n---\n", string(readContent))

	// --- 方式二：打开文件进行追加 (Append) ---
	// O_APPEND: 追加模式, O_CREATE: 不存在则创建, O_WRONLY: 只写
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer file.Close() // 确保文件关闭

	if _, err := file.WriteString("Appended Line: 追加的一行内容。\n"); err != nil {
		fmt.Println("追加内容失败:", err)
	} else {
		fmt.Println("追加内容成功")
	}

	// 清理演示文件 (可选)
	// os.Remove(fileName)
}

// demoIOTools 演示 io 包常用工具
func demoIOTools() {
	fmt.Println("\n>>> 4. IO 工具接口 (io.Copy)")

	// io.Copy 可以在两个流之间直接传输数据
	// 例如：将字符串读取器的数据，直接拷贝到标准输出
	src := strings.NewReader("我是通过 io.Copy 直接流转的数据\n")
	dst := os.Stdout

	if _, err := io.Copy(dst, src); err != nil {
		fmt.Println("io.Copy 失败:", err)
	}
}
