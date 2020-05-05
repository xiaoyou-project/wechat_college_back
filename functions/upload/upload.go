package upload

import (
	"../../common"
	"../sql"
	"github.com/labstack/echo"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func UploadFile(c echo.Context) error {
	//接收上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"没有文件"}`))
	}

	//打开用户上传的文件
	src, err := file.Open()
	if err != nil {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"打开文件失败"}`))
	}

	//获取文件后缀
	suffix := common.FindMatch("(\\.[a-z]+)$", file.Filename)
	//如果没有就自己设置一个
	if suffix == "" {
		suffix = ".jpg"
	}
	//创建文件
	year := "static/images/" + strconv.Itoa(time.Now().Year())
	month := strconv.Itoa(int(time.Now().Month()))
	filename := year + "/" + month + "/" + time.Now().Format("150405") + common.GetRandomNum(5) + suffix
	err = common.CreatePath(year + "/" + month)
	if err == nil {
		dst, err := os.Create(filename)
		defer dst.Close()
		if err == nil {
			if _, err := io.Copy(dst, src); err == nil {
				//读取博客配置
				option := sql.GetConfig("site")
				//获取图片地址
				src := option["server"] + "/" + filename
				return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":{"src":"`+src+`"},"msg":"上传文件成功"}`))
			}
		}
	}
	return c.JSONBlob(http.StatusOK, []byte(`{"type":"error","msg":"上传文件失败"}`))
}
