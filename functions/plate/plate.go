package plate

import (
	"../sql"
	"../tools"
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
)

/**
获取板块列表
*/
func GetPlateList(c echo.Context) error {
	//连表查询获取个人分享的内容
	result, err := sql.Sql_dql("select a.ID,a.name,a.imgUrl,(select count(ID) from share where a.ID=plateID) from plate a where a.plateType=0 and a.status=1")
	if err != nil {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"读取数据库失败"}`))
	}
	datas := new([]map[string]string)
	//遍历result
	for _, v := range result {
		data := make(map[string]string)
		//判断类型获取标题和id
		data["id"] = v[0]
		data["name"] = v[1]
		data["imgUrl"] = v[2]
		data["total"] = v[3]
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
获取经验列表
*/
func GetShareList(c echo.Context) error {
	//获取参数
	id := c.FormValue("plateID")
	if id == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//连表查询获取个人分享的内容
	result, err := sql.Sql_dql("select a.ID,a.date,a.title,u.imgUrl,u.name,a.content,a.imgUrl,a.view,a.good,(select count(ID) from comment where commentType=1 and shareID=a.ID) from share a,user_info u,plate b where a.userID=u.ID and a.plateID=b.ID and b.ID=" + id)
	if err != nil {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"读取数据库失败"}`))
	}
	datas := new([]map[string]string)
	//遍历result
	for _, v := range result {
		data := make(map[string]string)
		//判断类型获取标题和id
		data["id"] = v[0]
		data["time"] = v[1]
		data["title"] = v[2]
		data["avatar"] = v[3]
		data["name"] = v[4]
		data["content"] = tools.GetDescription(v[5])
		data["img"] = tools.GetDefaultImg(v[6])
		data["total"] = v[8]
		data["view"] = v[7]
		data["comments"] = v[9]
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
获取经验内容
*/
func GetShareContent(c echo.Context) error {
	//获取参数
	userId := c.FormValue("userID")
	shareId := c.FormValue("shareID")
	//判断参数是否为空
	if userId == "" || shareId == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//连表查询获取个人分享的内容
	result, err := sql.Sql_dql("select a.content,a.title,a.date,b.imgUrl,b.name,a.imgUrl,a.view,a.good from share a,user_info b where a.userID=b.ID and a.ID=" + shareId)
	if err != nil {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"读取数据库失败"}`))
	}
	//阅读数+1
	sql.Sql_dml("update share set view=view+1 where ID=" + shareId)
	data := make(map[string]string)
	//判断用户是否点赞
	if goods, _ := sql.Sql_dql("select ID from good where postType=1 and userID=" + userId + " and postID=" + shareId); goods != nil && goods[0][0] != "" {
		data["goodStatus"] = "1"
	} else {
		data["goodStatus"] = "0"
	}
	//判断用户是否收藏
	if collects, _ := sql.Sql_dql("select ID from user_collect where collectType=2 and userID=" + userId + " and shareID=" + shareId); collects != nil && collects[0][0] != "" {
		data["collectStatus"] = "1"
	} else {
		data["collectStatus"] = "0"
	}
	//给其他内容赋值
	data["content"] = result[0][0]
	data["title"] = result[0][1]
	data["time"] = result[0][2]
	data["avatar"] = result[0][3]
	data["name"] = result[0][4]
	data["img"] = result[0][5]
	data["view"] = result[0][6]
	data["good"] = result[0][7]
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
修改经验
*/
func EditShareContent(c echo.Context) error {
	//获取参数
	userId := c.FormValue("userID")
	shareId := c.FormValue("shareID")
	content := c.FormValue("content")
	imgUrl := c.FormValue("imgUrl")
	title := c.FormValue("title")
	//判断参数是否为空
	if userId == "" || shareId == "" || content == "" || imgUrl == "" || title == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//更新数据
	if sql.Sql_dml("update share set content='" + content + "',title='" + title + "',imgUrl='" + imgUrl + "',editTime=now() where userID=" + userId + " and ID=" + shareId) {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":{},"msg":"更新文章成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"更新文章失败"}`))
	}
}

