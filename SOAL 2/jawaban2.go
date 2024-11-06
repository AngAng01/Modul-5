package main

import (
    "fmt"
)

func cetakBintang(n int) {
    if n == 0 {
        return
    }
    fmt.Print("*")
    cetakBintang(n - 1)
}

func cetakPola(n, current int) {
    if current > n {
        return
    }
    cetakBintang(current)
    fmt.Println()
    cetakPola(n, current + 1)
}

func main() {
    var n int
    fmt.Scan(&n)
    cetakPola(n, 1)
}
