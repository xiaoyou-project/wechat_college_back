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
		//更新个人的信息
		sql.Sql_dml("update user_info set imgUrl='" + avatar + "',nickName='" + nickname + "' where openid='" + openid + "'")
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":{"userID":`+data[0][0]+`},"msg":"获取用户信息成功"}`))
	}
	//用户没有注册，自动插入数据
	if result, id := sql.Sql_dml_id("insert into user_info (`imgUrl`,`nickName`,`registeredTime`,`sex`,`name`,`college`,`openid`) values ('" + avatar + "','" + nickname + "',now(),'保密','无名侠','保密','" + openid + "')"); result {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":{"userID":`+id+`},"msg":"注册成功"}`))
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
	userId := c.FormValue("userId")
	//参数不能为空
	if userId == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//连表查询获取个人分享的内容
	result, err := sql.Sql_dql("select b.ID,b.title,b.date,b.view,b.content,b.imgUrl,c.name,b.good from share b,plate c where b.plateID=c.ID and b.userID=" + userId + "")
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
		data["good"] = v[7]
		*datas = append(*datas, data)
	}
	//解析为json数据
	str, _ := json.Marshal(datas)
	//判断数据是否为空
	if result[0][0] == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":[],"msg":"获取数据成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":`+string(str)+`,"msg":"获取数据成功"}`))
	}
}

/**
获取个人中心的打卡列表
*/
func GetCardList(c echo.Context) error {
	//获取参数
	userId := c.FormValue("userId")
	//参数不能为空
	if userId == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//获取加入的打卡
	result, err := sql.Sql_dql("select a.ID,a.title,a.content,a.totalDay,b.keepDay from card a,user_card b where b.cardID=a.ID and b.userID=" + userId)
	if err != nil {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"读取数据库失败"}`))
	}
	datas := new([]map[string]string)
	//遍历result
	for _, v := range result {
		data := make(map[string]string)
		data["id"] = v[0]
		data["title"] = v[1]
		data["description"] = tools.GetDescription(v[2])
		data["keepDay"] = v[4]
		data["totalDay"] = v[3]
		*datas = append(*datas, data)
	}
	//解析为json数据
	str, _ := json.Marshal(datas)
	//判断数据是否为空
	if result[0][0] == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":[],"msg":"获取数据成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":`+string(str)+`,"msg":"获取数据成功"}`))
	}
}

/**
获取个人中心的话题列表
*/
func GetTopicalList(c echo.Context) error {
	//获取参数
	userId := c.FormValue("userId")
	//参数不能为空
	if userId == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//获取加入的打卡
	result, err := sql.Sql_dql("select `ID`,`name`,`content`,`view`,`date`,`good` from plate where plateType=1 and userID=" + userId + "")
	if err != nil {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"读取数据库失败"}`))
	}
	datas := new([]map[string]string)
	//遍历result
	for _, v := range result {
		data := make(map[string]string)
		data["id"] = v[0]
		data["name"] = v[1]
		data["description"] = tools.GetDescription(v[2])
		data["view"] = v[3]
		data["time"] = v[4]
		data["good"] = v[5]
		*datas = append(*datas, data)
	}
	//解析为json数据
	str, _ := json.Marshal(datas)
	//判断数据是否为空
	if result[0][0] == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":[],"msg":"获取数据成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":`+string(str)+`,"msg":"获取数据成功"}`))
	}
}

/**
获取个人收藏列表
*/
func GetCollectTopicalList(c echo.Context) error {
	//获取参数
	userId := c.FormValue("userId")
	//参数不能为空
	if userId == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//获取加入的打卡
	result, err := sql.Sql_dql("select a.ID,a.name,a.content,a.view,a.date from plate a,user_collect b where a.ID=b.shareID and b.collectType=3 and b.userID=" + userId + "")
	if err != nil {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"读取数据库失败"}`))
	}
	datas := new([]map[string]string)
	//遍历result
	for _, v := range result {
		data := make(map[string]string)
		data["id"] = v[0]
		data["name"] = v[1]
		data["description"] = tools.GetDescription(v[2])
		data["view"] = v[3]
		data["time"] = v[4]
		*datas = append(*datas, data)
	}
	//解析为json数据
	str, _ := json.Marshal(datas)
	//判断数据是否为空
	if result[0][0] == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":[],"msg":"获取数据成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":`+string(str)+`,"msg":"获取数据成功"}`))
	}
}

/**
获取个人收藏的经验
*/
func GetCollectShareList(c echo.Context) error {
	//获取参数
	userId := c.FormValue("userId")
	//参数不能为空
	if userId == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//连表查询获取个人分享的内容
	result, err := sql.Sql_dql("select b.ID,b.title,b.date,b.view,b.content,b.imgUrl,c.name from user_collect a,share b,plate c where b.plateID=c.ID and a.shareID=b.ID and a.collectType=2 and a.userID=" + userId + "")
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
	//判断数据是否为空
	if result[0][0] == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":[],"msg":"获取数据成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":`+string(str)+`,"msg":"获取数据成功"}`))
	}
}

