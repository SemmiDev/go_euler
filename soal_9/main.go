package main

import (
	"fmt"
	"time"
)

func main() {
	before := time.Now()
out:
	for a := 1; a < 1_000; a++ {
		for b := 1; b < 1_000; b++ {
			for c := 1; c < 1_000; c++ {
				if ((a*a)+(b*b) == (c * c)) && (a+b+c == 1_000) && ((a < b) && (b < c)) {
					product := a * b * c
					fmt.Printf("%d + %d + %d = 1000, product = %d\n", a, b, c, product)
					break out
				}
			}
		}
	}
	after := time.Now()
	fmt.Println("Waktu eksekusi", (after.Nanosecond()-before.Nanosecond())/1_000_000, "milli second")
}