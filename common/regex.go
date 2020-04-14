package common

import "regexp"

/*寻找匹配的正则*/
func FindMatch(rege string, content string) string {
	r, err := regexp.Compile(rege)
	if err != nil {
		return ""
	}
	if data := r.FindStringSubmatch(content); data != nil {
		return data[1]
	}
	return ""
}
