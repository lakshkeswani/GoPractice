package concurrency

import (
	"fmt"
	_ "fmt"
	_ "sync"
)

//var  wg = sync.WaitGroup{}
func SolveRaceCondition() {
	var msg = "hello i am thread"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)

	msg = "hello i am main"
	wg.Wait()
}
