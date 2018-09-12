package util

import (
	"sync"
	"regexp"
)

const (
	CAS     = "http://gzb.szsy.cn:3000/cas"
	LOGIN   = "http://passport-yun.szsy.cn/login?service=http://gzb.szsy.cn/card/Default.aspx"
	CARD    = "http://gzb.szsy.cn/card"
	SERVICE = "?service=" + CAS + "/Default.aspx"
	DATE    = CARD + "/Restaurant/RestaurantUserMenu/RestaurantUserSelect.aspx"
	MENU    = CARD + "/Restaurant/RestaurantUserMenu/RestaurantUserMenu.aspx?Date="
	LOGOUT  = "http://gzb.szsy.cn:4000/lcconsole/"
)

var Waitgroup sync.WaitGroup

func Find(data *string, match string, body string) () {
	re   := regexp.MustCompile(match)
	*data = re.FindStringSubmatch(string(body))[1]
	Waitgroup.Done()
}