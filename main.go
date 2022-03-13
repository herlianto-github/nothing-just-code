package main

import "fmt"

func staircase(n int) {
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
	fmt.Println("-------")
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
	fmt.Println("-------")
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		for j := i + 1; j <= n; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
	fmt.Println("-------")
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

}

func main() {
	fmt.Println("nothing just code")
	staircase(6)
}
