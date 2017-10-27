package slack

import (
	"peaberry/config"
	"fmt"
	"net/http"
	"bytes"
	"peaberry/adaptor/ghttp"
)

var HOST = "https://hooks.slack.com"
var PATH = "/services/T7CHGPPBN/"
var conf = config.GetInstance().Slack
var URL = HOST + PATH + conf.Token

func Notify(title string, msg string) {
	var jsonStr = []byte(
		fmt.Sprintf(
			`{"channel": "%s", "username": "%s", "text": "%s\n%s", "icon_emoji": "%s"}`,
			conf.Channel,
			conf.UserName,
			title,
			msg,
			conf.Icon))
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
	}

	ghttp.Do(req)
}
