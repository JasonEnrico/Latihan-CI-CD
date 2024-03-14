package main

import (
	"fmt"
)

func a() {

	var n int

	fmt.Scan(&n)
	// Nomor i

	for i := 1; i <= (n+1)/2; i++ {
		for j := 1; j <= n-(i+1); j++ {
			fmt.Print(" ")
		}

		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	for i := 1; i <= (n-1)/2; i++ {
		for j := 1; j <= i+1; j++ {
			fmt.Print(" ")
		}

		for k := n - (2 * i); k >= 1; k-- {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	// Nomor ii

	for i := 1; i <= n; i++ {
		fmt.Print("*")
	}
	fmt.Println("")

	for i := 1; i <= n-2; i++ {
		fmt.Print("*")

		for j := 1; j <= n-2; j++ {
			fmt.Print(" ")
		}

		fmt.Print("*")

		fmt.Println("")
	}

	for i := 1; i <= n; i++ {
		fmt.Print("*")
	}
}

func b() {

	// var arr = [5]int{1, 99, 29, 14, 32}

	var arr [5]int

	for i := 0; i < len(arr); i++ {
		fmt.Printf("Input data %d: ", i)
		fmt.Scan(&arr[i])
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	max := arr[2:]

	// Tampilkan hasil
	fmt.Println("3 angka terbesar secara ascending:", max)
}
