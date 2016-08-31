// Selection Sort demonstration in Go

package main

import "fmt"
import "time"

func main() {
	start := time.Now()

	a := []int{6, 5, 4, 3, 2, 1}
	SelectionSort(a)
	fmt.Println(a)

	fmt.Println(time.Since(start))
}

func SelectionSort(a []int) {
	// iterate over the entire array
	for key, _ := range a {
		// iterate over the array starting from 1 + key
		for i := key + 1; i < len(a); i++ {
			// if an element is less than the key swap them
			if a[i] < a[key] {
				a[i], a[key] = a[key], a[i]
			}
		}
	}
}
