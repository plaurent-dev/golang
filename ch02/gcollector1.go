package main

import (
	"fmt"
	"runtime"
	"time"
)

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func printStats() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc :", bToMb(mem.Alloc))
	fmt.Println("mem.TotalAlloc :", bToMb(mem.TotalAlloc))
	fmt.Println("mem.HeapAlloc :", bToMb(mem.HeapAlloc))
	fmt.Println("mem.NumGC :", mem.NumGC)
	fmt.Println("____________________________")
}

func main() {

	printStats()

	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation échouée")
		}
	}

	printStats()

	for i := 0; i < 10; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation échouée")
		}
		time.Sleep(5 * time.Second)
	}

	printStats()

}
