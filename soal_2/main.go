package main

import (
	"fmt"
	"time"
)
func main() {
	start := time.Now()

	firstNumber := 0
	secondNumber := 1
	sumEven,sumOdd := 0,0
	const max = 4_000_000

	for secondNumber < max {
		buff := firstNumber + secondNumber
		if secondNumber % 2 == 0 {
			fmt.Println("EVEN FROM FIB :", secondNumber)
			sumEven += secondNumber
		} else {
			fmt.Println("ODD FROM FIB :", secondNumber)
			sumOdd += secondNumber
		}

		firstNumber = secondNumber
		secondNumber = buff
	}

	fmt.Printf("Jumlah dari bilangan ganjil : [%d]\n", sumOdd)
	fmt.Printf("Jumlah dari bilangan genap : [%d]\n", sumEven)

	finish := time.Now()
	fmt.Println("Waktu eksekusi ", finish.Nanosecond() - start.Nanosecond(), "nano second")
}