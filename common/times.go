// @Title  和时间处理有关的接口
package common

import "time"

func TimeChange(times string) string {
	loc, _ := time.LoadLocation("Local")
	theTime, err := time.ParseInLocation("2006-01-02 15:04:05", times, loc)
	if err == nil {
		return theTime.Format("2006-01-02")
	}
	return ""
}
