package main

import "fmt"

// 定义一个"说话"接口：任何能 Speak 的东西都算它
type Speaker interface {
    Speak() string
}

// 定义两个"东西"：狗 和 人
type Dog struct {
    Name string
}  

type Person struct {
    Name string
}

// 让狗实现 Speak 方法
func (d Dog) Speak() string {
    return "汪汪！我是" + d.Name
}

// 让人实现 Speak 方法
func (p Person) Speak() string {
    return "你好，我是" + p.Name
}

func main() {
    var s Speaker

    s = Dog{Name: "旺财"}
    fmt.Println(s.Speak())

    s = Person{Name: "陈静一"}
    fmt.Println(s.Speak())
}

