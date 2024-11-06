package main

import "fmt"


func buatDeretGanjil(n, current int) {
	if current > n {
		return
	}

	if current % 2 != 0 {
		fmt.Print(current, " ")
	}

	buatDeretGanjil(n, current+1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N : ")
	fmt.Scan(&n)
	fmt.Println(n)
	buatDeretGanjil(n, 1)
	fmt.Println() 
}
