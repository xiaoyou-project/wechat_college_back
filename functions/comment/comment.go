package comment

import (
	"../sql"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

/**
获取评论内容
*/
func GetCommentList(c echo.Context) error {
	//获取评论的内容
	//获取参数
	id := c.FormValue("postID")
	postType := c.FormValue("postType")
	userId := c.FormValue("userId")
	//判断参数
	if userId == "" || id == "" || postType == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//连表查询获取个人分享的内容
	result, err := sql.Sql_dql("select a.ID,b.name,a.good,a.content,b.imgUrl,b.ID,a.date from comment a,user_info b where a.commentType=" + postType + " and a.userID=b.ID and a.shareID=" + id)
	if err != nil {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"读取数据库失败"}`))
	}
	datas := new([]map[string]string)
	//遍历result
	for _, v := range result {
		data := make(map[string]string)
		data["id"] = v[0]
		data["name"] = v[1]
		data["good"] = v[2]
		data["content"] = v[3]
		data["imgUrl"] = v[4]
		data["userID"] = v[5]
		data["time"] = v[6]
		//判断是否点赞该评论
		if good, _ := sql.Sql_dql("select ID from good where postType=3 and userID=" + userId + " and postID=" + v[0]); good != nil && good[0][0] != "" {
			data["state"] = "1"
		} else {
			data["state"] = "0"
		}
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
删除评论
*/
func DeleteComment(c echo.Context) error {
	commentId := c.FormValue("commentID")
	userId := c.FormValue("userID")
	if commentId == "" || userId == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//删除评论
	if sql.Sql_dml("delete from comment where userID=" + userId + " and ID=" + commentId) {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":[],"msg":"删除评论成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"删除评论失败"}`))
	}
}

/**
发表评论
*/
func PublishComment(c echo.Context) error {
	//获取参数
	commentType := c.FormValue("commentType")
	userId := c.FormValue("userID")
	content := c.FormValue("content")
	postID := c.FormValue("postID")
	fmt.Println(content)
	//判断参数是否为空
	if commentType == "" || userId == "" || content == "" || postID == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//插入数据
	if errs, id := sql.Sql_dml_id("insert into comment (`shareID`,`commentType`,`userID`,`content`,`date`,`good`) values (" + postID + "," + commentType + "," + userId + ",'" + content + "',now(),0)"); errs {
		//根据不同的文章查找用户id
		var result [][]string
		var err error
		posType := ""
		if commentType == "1" {
			posType = "2"
			//经验分享
			result, err = sql.Sql_dql("select userID from share where ID=" + postID)
		} else if commentType == "2" {
			posType = "1"
			//话题
			result, err = sql.Sql_dql("select userID from plate where ID=" + postID)
		} else if commentType == "3" {
			posType = "4"
			//打卡
			result, err = sql.Sql_dql("select userID from card where ID=" + postID)
		}
		//插入数据
		if err == nil && result != nil && result[0][0] != "" {
			sql.Sql_dml("insert into message (`userID`,`messageType`,`postID`,`actionID`,`status`,`date`,`detail`,`postType`) values (" + result[0][0] + ",2," + id + "," + userId + ",0,now(),''," + posType + ")")
		}

		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":[],"msg":"发表评论成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"发表评论失败"}`))
	}
}

/**
评论点赞或者取消赞
*/
func UpdateGood(c echo.Context) error {
	//获取参数
	userId := c.FormValue("userID")
	id := c.FormValue("commentID")
	//判断参数是否为空
	if userId == "" || id == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	result := false
	//判断是否点赞
	if goods, _ := sql.Sql_dql("select ID from good where postType=3 and userID=" + userId + " and postID=" + id); goods != nil && goods[0][0] != "" {
		//说明点赞了
		//点赞数-1
		sql.Sql_dml("update comment set good=good-1 where ID=" + id)
		//删除点赞数据
		result = sql.Sql_dml("delete from good where postType=3 and userID=" + userId + " and postID=" + id)
	} else {
		//说明没有点赞
		//点赞数+1
		sql.Sql_dml("update comment set good=good+1 where ID=" + id)
		//添加点赞数据
		//添加点赞数据
		result = sql.Sql_dml("insert into good (`userID`,`postType`,`postID`) values (" + userId + ",3," + id + ")")

		//点赞了才发生消息通知
		//查询话题的用户id
		if result, err := sql.Sql_dql("select userID from comment where ID=" + id); err == nil && result[0][0] != "" {
			//点赞成功后自动添加用户消息里面
			sql.Sql_dml("insert into message (`userID`,`messageType`,`postID`,`actionID`,`status`,`date`,`detail`,`postType`) values (" + result[0][0] + ",1," + id + "," + userId + ",0,now(),'',3)")
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
判断评论的类型
*/
func CommentType(c echo.Context) error {
	id := c.FormValue("postID")
	if id == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//返回评论的类型
	data := make(map[string]string)
	result, err := sql.Sql_dql("select commentType,shareID from comment where ID=" + id)
	if err != nil {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"读取数据库失败"}`))
	}
	//判断类型获取标题和id
	data["type"] = result[0][0]
	data["postID"] = result[0][1]
	str, _ := json.Marshal(data)
	//判断数据是否为空
	if result[0][0] == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":[],"msg":"获取数据成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":`+string(str)+`,"msg":"获取数据成功"}`))
	}
}
