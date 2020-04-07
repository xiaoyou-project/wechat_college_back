package user

import (
	"../sql"
	"github.com/labstack/echo"
	"net/http"
)

/**
用户注册函数
*/
func UserRegistered(c echo.Context) error {
	//获取参数
	avatar := c.FormValue("imgUrl")
	openid := c.FormValue("openid")
	nickname := c.FormValue("nickname")
	//参数不能为空
	if avatar == "" || openid == "" || nickname == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":"参数错误"}`))
	}
	//判断用户是否注册
	if data, err := sql.Sql_dql("select * from user_info where openid='" + openid + "'"); data[0][0] != "" || err != nil {
		if err != nil {
			return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":"数据库查询失败"}`))
		}
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":"该用户已注册"}`))
	}
	//用户没有注册，自动插入数据
	if sql.Sql_dml("insert into user_info (`imgUrl`,`nickName`,`registeredTime`,`sex`,`name`,`college`,`openid`) values ('" + avatar + "','" + nickname + "',now(),'保密','无名侠','保密','" + openid + "')") {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":"注册成功"}`))
	}
	return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":"注册失败"}`))
}
