package main

import (
	"runtime"
	"net/http"
	. "handler"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	http.HandleFunc("/login/", LoginHandler)
	http.HandleFunc("/dates/", DatesHandler)
	http.HandleFunc("/menu/",  MenuHandler)
	http.HandleFunc("/order/", OrderHandler)
	http.ListenAndServe(":2018", nil)
}