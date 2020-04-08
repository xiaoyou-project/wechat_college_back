package tools

import (
	"io/ioutil"
	"net/http"
	"strings"
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

/*获取简略信息*/
func GetDescription(content string) string {
	num := 30
	//先去除前后空格和换行
	content = strings.Trim(content, " \n")
	//判断内容是不是信息框
	if len([]rune(content)) < num {
		return content
	} else {
		return string([]rune(content)[:num]) + "...."
	}
}

/*获取代表性的图片*/
func GetDefaultImg(img string) string {
	imgs := strings.Split(img, "&&")
	return imgs[0]
}
