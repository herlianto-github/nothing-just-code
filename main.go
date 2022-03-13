package main

import (
	"fmt"
	"sort"
)

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

func miniMaxSum(arr []int32) {
	var max int64
	var min int64
	arrint := []int{}
	for i := 0; i < len(arr); i++ {
		arrint = append(arrint, int(arr[i]))
	}
	sort.Ints(arrint)
	for i := 0; i < len(arrint)-1; i++ {
		min += int64(arrint[i])
	}
	for i := 1; i < len(arrint); i++ {
		max += int64(arrint[i])
	}
	fmt.Println("min", min, "max", max)
	fmt.Println("-----------------")
}

func main() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("nothing just code")
	staircase(6)
	fmt.Println("=========================================================")
	miniMaxSum([]int32{1, 2, 3, 4, 5})
	miniMaxSum([]int32{7, 69, 2, 221, 8974})
	miniMaxSum([]int32{256741038, 623958417, 467905213, 714532089, 938071625})

}
