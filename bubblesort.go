// Bubblesort demonstration in Go

package main

import "fmt"
import "time"

func main() {
	start := time.Now()

	a := []int{5, 6, 4, 3, 2, 1}
	BubbleSort(a)
	fmt.Println(a)

	fmt.Println(time.Since(start))
}

func BubbleSort(a []int) {
	// stores the amount of sorted elements
	j := 0
	// while elements are being swapped, continue another loop
	for swapped := true; swapped != false; {
		j++
		swapped = false
		// iterates over the array and swaps values
		for i := 0; i < len(a)-j; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				swapped = true
			}
		}
	}
}
