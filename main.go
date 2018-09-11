package main

import (
	"fmt"
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

	fmt.Println(time.Now().Format("2006/01/02 15:04:05"), "szsy-canteen-api 开始运行")

	http.ListenAndServe(":2018", nil)
}