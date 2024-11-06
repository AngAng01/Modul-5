package main

import (
    "fmt"
)

func Factors(n, i int) {
    if i > n {
        return
    }
    if n%i == 0 {
        fmt.Print(i, " ")
    }
    Factors(n, i+1)
}

func main() {
    var n int
    fmt.Print("Masukkan bilangan bulat positif : ")
    fmt.Scan(&n)
    fmt.Println(n)
    Factors(n, 1)
    fmt.Println()
}
