package channel

import (
	"fmt"
	_ "fmt"
	"sync"
	_ "sync"
	"time"
	_ "time"
)

var wg = sync.WaitGroup{}

func BasicChannel() {
	fmt.Println("code started")
	ch := make(chan int)
	wg.Add(2)
	go func() {
		fmt.Println("waiting")
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		time.Sleep(10 * time.Second)
		ch <- 42
		wg.Done()
	}()
	wg.Wait()
}
