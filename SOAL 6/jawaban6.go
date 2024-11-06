package main

import "fmt"

func Pangkat(n1, n2 int) int {
	if n2 == 0 {
		return 1 
	}
	return n1 * Pangkat(n1, n2-1) 
}

func main() {
	var n1, n2 int

	fmt.Scan(&n1, &n2)
	
	fmt.Println("Hasil :", Pangkat(n1, n2))
}
