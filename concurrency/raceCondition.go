package concurrency

import (
	"fmt"
	_ "fmt"
	"time"
	_ "time"
)

func RaceCondition() {
	var msg = "hello i am thread"
	go func(msg string) {
		fmt.Println(msg)
	}(msg)
	msg = "hello i am main"
	time.Sleep(100 * time.Millisecond)
}
