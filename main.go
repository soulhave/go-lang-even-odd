package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, d := range slice {
		result := "odd"
		if d%2 == 0 {
			result = "even"
		}

		fmt.Println(d, "is", result)
	}

}
