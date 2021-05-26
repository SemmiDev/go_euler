package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	sumOf3, sumOf5 := 0,0

	for i := 1; i < 1000; i++ {
		if i % 5 == 0 {
			sumOf5 += i
		} else if i % 3 == 0 {
			sumOf3 += i
		} else {}
	}

	fmt.Println("total kelipatan 3 adalah  :", sumOf3)
	fmt.Println("total kelipatan 5  adalah :", sumOf5)
	fmt.Println("total kelipatan 3 dan 5   :", sumOf5 + sumOf3)

	finish := time.Now()
	fmt.Println("time elapsed : ", finish.Nanosecond() - start.Nanosecond(), " ns")
}