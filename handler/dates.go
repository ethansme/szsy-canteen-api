package handler

import (
	"fmt"
	"time"
	"regexp"
	"strconv"
	"strings"
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
	. "util"
)

type Dates struct {
	Dates []string `json:"可订日期"`
}

func DatesHandler(w http.ResponseWriter, request *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	var client http.Client

	// Take values in the form from request

	request.ParseForm()
	cookies      := request.FormValue("口令")
	LcSoftCardV2 := strings.Split(cookies, ",")[0]
	SessionId    := strings.Split(cookies, ",")[1]

	// Make Cookies

	cookie_CasModule    := http.Cookie{
		Name:  "LcSoftCardV2.CasModule",
		Value: "/card/",
	}
	cookie_LcSoftCardV2 := http.Cookie{
		Name:  ".LcSoftCardV2",
		Value: LcSoftCardV2,
	}
	cookie_SessionId    := http.Cookie{
		Name:  "ASP.NET_SessionId",
		Value: SessionId,
	}

	// GET DATE Page

	request, err := http.NewRequest("GET", DATE, nil)
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

	// Match for available dates

	var (
		VIEWSTATE          string
		VIEWSTATEGENERATOR string
		EVENTVALIDATION    string
		dates_raw          [][]string
	)

	Waitgroup.Add(1)
	go func() {
		re_dates := regexp.MustCompile("Date=(.*)\" target=\"RestaurantContent\">(.*)<font color='red'>订餐</font>")
		dates_raw = re_dates.FindAllStringSubmatch(string(body), 32)
		Waitgroup.Done()
	}()

	// Never forget the next month

	month := int(time.Now().Month())
	year  := int(time.Now().Year())
	if month == 12 {
		year = year + 1
		month = 1
	} else {
		month = month + 1
	}

	data := make(url.Values)
	data["__EVENTTARGET"]              = []string{"DrplstMonth1$DrplstControl"}
	data["__EVENTARGUMENT"]            = []string{""}
	data["__LASTFOCUS"]                = []string{""}
	data["DrplstYear1$DrplstControl"]  = []string{strconv.Itoa(year)}
	data["DrplstMonth1$DrplstControl"] = []string{strconv.Itoa(month)}

	Waitgroup.Wait()

	data["__VIEWSTATE"]          = []string{VIEWSTATE}
	data["__VIEWSTATEGENERATOR"] = []string{VIEWSTATEGENERATOR}
	data["__EVENTVALIDATION"]    = []string{EVENTVALIDATION}

	return
}
