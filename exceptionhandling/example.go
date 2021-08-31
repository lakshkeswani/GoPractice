package exceptionhandling

import (
	"fmt"
	_ "fmt"
	"io/ioutil"
	_ "io/ioutil"
	"log"
	_ "log"
	"net/http"
	_ "net/http"
)

func Deferexp() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robot, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("%s", robot)
}
