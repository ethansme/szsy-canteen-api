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