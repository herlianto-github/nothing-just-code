package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
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

func birthdayCakeCandles(candles []int32) int32 {

	arr := make(map[int32]int32)
	var result int32
	for i := 0; i < len(candles); i++ {
		arr[candles[i]]++
	}

	for _, value := range arr {
		if result < value {
			result = value
		}
	}

	return result

}

func timeConversion(s string) string {
	// Write your code here
	result := ""

	kode := s[len(s)-2:]
	waktu := s[:len(s)-2]
	arr := strings.Split(waktu, ":")
	jam, _ := strconv.Atoi(arr[0])

	if kode == "PM" && arr[0] != "12" {
		jam += 12
	}

	if kode == "AM" && arr[0] == "12" {
		jam = 0
	}

	result = fmt.Sprintf("%02v:%v:%v", jam, arr[1], arr[2])

	return result

}

func main() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("nothing just code")
	staircase(6)
	fmt.Println("=========================================================")
	miniMaxSum([]int32{1, 2, 3, 4, 5})
	miniMaxSum([]int32{7, 69, 2, 221, 8974})
	miniMaxSum([]int32{256741038, 623958417, 467905213, 714532089, 938071625})
	fmt.Println("=========================================================")
	fmt.Println(birthdayCakeCandles([]int32{3, 2, 1, 3}))
	fmt.Println("=========================================================")
	fmt.Println(timeConversion("07:05:45PM"))
	fmt.Println(timeConversion("12:40:22AM"))
}
