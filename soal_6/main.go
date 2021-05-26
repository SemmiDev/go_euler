package main

import (
	"fmt"
	"time"
)

func main() {
	before := time.Now()

	sumSquares := 0
	numberOfData := 100_000
	squaresOfTheSum:=0

	sumSquares = (numberOfData * (numberOfData + 1) * (2 * numberOfData + 1)) / 6
	sumData := (numberOfData*(numberOfData+1))/2
	squaresOfTheSum = sumData*sumData
	diffSums :=  squaresOfTheSum - sumSquares

	fmt.Println("Selisis Jumlah Kuadrat Adalah : ",diffSums)
	after := time.Now()
	fmt.Println("Waktu eksekusi", after.Nanosecond()-before.Nanosecond(), "nano second")
}