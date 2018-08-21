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

	return
}
