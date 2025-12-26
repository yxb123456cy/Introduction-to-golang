package main

import (
	"LuFei_go_study/constants"
	"fmt"
)

/**
 * @Author: 轻叶
 * @Desc: 可变全局变量;
 */
var AGE = 19
var SCHOOL = "南昌航空大学"

/**
 * @Author: 轻叶
 * @Desc: 块范围可变全局变量;
 */
var (
	DB       string = "mysql"
	PORT     int    = 3306
	USERNAME        = "root"
	PASSWORD        = "123456"
)

/**
 * @Author: 轻叶
 * @Desc: 常量块定义;
 */
const (
	DB_URL = "jdbc:mysql://localhost:3306/test"
)

// 常量定义 必须初始化;
const DB_VERSION float64 = 8.030
const DB_TYPE = "postGreSQL"
const DB_NAME = "testDB"

func main() {
	var name = "docker"
	elk := "hello,world"
	docker := "docker+K8s"
	k8s := "k8s"
	fmt.Println(k8s)
	fmt.Println(elk)
	fmt.Println(docker)
	fmt.Println(name)
	mybatis := "学习框架"
	fmt.Println(mybatis)
	var spring string = "SpringBoot"
	fmt.Println(spring)
	fmt.Println("你好Golang")
	fmt.Println(constants.DOCKER)
	// 1.先声明并带类型后赋值
	var nacos string
	nacos = "NACOS配置中心"
	fmt.Println(nacos)
	// 2.先声明并带类型后赋值
	var dubbo string = "Dubbo框架"
	fmt.Println(dubbo)
	// 3.最简短且省略类型的写法 只能在函数作用域中才可以使用该写法;
	zookeeper := "zookeeper"
	fmt.Println(zookeeper)
	fmt.Println(AGE)
	fmt.Println(SCHOOL)
	AGE = 38
	fmt.Println(AGE)
	constants.GOLANG_VERSION = "3.6"
	fmt.Printf("当前使用的golangSDK版本为:%s", constants.GOLANG_VERSION)
	var stack string
	fmt.Scan(&stack)
	fmt.Printf("你的技术栈为: %s", stack)

}
