// @Title  请填写文件名称
// @Description  请填写文件描述
package common

import (
	"../functions/sql"
	"crypto/sha256"
	"encoding/base64"
	"golang.org/x/crypto/pbkdf2"
	"math/rand"
	"time"
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

//随机生成多位随机数
func GetRandomNum(n int) string {
	var randomStr = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	data := ""
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		data += randomStr[rand.Intn(len(randomStr))]
	}
	return data
}
