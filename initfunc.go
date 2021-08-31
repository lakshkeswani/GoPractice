package main

import (
	_ "fmt"
	"time"
	_ "time"
)

var weekday string

func init() {
	weekday = time.Now().Weekday().String()
}

func initfunc() {
	println("today is %s", weekday)
}
