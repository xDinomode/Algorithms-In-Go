// Insertion sort demonstration in Go

package main

import "fmt"
import "time"

func main() {
	start := time.Now()

	a := []int{5, 2, 4, 6, 1, 3}
	InsertionSort(a)
	fmt.Println(a)

	fmt.Println(time.Since(start))
}

func InsertionSort(a []int) {
	// iterate over the slice and store the index in key
	for key, _ := range a {
		// skip the first index since it's already sorted
		if key == 0 {
			continue
		}

		// iterate over the sorted list
		for j := (key - 1); j >= 0; j-- {
			// compare j with j + 1 and swap values if
			// j + 1 is less than j
			if a[j+1] < a[j] {
				a[j+1], a[j] = a[j], a[j+1]
			}
		}
	}
}
