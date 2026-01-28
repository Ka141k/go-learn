package main

import (
	"fmt"
)

func main() {
	arr := [5]int{5, 66, 7, 100, 1}

	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			arr[i] *= 2
		}

		fmt.Println(i, "-", arr[i])

	}

	// fmt.Println("Результат цикла:", arr)
}
