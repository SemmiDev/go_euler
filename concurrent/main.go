package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	result := addWithConcurrency([]int{1,2,3,4,5,6,7,8})
	fmt.Println(result)
}

func addWithConcurrency(numbers []int) int64 {
	numberOfCores := runtime.NumCPU() // 2
	fmt.Println(numberOfCores) // 2
	runtime.GOMAXPROCS(numberOfCores) // 2

	var sum int64
	max := len(numbers) // 8

	sizeOfPart := max / numberOfCores // 8 / 2 = 4

	var wg sync.WaitGroup
	for i := 0; i < numberOfCores; i++ {
		start := i * sizeOfPart
		end := start + sizeOfPart
		part := numbers[start:end]
		wg.Add(1)
		go func(nums []int) {
			defer wg.Done()
			var partSum int64
			for _, n := range nums {
				partSum += int64(n)
			}
			atomic.AddInt64(&sum, partSum)
		}(part)
	}
	wg.Wait()
	return sum
}

