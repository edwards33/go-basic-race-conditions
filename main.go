package main

import (
	"fmt"
	"sync"
        "runtime"
)

var wg sync.WaitGroup
var counter int

func main() {
        fmt.Println("Num of CPUs:", runtime.NumCPU())
	wg.Add(2)
	go increment("Proc01:")
	go increment(" Proc02:")
	wg.Wait()
        fmt.Println("Final Count:", counter)
}

func increment(s string) {
	for i := 0; i < 50; i++ {
                x := counter
                x++
                counter = x
		fmt.Println(s, i, "Count:", counter)
	}
	wg.Done()
}

// go run -race main.go