/**
点赞或者取消点赞
*/
func UpdateGood(c echo.Context) error {
	//获取参数
	userId := c.FormValue("userID")
	id := c.FormValue("shareID")
	//判断参数是否为空
	if userId == "" || id == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	result := false
	//判断是否点赞
	if goods, _ := sql.Sql_dql("select ID from good where postType=1 and userID=" + userId + " and postID=" + id); goods != nil && goods[0][0] != "" {
		//说明点赞了
		//点赞数-1
		sql.Sql_dml("update share set good=good-1 where ID=" + id)
		//删除点赞数据
		result = sql.Sql_dml("delete from good where postType=1 and userID=" + userId + " and postID=" + id)
	} else {
		//说明没有点赞
		//点赞数+1
		sql.Sql_dml("update share set good=good+1 where ID=" + id)
		//添加点赞数据
		//添加点赞数据
		result = sql.Sql_dml("insert into good (`userID`,`postType`,`postID`) values (" + userId + ",1," + id + ")")
	}
	//判断操作是否成功
	if result {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":{},"msg":"操作成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"操作失败"}`))
	}
}

/**
收藏经验或者取消收藏经验
*/
func UpdateCollect(c echo.Context) error {
	//获取参数
	id := c.FormValue("shareID")
	userId := c.FormValue("userID")
	//判断参数
	if userId == "" || id == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	result := false
	//判断是否收藏
	if goods, _ := sql.Sql_dql("select ID from user_collect where collectType=2 and userID=" + userId + " and shareID=" + id); goods != nil && goods[0][0] != "" {
		//收藏了，删除收藏
		result = sql.Sql_dml("delete from user_collect where collectType=2 and userID=" + userId + " and shareID=" + id)
	} else {
		//没有收藏，添加收藏
		result = sql.Sql_dml("insert into user_collect (`userID`,`collectType`,`shareID`) values (" + userId + ",2," + id + ")")
	}
	//判断操作是否成功
	if result {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":{},"msg":"操作成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"操作失败"}`))
	}
}

/**
删除经验
*/
func DeleteShare(c echo.Context) error {
	//获取参数
	id := c.FormValue("shareID")
	userId := c.FormValue("userID")
	//判断参数
	if userId == "" || id == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//判断结构
	if sql.Sql_dml("delete from share where ID=" + id + " and userID=" + userId) {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":{},"msg":"操作成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"操作失败"}`))
	}
}

/**
获取我收藏的板块
*/
func GetCollectPlate(c echo.Context) error {
	//获取参数
	userId := c.FormValue("userID")
	//判断参数
	if userId == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//连表查询获取个人分享的内容
	result, err := sql.Sql_dql("select a.ID,a.name,a.imgUrl,(select count(ID) from share where plateID=a.ID) from plate a,user_collect b where b.collectType=1 and a.ID=b.shareID and b.userID=" + userId)
	if err != nil {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"读取数据库失败"}`))
	}
	datas := new([]map[string]string)
	//遍历result
	for _, v := range result {
		data := make(map[string]string)
		//判断类型获取标题和id
		data["id"] = v[0]
		data["name"] = v[1]
		data["imgUrl"] = v[2]
		data["total"] = v[3]
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
收藏板块
*/
func CollectPlate(c echo.Context) error {
	//获取参数
	id := c.FormValue("plateID")
	userId := c.FormValue("userID")
	//判断参数
	if userId == "" || id == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	result := false
	//判断是否收藏
	if goods, _ := sql.Sql_dql("select ID from user_collect where collectType=1 and userID=" + userId + " and shareID=" + id); goods != nil && goods[0][0] != "" {
		//收藏了，删除收藏
		result = sql.Sql_dml("delete from user_collect where collectType=1 and userID=" + userId + " and shareID=" + id)
	} else {
		//没有收藏，添加收藏
		result = sql.Sql_dml("insert into user_collect (`userID`,`collectType`,`shareID`) values (" + userId + ",1," + id + ")")
	}
	//判断操作是否成功
	if result {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":{},"msg":"操作成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"操作失败"}`))
	}
}

/**
申请板块
*/
func ApplicationPlate(c echo.Context) error {
	//获取参数
	content := c.FormValue("content")
	name := c.FormValue("name")
	userId := c.FormValue("userID")
	//判断参数
	if userId == "" || content == "" || name == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//判断结构
	if sql.Sql_dml("insert into plate (`name`,`imgUrl`,`content`,`userID`,`status`,`view`,`date`,`good`,`plateType`) values ('" + name + "','','" + content + "'," + userId + ",0,0,now(),0,0)") {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":{},"msg":"申请板块成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"申请板块失败"}`))
	}
}
