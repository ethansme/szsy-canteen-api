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

	return
}