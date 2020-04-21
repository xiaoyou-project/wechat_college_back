package admin

import (
	"../../common"
	"../sql"
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
)

func AdminToken(c echo.Context) error {
	//获取用户名和密码
	user := c.FormValue("user")
	password := c.FormValue("password")
	if user == "" || password == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//判断密码是否正确
	if result, err := sql.Sql_dql("select user,password from admin where user='" + user + "'"); err != nil || result[0][0] == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"没有该用户"}`))
	} else {
		//判断密码是否错误
		//对密码进行加密
		password = common.Pdkf(password)
		//比对数据库密码
		if password == result[0][1] {
			return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":{},"msg":"登录成功"}`))
		}
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"用户名或者密码错误"}`))
	}
}

/*获取用户数，话题数，经验数，打开数*/
func VisualizationOverall(c echo.Context) error {
	if result, err := sql.Sql_dql("select count(ID),(select count(ID) from card),(select count(ID) from share),(select count(ID) from plate where plateType=1) from  user_info"); err != nil || result[0][0] == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"查询数据失败"}`))
	} else {
		data := make(map[string]string)
		data["user"] = result[0][0]
		data["card"] = result[0][1]
		data["share"] = result[0][2]
		data["topical"] = result[0][3]
		str, _ := json.Marshal(data)
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":`+string(str)+`,"msg":"查询数据成功"}`))
	}
}
