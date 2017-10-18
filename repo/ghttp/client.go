package ghttp

import (
	"net/http"
	"fmt"
	"encoding/json"
	"peaberry/repo/dto"
	"peaberry/util"
)

var HOST = "https://imt.services.isca.jp"
var COOKIE = "schSchDialogReferWidth=725; schSchDialogReferHeight=600; jp.co.intra_mart.session.cookie=11rrowk5|1508200819833; __zlcmid=iognBu9o4TNeii; _ga=GA1.2.2011196724.1507019323; _gid=GA1.2.48615766.1508113015; AWSELB=0DFB4119120A7C8E000FD729A065620BDD13FCDC87B5EA90D77EBEB2DDCF22CB4F694EA1144A3A4103310D5C5D19B64D16892963A262F00439F7E1466908CC95E55DD753D6; opentoken=T1RLAQIXLBLg2fG3gCeloasti8yJHPKOUBBCmwwsyOMWegSN5oJ04xjCAACA7dix61XmMo1x8bgs6eYk7pjBm_JY24xbRhB7k0jdTDzg2LcU9vvrMQDgerwXVRlo2vXVq-FdLKe2w3Sz5GfTe6g-23T04GOInSyf8snBEkzxlvMmetBtIeLbVADGI4bD9pjLWhlC94FOM9V65v-AKqG-FbTAZZs-rBvzJdbydWs*; JSESSIONID=baakKk42mGwL2eFYoiv8v; _gat=1"

func Post(path string) *dto.Response{
	client := &http.Client{}
	req, err := http.NewRequest("POST", HOST + path, nil)
	req.Header.Set("Cookie", COOKIE)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer resp.Body.Close()

	data := &dto.Response{}
	err2 := json.NewDecoder(resp.Body).Decode(&data)
	if err2 != nil {
		fmt.Println(err2)
		panic(err)
	}
	util.PrettyPrint(data)
	return data
}
