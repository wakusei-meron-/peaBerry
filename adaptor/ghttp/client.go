package ghttp

import (
	"net/http"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
)

var client = &http.Client{}

func Do(req *http.Request) []byte {
	fmt.Println("http request", zap.Reflect("request", req))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return body
}
