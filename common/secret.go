// @Title  请填写文件名称
// @Description  请填写文件描述
package common

import (
	"../functions/sql"
	"crypto/sha256"
	"encoding/base64"
	"golang.org/x/crypto/pbkdf2"
)

//密码加密算法
func Pdkf(password string) string {
	data := sql.GetConfig("password")
	if data == nil {
		return ""
	}
	dk := pbkdf2.Key([]byte(password), []byte(data["salt"]), 4096, 32, sha256.New)
	str := base64.StdEncoding.EncodeToString(dk)
	return str
}
