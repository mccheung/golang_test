package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)


func main() {
	client := &http.Client{}

	request, _ := http.NewRequest("GET", "http://www.baidu.com", nil )

	request.Header.Set("Accept", "text/html, application/xhtml+xml, application/xml;q=0.9,*/*;q=0.8")
	request.Header.Set("Accept-Charset", "GBK, utf-8;q=0.7,*;q=0.3")
	request.Header.Set("Accept-Encoding","gzip,deflate,sdch")
	request.Header.Set("Accept-Language","zh-CN,zh;q=0.8")
	request.Header.Set("Cache-Control","max-age=0")
	request.Header.Set("Connection","keep-alive")

	response,_ := client.Do(request)

	if response.StatusCode == 200 {
		body,_ := ioutil.ReadAll(response.Body)

		bodystr := string(body)
		fmt.Println(bodystr)
	}
}

