package main

import "fmt"

func compare(array1 []string, array2 []string, value int) bool {
	same := true
	for i := 0; i < value; i++ {
		if array1[i] != array2[i] {
			same = false
			break
		}
	}
	return same
} 
func main() {    
    array1 := make([]string, 3)//dideklare dengan value=3 sesuai soal
    array2 := make([]string, 3)//dideklare dengan value=3 sesuai soal

    fmt.Print("Array 1: ")
    for i := 0; i < 3; i++ {
        fmt.Scan(&array1[i])
    }
    fmt.Print("Array 2: ")
    for i := 0; i < 3; i++ {
        fmt.Scan(&array2[i])
    }
    if compare(array1, array2, 3) {
		fmt.Println("Kedua array tersebut isinya yang sama!")
	} else {
		 for i := 0; i < len(array1); i++ {
			if array1[i] != array2[i] {
				fmt.Printf("Index ke %d berbeda\n", i)
			}
		}
	}
}