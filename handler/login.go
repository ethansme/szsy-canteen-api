package handler

import (
	"fmt"
	"time"
	"regexp"
	"strings"
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"net/http/cookiejar"
	. "util"
)

type Info struct {
	Name    string `json:"姓名"`
	Money   string `json:"余额"`
	Cookies string `json:"口令"`
}

func LoginHandler(w http.ResponseWriter, request *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	var client http.Client

	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client.Jar = jar

	// GET CAS Login Page

	response, err := client.Get(LOGIN)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	// Take the value of 'lt' and 'execution' from the CAS Login Page source

	var (
		lt        string
		execution string
	)

	Waitgroup.Add(1)
	go Reg(&lt, "lt\" value=\"(.*)\"", string(body))
	Waitgroup.Add(1)
	go Reg(&execution, "execution\" value=\"(.*)\"", string(body))

	// Take values in the form from request

	request.ParseForm()
	username := request.FormValue("账号")
	password := request.FormValue("密码")

	if username == "" && password == "" {
		fmt.Fprintf(w, "缺少账号或密码")
		return
	}
	
	return
}