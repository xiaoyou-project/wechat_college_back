package tools

import (
	"io/ioutil"
	"net/http"
)

/*发送get请求*/
func Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(s)
}
