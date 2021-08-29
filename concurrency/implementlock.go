package concurrency

import (
	"fmt"
	_ "fmt"
	"runtime"
	_ "runtime"
	"sync"
	_ "sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func ImplementLock() {
	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayhello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayhello() {
	fmt.Println("counter =", counter)
	m.RUnlock()
	wg.Done()
}
func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
