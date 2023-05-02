/*Ellipsis dalam pemrograman Go adalah operator variadic yang digunakan untuk mengirimkan 
argumen tanpa jumlah tetap ke fungsi.  Operator ini ditandai dengan tiga titik (...), 
yang ditempatkan di depan tipe data parameter terakhir dalam daftar parameter fungsi.

Dalam konteks ini, variadic artinya dapat menerima sejumlah argumen tanpa batas dalam bentuk 
slice, dan akan dibaca sebagai sebuah array. Variadic function adalah sebuah fungsi yang 
menerima jumlah parameter yang tidak pasti atau bisa dibilang bisa menerima berapa pun argumen 
yang diinginkan oleh pengguna.
*/

package main

import "fmt"

func addNumbers(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main () {
	result := addNumbers(1, 2, 3, 4, 5)
	fmt.Println(result) 
	// Output: 15
}