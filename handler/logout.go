package handler

import (
	"fmt"
	"net/http"
	. "util"
)

func LogoutHandler(w http.ResponseWriter, request *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	var client http.Client

	// Take values in the form from request

	request.ParseForm()
	LcSoftCardV2 := request.FormValue("LcSoftCardV2")
	SessionId    := request.FormValue("SessionId")

	// Make Cookies

	cookie_CasModule    := http.Cookie {
		Name: "LcSoftCardV2.CasModule",
		Value: "/card/",
	}
	cookie_LcSoftCardV2 := http.Cookie {
		Name: ".LcSoftCardV2",
		Value: LcSoftCardV2,
	}
	cookie_SessionId    := http.Cookie {
		Name: "ASP.NET_SessionId",
		Value: SessionId,
	}

	// GET DATE Page

	request, err := http.NewRequest("GET", LOGOUT, nil)
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
	
	fmt.Fprintln(w, "登出成功")
	return
}