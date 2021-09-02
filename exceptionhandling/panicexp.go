package exceptionhandling

import (
	_ "fmt"
	"net/http"
	_ "os/user"
)

func PanicExp() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello Go"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}
