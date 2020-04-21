// @Title  请填写文件名称
// @Description  请填写文件描述
package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println(common.Pdkf("123456"))
	//转换时间
	loc, _ := time.LoadLocation("Local")
	the_time, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-04-09 08:55:02", loc)
	if err == nil {
		fmt.Println(the_time.Format("2006-01-02"))
	}
}
