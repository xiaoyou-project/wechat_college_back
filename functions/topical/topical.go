package topical

import (
	"../../common"
	"../sql"
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
)

/**
获取话题列表
*/
func GetTopicalList(c echo.Context) error {
	//连表查询获取个人分享的内容
	result, err := sql.Sql_dql("select ID,name,view from plate where plateType=1 order by view desc")
	if err != nil {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"读取数据库失败"}`))
	}
	datas := new([]map[string]string)
	//遍历result
	for _, v := range result {
		data := make(map[string]string)
		data["id"] = v[0]
		data["title"] = v[1]
		data["view"] = v[2]
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
发布话题
*/
func ReleaseTopical(c echo.Context) error {
	//获取参数
	userId := c.FormValue("userId")
	title := c.FormValue("title")
	content := c.FormValue("content")
	//判断参数是否为空
	if userId == "" || title == "" || content == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//插入数据
	if sql.Sql_dml("insert into plate (`name`,`imgUrl`,`content`,`userID`,`status`,`view`,`date`,`good`,`plateType`) values ('" + title + "','','" + content + "'," + userId + ",1,0,now(),0,1)") {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":[],"msg":"发布话题成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"发布话题失败"}`))
	}
}

/**
获取话题的内容
*/
func GetTopicalContent(c echo.Context) error {
	//获取参数
	userId := c.FormValue("userId")
	topicalId := c.FormValue("topicalID")
	//判断参数是否为空
	if userId == "" || topicalId == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//连表查询获取个人分享的内容
	result, err := sql.Sql_dql("select a.name,a.content,a.view,a.good,b.name,b.imgUrl,a.date,a.userID from plate a,user_info b where a.userID=b.ID and a.ID=" + topicalId)
	if err != nil {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"读取数据库失败"}`))
	}
	//阅读数+1
	sql.Sql_dml("update plate set view=view+1 where ID=" + topicalId)
	data := make(map[string]string)
	//判断用户是否点赞
	if goods, _ := sql.Sql_dql("select ID from good where postType=2 and userID=" + userId + " and postID=" + topicalId); goods != nil && goods[0][0] != "" {
		data["goodStatus"] = "1"
	} else {
		data["goodStatus"] = "0"
	}
	//判断用户是否收藏
	if collects, _ := sql.Sql_dql("select ID from user_collect where collectType=3 and userID=" + userId + " and shareID=" + topicalId); collects != nil && collects[0][0] != "" {
		data["collectStatus"] = "1"
	} else {
		data["collectStatus"] = "0"
	}
	//给其他内容赋值
	data["title"] = result[0][0]
	data["content"] = result[0][1]
	data["view"] = result[0][2]
	data["good"] = result[0][3]
	data["name"] = result[0][4]
	data["imgUrl"] = result[0][5]
	data["time"] = common.TimeChange(result[0][6])
	data["authorID"] = result[0][7]
	//解析为json数据
	str, _ := json.Marshal(data)
	//判断数据是否为空
	if result[0][0] == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":[],"msg":"获取数据成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":`+string(str)+`,"msg":"获取数据成功"}`))
	}
}

/**
话题点赞或者取消赞
*/
func TopicalGood(c echo.Context) error {
	//获取参数
	id := c.FormValue("topicalID")
	userId := c.FormValue("userId")
	//判断参数
	if userId == "" || id == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	result := false
	//判断是否点赞
	if goods, _ := sql.Sql_dql("select ID from good where postType=2 and userID=" + userId + " and postID=" + id); goods != nil && goods[0][0] != "" {
		//说明点赞了
		//点赞数-1
		sql.Sql_dml("update plate set good=good-1 where ID=" + id)
		//删除点赞数据
		result = sql.Sql_dml("delete from good where postType=2 and userID=" + userId + " and postID=" + id)
	} else {
		//说明没有点赞
		//点赞数+1
		sql.Sql_dml("update plate set good=good+1 where ID=" + id)
		//添加点赞数据
		//添加点赞数据
		result = sql.Sql_dml("insert into good (`userID`,`postType`,`postID`) values (" + userId + ",2," + id + ")")

		//点赞了才发生消息通知
		//查询话题的用户id
		if result, err := sql.Sql_dql("select userID from plate where ID=" + id); err == nil && result[0][0] != "" {
			//点赞成功后自动添加用户消息里面
			sql.Sql_dml("insert into message (`userID`,`messageType`,`postID`,`actionID`,`status`,`date`,`detail`,`postType`) values (" + result[0][0] + ",1," + id + "," + userId + ",0,now(),'',1)")
		}
	}
	//判断操作是否成功
	if result {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":{},"msg":"操作成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"操作失败"}`))
	}
}

/**
话题收藏或者取消收藏
*/
func UpdateCollect(c echo.Context) error {
	//获取参数
	id := c.FormValue("topicalID")
	userId := c.FormValue("userId")
	//判断参数
	if userId == "" || id == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	result := false
	//判断是否收藏
	if goods, _ := sql.Sql_dql("select ID from user_collect where collectType=3 and userID=" + userId + " and shareID=" + id); goods != nil && goods[0][0] != "" {
		//收藏了，删除收藏
		result = sql.Sql_dml("delete from user_collect where collectType=3 and userID=" + userId + " and shareID=" + id)
	} else {
		//没有收藏，添加收藏
		result = sql.Sql_dml("insert into user_collect (`userID`,`collectType`,`shareID`) values (" + userId + ",3," + id + ")")
	}
	//判断操作是否成功
	if result {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":{},"msg":"操作成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"操作失败"}`))
	}
}

/**
删除话题
*/
func DeleteTopical(c echo.Context) error {
	//获取参数
	id := c.FormValue("topicalID")
	userId := c.FormValue("userId")
	//判断参数
	if userId == "" || id == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//删除话题
	if sql.Sql_dml("delete from plate where plateType=1 and userID=" + userId + " and ID=" + id) {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":{},"msg":"删除话题成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"删除话题失败"}`))
	}
}
