package user

import (
	"../sql"
	"../tools"
	"encoding/json"
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
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//判断用户是否注册
	if data, err := sql.Sql_dql("select * from user_info where openid='" + openid + "'"); data[0][0] != "" || err != nil {
		if err != nil {
			return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"数据库查询失败"}`))
		}
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"该用户已注册"}`))
	}
	//用户没有注册，自动插入数据
	if sql.Sql_dml("insert into user_info (`imgUrl`,`nickName`,`registeredTime`,`sex`,`name`,`college`,`openid`) values ('" + avatar + "','" + nickname + "',now(),'保密','无名侠','保密','" + openid + "')") {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":{},"msg":"注册成功"}`))
	}
	return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"注册失败"}`))
}

/**
获取openid
*/
func GetOpenId(c echo.Context) error {
	code := c.FormValue("code")
	if code == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//读取小程序配置
	data := sql.GetConfig("miniProgram")
	//发送请求
	result := tools.Get("https://api.weixin.qq.com/sns/jscode2session?appid=" + data["appId"] + "&secret=" + data["secret"] + "&js_code=" + code + "&grant_type=authorization_code")
	//解析获取到的json数据
	var v interface{}
	json.Unmarshal([]byte(result), &v)
	openid := v.(map[string]interface{})["openid"]
	if openid != nil {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":{"openid":"`+openid.(string)+`"},"msg":"获取openid成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"无法获取openid"}`))
	}
}

/**
获取用户信息
*/
func GetUserInfo(c echo.Context) error {
	//获取openid或者userID
	openid := c.FormValue("openid")
	userId := c.FormValue("userId")
	//两个都为空报错
	if openid == "" && userId == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//userId为空的时候置为0
	if userId == "" {
		userId = "0"
	}
	//获取数据
	result, err := sql.Sql_dql("select * from user_info where ID=" + userId + " or openid='" + openid + "'")
	if err != nil {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"读取数据库失败"}`))
	}
	//赋值
	data := make(map[string]string)
	data["name"] = result[0][5]
	data["sex"] = result[0][4]
	data["college"] = result[0][6]
	data["userID"] = result[0][0]
	str, _ := json.Marshal(data)
	//判断是否有这个用户
	if result[0][0] == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":{},"msg":"没有找到该用户"}`))
	}
	return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":`+string(str)+`,"msg":"获取数据成功"}`))
}

/**
更新用户数据
*/
func UpdateUserInfo(c echo.Context) error {
	//获取参数
	name := c.FormValue("name")
	sex := c.FormValue("sex")
	college := c.FormValue("college")
	openid := c.FormValue("openid")
	//参数不能为空
	if name == "" || sex == "" || college == "" || openid == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//数据库更新数据
	if sql.Sql_dml("update user_info set name='" + name + "',sex='" + sex + "',college='" + college + "' where openid='" + openid + "'") {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":{},"msg":"更新数据成功"}`))
	}
	return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"更新数据失败"}`))
}

/**
进入个人中心获取分享列表
*/
func GetShareList(c echo.Context) error {
	//获取参数
	openid := c.FormValue("openid")
	//参数不能为空
	if openid == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//连表查询获取个人分享的内容
	result, err := sql.Sql_dql("select b.ID,b.title,b.date,b.view,b.content,b.imgUrl,c.name from user_info a,share b,plate c where a.ID=b.userID and b.plateID=c.ID and a.openid='" + openid + "'")
	if err != nil {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"读取数据库失败"}`))
	}
	datas := new([]map[string]string)
	//遍历result
	for _, v := range result {
		data := make(map[string]string)
		data["id"] = v[0]
		data["title"] = v[1]
		data["time"] = v[2]
		data["view"] = v[3]
		data["description"] = tools.GetDescription(v[4])
		data["img"] = tools.GetDefaultImg(v[5])
		data["plate"] = v[6]
		*datas = append(*datas, data)
	}
	//解析为json数据
	str, _ := json.Marshal(datas)
	return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":`+string(str)+`,"msg":"获取数据成功"}`))
}

/**
获取个人中心的打卡列表
*/
func GetCardList(c echo.Context) error {
	//获取参数
	openid := c.FormValue("openid")
	//参数不能为空
	if openid == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"接口开发中"}`))
}
