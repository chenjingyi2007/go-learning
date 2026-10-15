package main

import "fmt"

func main() {
    age := 20
    if age >= 18 {
        fmt.Println("成年了")
    } else {
        fmt.Println("未成年")
    }

    sum := 0
    for i := 1; i <= 100; i++ {
        sum = sum + i
    }
    fmt.Println("1加到100等于:", sum)
}

