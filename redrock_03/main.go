package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	over := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
		wg.Wait()
		if i == 9 {
			over <- true
		}
	}
	<-over
	fmt.Println("over!!!")
}