/**
获取收到的评论信息
*/
func GetCommentMessage(c echo.Context) error {
	//获取参数
	userId := c.FormValue("userId")
	//参数不能为空
	if userId == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//连表查询获取个人分享的内容
	result, err := sql.Sql_dql("select a.ID,b.nickName,a.status,a.date,c.commentType,c.shareID,c.content from message a,user_info b,comment c where a.messageType=2 and a.userID=b.ID and a.postID=c.ID and a.userID=" + userId + "")
	if err != nil {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"读取数据库失败"}`))
	}
	datas := new([]map[string]string)
	//遍历result
	for _, v := range result {
		data := make(map[string]string)
		//判断类型获取标题和id
		if v[4] == "2" {
			data["type"] = "share"
			result, err := sql.Sql_dql("select title from share where ID=" + v[5])
			if err != nil {
				return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"读取数据库失败"}`))
			}
			data["postTitle"] = result[0][0]
		} else if v[4] == "1" {
			data["type"] = "topical"
			result, err := sql.Sql_dql("select name from plate where ID=" + v[5])
			if err != nil {
				return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"读取数据库失败"}`))
			}
			data["postTitle"] = result[0][0]
		}
		data["id"] = v[0]
		data["name"] = v[1]
		data["status"] = v[2]
		data["time"] = v[3]
		data["postID"] = v[5]
		data["content"] = v[6]
		*datas = append(*datas, data)
	}
	//解析为json数据
	str, _ := json.Marshal(datas)
	//判断数据是否为空
	if result[0][0] == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":[],"msg":"获取数据成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":`+string(str)+`,"msg":"获取数据成功"}`))
	}
}

/**
获取收到的赞的消息
*/
func GetGoodMessage(c echo.Context) error {
	//获取参数
	userId := c.FormValue("userId")
	//参数不能为空
	if userId == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//连表查询获取个人分享的内容
	result, err := sql.Sql_dql("select a.ID,b.name,a.status,a.date,a.postType,a.postID from message a,user_info b where a.messageType=1 and a.userID=b.ID and a.userID=" + userId + "")
	if err != nil {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"读取数据库失败"}`))
	}
	datas := new([]map[string]string)
	//遍历result
	for _, v := range result {
		data := make(map[string]string)
		//判断类型获取标题和id
		if v[4] == "2" {
			data["type"] = "share"
			result, err := sql.Sql_dql("select title from share where ID=" + v[5])
			if err != nil {
				return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"读取数据库失败"}`))
			}
			data["postTitle"] = result[0][0]
		} else if v[4] == "1" {
			data["type"] = "topical"
			result, err := sql.Sql_dql("select name from plate where ID=" + v[5])
			if err != nil {
				return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"读取数据库失败"}`))
			}
			data["postTitle"] = result[0][0]
		} else if v[4] == "3" {
			data["type"] = "comment"
			result, err := sql.Sql_dql("select content from comment where ID=" + v[5])
			if err != nil {
				return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"读取数据库失败"}`))
			}
			data["postTitle"] = tools.GetDescription(result[0][0])
		}
		data["id"] = v[0]
		data["name"] = v[1]
		data["status"] = v[2]
		data["time"] = v[3]
		data["postID"] = v[5]
		*datas = append(*datas, data)
	}
	//解析为json数据
	str, _ := json.Marshal(datas)
	//判断数据是否为空
	if result[0][0] == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":[],"msg":"获取数据成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":`+string(str)+`,"msg":"获取数据成功"}`))
	}
}

/**
获取系统消息
*/
func GetSystemMessage(c echo.Context) error {
	//获取参数
	userId := c.FormValue("userId")
	//参数不能为空
	if userId == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//连表查询获取个人分享的内容
	result, err := sql.Sql_dql("select ID,detail,date,status from message where messageType=3 and userID=" + userId + "")
	if err != nil {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"读取数据库失败"}`))
	}
	datas := new([]map[string]string)
	//遍历result
	for _, v := range result {
		data := make(map[string]string)
		//判断类型获取标题和id
		data["id"] = v[0]
		data["content"] = v[1]
		data["time"] = v[2]
		data["status"] = v[3]
		*datas = append(*datas, data)
	}
	//解析为json数据
	str, _ := json.Marshal(datas)
	//判断数据是否为空
	if result[0][0] == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":[],"msg":"获取数据成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":`+string(str)+`,"msg":"获取数据成功"}`))
	}
}

/**
修改消息的状态
*/
func SetMessageStatus(c echo.Context) error {
	id := c.FormValue("messageId")
	status := c.FormValue("status")
	//判断参数是否为空
	if id == "" || status == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//判断是已读还是删除
	result := false
	if status == "2" {
		result = sql.Sql_dml("update message set messageType=4 where ID=" + id)
	} else if status == "1" {
		result = sql.Sql_dml("update message set status=1 where ID=" + id)
	}
	//返回相应结果
	if result {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":{},"msg":"更新状态成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"更新状态失败"}`))
	}
}
