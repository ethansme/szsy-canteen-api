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
	go Find(&lt, "lt\" value=\"(.*)\"", string(body))
	Waitgroup.Add(1)
	go Find(&execution, "execution\" value=\"(.*)\"", string(body))

	// Take values in the form from request

	request.ParseForm()
	username := request.FormValue("账号")
	password := request.FormValue("密码")

	if username == "" && password == "" {
		fmt.Fprintf(w, "缺少账号或密码")
		return
	}

	// Make the form for login post

	data := make(url.Values)
	data["username"] = []string{username}
	data["password"] = []string{password}

	Waitgroup.Wait()

	data["lt"]        = []string{lt}
	data["execution"] = []string{execution}
	data["_eventId"]  = []string{"submit"}

	// POST to login with the form above

	response, err = client.PostForm(LOGIN, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, _ = ioutil.ReadAll(response.Body)

	// Check the status

	success, _ := regexp.MatchString("当前用户", string(body))
	if success {
		fmt.Println(time.Now().Format("2006/01/02 15:04:05"), username, "登录成功")
	} else {
		fmt.Println(time.Now().Format("2006/01/02 15:04:05"), username, "登录失败")
		fmt.Fprintf(w, "用户名或密码错误")
		return
	}

	response, err = client.Get(CARD)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, _ = ioutil.ReadAll(response.Body)
	
	var (
		user  string
		money string
	)

	Waitgroup.Add(1)
	go Find(&user, "用户：(.*)</span></span><span style", string(body))
	Waitgroup.Add(1)
	go Find(&money, "余额：(.*)</span></span><span style", string(body))

	cookies_raw  := strings.Split(response.Request.Header["Cookie"][0], "=")
	SessionId    := cookies_raw[2]
	LcSoftCardV2 := strings.Split(cookies_raw[1], ";")[0]
	cookies      := LcSoftCardV2 + "," + SessionId

	return
}