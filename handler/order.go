package handler

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"net/url"
	"net/http"
	"io/ioutil"
	. "util"
)

func OrderHandler(w http.ResponseWriter, request *http.Request) {

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
	breakfast    := request.FormValue("早餐")
	lunch        := request.FormValue("午餐")
	dinner       := request.FormValue("晚餐")
	LcSoftCardV2 := strings.Split(cookies, ",")[0]
	SessionId    := strings.Split(cookies, ",")[1]

	// Make Cookies

	cookie_CasModule := http.Cookie{
		Name:  "LcSoftCardV2.CasModule",
		Value: "/card/",
	}
	cookie_LcSoftCardV2 := http.Cookie{
		Name:  ".LcSoftCardV2",
		Value: LcSoftCardV2,
	}
	cookie_SessionId := http.Cookie{
		Name:  "ASP.NET_SessionId",
		Value: SessionId,
	}

	request, err = http.NewRequest("GET", MENU+date, nil)
	if err != nil {
		panic(err)
	}

	// Add Cookies

	request.AddCookie(&cookie_CasModule)
	request.AddCookie(&cookie_LcSoftCardV2)
	request.AddCookie(&cookie_SessionId)

	return
}
