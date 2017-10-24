package slack

import (
	"net/http"
	"fmt"
	"bytes"
	"io/ioutil"
)

var HOST = "https://hooks.slack.com"
var PATH = "/services/T7CHGPPBN/B7PNQ2HN2/AFY1XuS2bTn52Js6atMLWCoM"
var URL = HOST + PATH

func Notify(title string, msg string, channel string, username string, icon_emoji string) {
	var jsonStr = []byte(fmt.Sprintf(`{"channel": "%s", "username": "%s", "text": "%s\n%s", "icon_emoji": "%s"}`, channel, username, title, msg, icon_emoji))
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}