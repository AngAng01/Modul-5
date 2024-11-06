package main

import (
	"fmt"
)

func buatUrutan(n, current int) {
	fmt.Print(current, " ")

	if current == 1 {
		return
	}

	buatUrutan(n, current-1)

	fmt.Print(current, " ")
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N : ")
	fmt.Scan(&n)
	fmt.Println(n)
	buatUrutan(n, n)
	fmt.Println() 
}
