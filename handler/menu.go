package handler

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
	. "util"
)

type Menu struct {
	Notorder  []string   `json:"不订餐"`
	Breakfast [][]string `json:"早餐"`
	Lunch     [][]string `json:"午餐"`
	Dinner    [][]string `json:"晚餐"`
}

func MenuHandler(w http.ResponseWriter, request *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Take value of 'date' from the url

	var date string
	form, err := url.ParseQuery(request.URL.RawQuery)
	if err == nil && len(form["date"]) > 0 {
		date = form["date"][0]
	}

	var client http.Client

	// Take values in the form from request

	request.ParseForm()
	cookies      := request.FormValue("口令")
	LcSoftCardV2 := strings.Split(cookies, ",")[0]
	SessionId    := strings.Split(cookies, ",")[1]

	// Make Cookies

	cookie_CasModule    := http.Cookie {
		Name:  "LcSoftCardV2.CasModule",
		Value: "/card/",
	}
	cookie_LcSoftCardV2 := http.Cookie {
		Name:  ".LcSoftCardV2",
		Value: LcSoftCardV2,
	}
	cookie_SessionId    := http.Cookie {
		Name:  "ASP.NET_SessionId",
		Value: SessionId,
	}
	
	request, err = http.NewRequest("GET", MENU + date, nil)
	if err != nil {
		panic(err)
	}

	// Add Cookies

	request.AddCookie(&cookie_CasModule)
	request.AddCookie(&cookie_LcSoftCardV2)
	request.AddCookie(&cookie_SessionId)

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	// Check if Cookies are still alive

	re_check := regexp.MustCompile("<input id=\"username\" name=\"username\"")
	if len(re_check.FindStringSubmatch(string(body))) != 0 {
		fmt.Fprintln(w, "口令错误或过期")
		return
	}

	// Fetch menu

	var (
		meal      	 []string
		menu      	 [][]string
		breakfast 	 [][]string
		lunch     	 [][]string
		dinner    	 [][]string

		num          [][]string
		notorder     []string
		notorder_raw [][]int
		index        string
		category     string
		name         string
		price        string
		number       string
	)
	
	exist, _ := regexp.MatchString("时间段", string(body))
	if !(exist) {
		fmt.Fprintf(w, "无可订菜单")
		return
	}

	Waitgroup.Add(1)
	go func() {
		re_menu := regexp.MustCompile("<td width=\"40\">(.*)</td><td width=\"50\">(.*)</td><td align=\"left\" nowrap=\"nowrap\">(.*)</td><td width=\"50\">(&nbsp;)?(<font color='red'>套餐</font>)?</td><td width=\"50\">&nbsp;</td><td width=\"50\">(.*)</td><td>")
		menu = re_menu.FindAllStringSubmatch(string(body), 30)
		Waitgroup.Done()
	}()

	Waitgroup.Add(1)
	go func() {
		re_num := regexp.MustCompile("<span id=\"Repeater1_GvReport_(.*)_TxtNum_[0-7]+\">(.*)</span>")
		num = re_num.FindAllStringSubmatch(string(body), 30)
		Waitgroup.Done()
	}()
	
	Waitgroup.Add(1)
	go func() {
		peroids := []string{"早餐", "午餐", "晚餐"}
		re_notorder := regexp.MustCompile("checked=\"checked\"")
		notorder_raw = re_notorder.FindAllStringIndex(string(body), 3)
		for _, n := range notorder_raw {
			number0, _ := strconv.Atoi(string(body)[(n[0] - 16):(n[0] - 15)])
			notorder = append(notorder, peroids[number0])
		}
		Waitgroup.Done()
	}()

	Waitgroup.Wait()

	menu_len := len(menu)

	for n := 0; n < menu_len; n++ {
		index    = menu[n][1]
		category = menu[n][2]
		name     = menu[n][3]
		price    = menu[n][6]
		number   = num[n][2]
	}

	return
}