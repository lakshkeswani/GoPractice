package channel

import (
	"fmt"
	_ "fmt"
	_ "sync"
)

func BufferChannel() {
	ch := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		i = <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
	}(ch)
	wg.Wait()
}
