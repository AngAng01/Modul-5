# MODUL EMPAT
 ## SOAL 1
   Program di atas adalah program yang menghitung permutasi dan kombinasi dari dua pasang bilangan yang dimasukkan oleh pengguna.
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Fungsi 'factorial(n int, hasil *int)', Fungsi ini menghitung faktorial dari bilangan n dan menyimpan hasilnya dalam variabel yang ditunjuk oleh pointer hasil. Faktorial dihitung menggunakan loop dari n hingga 1.
      - Fungsi 'permutasi(n, r int) int', Fungsi ini menghitung permutasi dari n elemen yang diambil sebanyak r. Rumus yang digunakan adalah P(n,r) = n!/(n-r)!, dimana faktorial dari n dan n-r dihitung menggunakan fungsi factorial.
      - Fungsi 'kombinasi(n, r int) int', Fungsi ini menghitung kombinasi dari n elemen yang diambil sebanyak r. Rumus yang digunakan adalah C(n,r) = n!/r!(n-r)!, di mana faktorial dari n, r, dan n-r dihitung menggunakan fungsi factorial.
      
   ## Code Explanation
      ```go
          func factorial(n int, hasil *int) {
	            *hasil = 1 
	            for i := n; i >= 1; i-- {
		             *hasil *= i
	            }
          }
      ```
   Kode di atas adalah kode fungsi yang digunakan untuk menghitung faktorial dari sebuah bilangan bulat n. 
   
   ```go
      func permutasi(n, r int) int {
      	var pembilang, penyebut int
      	factorial(n, &pembilang)
      	factorial(n-r, &penyebut)
      	return pembilang / penyebut
      }
   ```
   Kode di atas adalah Fungsi permutasi yang digunakan untuk menghitung permutasi dari n elemen yang diambil sebanyak r. Fungsi ini pertama-tama menghitung faktorial dari n (disimpan dalam variabel pembilang) dan faktorial dari n - r (disimpan dalam variabel penyebut) dengan memanfaatkan fungsi factorial. 
   
   ```go
      func kombinasi(n, r int) int {
      	var pembilang, penyebut1, penyebut2 int
      	factorial(n, &pembilang)               
      	factorial(r, &penyebut1)               
      	factorial(n-r, &penyebut2)             
      	return pembilang / (penyebut1 * penyebut2) 
      }
   ```
   Kode di atas adalah fungsi yang digunakan untuk menghitung kombinasi dari n elemen yang diambil sebanyak r. 
   
   ```go
      var a, b, c, d int
   ```
   kode diatas adalah kode yang digunakan untuk mendeklarasikan 4 variabel int yang digunakan sebagai inputan.
   
   ```go
      fmt.Scan(&a, &b, &c, &d)
   ```
   kode diatas adalah kode yang digunakan untuk pengguna menginputkan suatu nilai untuk dieksekusi.
   
   ```go
      fmt.Println(permutasi(a, c), kombinasi(a, c))
	  fmt.Println(permutasi(b, d), kombinasi(b, d))
   ```
   Kode di atas adalah bagian dari fungsi main() dalam program Go yang berfungsi untuk mencetak hasil perhitungan permutasi dan kombinasi dari dua pasangan bilangan (a, c dan b, d) menggunakan fungsi fmt.Println().
