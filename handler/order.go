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

	return
}
