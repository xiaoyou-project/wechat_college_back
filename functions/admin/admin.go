package admin

import (
	"../../common"
	"../sql"
	"../tools"
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
	if user == "admin" || password == "123" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":{},"msg":"登录成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"用户名或者密码错误"}`))
	}
	////判断密码是否正确
	//if result, err := sql.Sql_dql("select user,password from admin where user='" + user + "'"); err != nil || result[0][0] == "" {
	//	return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"没有该用户"}`))
	//} else {
	//	//判断密码是否错误
	//	//对密码进行加密
	//	password = common.Pdkf(password)
	//	//比对数据库密码
	//	if password == result[0][1] {
	//		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":{},"msg":"登录成功"}`))
	//	}
	//	return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"用户名或者密码错误"}`))
	//}
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

//获取所有的板块
func GetAllPlate(c echo.Context) error {
	if result, err := sql.Sql_dql("select ID,imgUrl,name,content,userID,date,status from plate where plateType=0"); err != nil || result[0][0] == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"查询数据失败"}`))
	} else {
		datas := *new([]map[string]string)
		for _, v := range result {
			data := make(map[string]string)
			data["id"] = v[0]
			data["imgUrl"] = v[1]
			data["name"] = v[2]
			data["content"] = v[3]
			data["userId"] = v[4]
			data["date"] = v[5]
			data["status"] = v[6]
			datas = append(datas, data)
		}
		str, _ := json.Marshal(datas)
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":`+string(str)+`,"msg":"查询数据成功"}`))
	}
}

//获取所有用户
func GetAllUser(c echo.Context) error {
	if result, err := sql.Sql_dql("select * from user_info"); err != nil || result[0][0] == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"查询数据失败"}`))
	} else {
		datas := *new([]map[string]string)
		for _, v := range result {
			data := make(map[string]string)
			data["id"] = v[0]
			data["imgUrl"] = v[1]
			data["nickName"] = v[2]
			data["registeredTime"] = v[3]
			data["sex"] = v[4]
			data["name"] = v[5]
			data["college"] = v[6]
			data["openid"] = v[7]
			datas = append(datas, data)
		}
		str, _ := json.Marshal(datas)
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":`+string(str)+`,"msg":"查询数据成功"}`))
	}
}

//删除用户
func DeleteUser(c echo.Context) error {
	//获取相关的参数
	user := c.FormValue("userID")
	if user == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	if sql.Sql_dml("delete from user_info where ID=" + user) {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":"","msg":"删除用户成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":"","msg":"删除用户失败"}`))
	}
}

//删除板块
func DeletePlate(c echo.Context) error {
	//获取相关的参数
	user := c.FormValue("plateID")
	if user == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	if sql.Sql_dml("delete from plate where ID=" + user + " and plateType=0") {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":"","msg":"删除板块成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":"","msg":"删除板块失败"}`))
	}
}

//修改板块信息
func ChangePlateInfo(c echo.Context) error {
	//获取相关的参数
	plate := c.FormValue("plateID")
	name := c.FormValue("name")
	img := c.FormValue("imgUrl")
	content := c.FormValue("content")
	if plate == "" || name == "" || img == "" || content == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//修改信息
	if sql.Sql_dml("update plate set name='" + name + "',imgUrl='" + img + "',content='" + content + "' where ID=" + plate) {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":"","msg":"修改信息成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":"","msg":"修改信息失败"}`))
	}
}

//修改板块状态
func ChangePlateStatus(c echo.Context) error {
	//获取相关的参数
	plate := c.FormValue("plateID")
	status := c.FormValue("status")
	userID := c.FormValue("userID")
	if plate == "" || status == "" || userID == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//修改信息
	if sql.Sql_dml("update plate set status='" + status + "' where ID=" + plate) {
		if status == "1" {
			//获取板块的名字
			if result, err := sql.Sql_dql("select name from plate where ID=" + plate); err == nil && result[0][0] != "" {
				sql.Sql_dml("insert into message (userID,messageType,postID,actionID,status,date,detail,postType) values (" + userID + ",3,0,0,0,now(),'你申请的板块(" + result[0][0] + ")已通过审核',0)")
			}
		}
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":"","msg":"修改信息成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":"","msg":"修改信息失败"}`))
	}
}

//修改话题内容
func ChangeTopicalContent(c echo.Context) error {
	//获取相关的参数
	plate := c.FormValue("topicalID")
	title := c.FormValue("title")
	content := c.FormValue("content")
	if plate == "" || title == "" || content == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//修改信息
	if sql.Sql_dml("update plate set name='" + title + "',content='" + content + "' where plateType=1 and ID=" + plate) {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":"","msg":"修改信息成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":"","msg":"修改信息失败"}`))
	}
}

//修改打卡的信息
func ChangeCardInfo(c echo.Context) error {
	//获取相关的参数
	plate := c.FormValue("cardID")
	title := c.FormValue("title")
	content := c.FormValue("content")
	total := c.FormValue("totalDay")
	if plate == "" || title == "" || content == "" || total == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//修改信息
	if sql.Sql_dml("update card set title='" + title + "',content='" + content + "',totalDay=" + total + " where ID=" + plate) {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":"","msg":"修改信息成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":"","msg":"修改信息失败"}`))
	}
}

//获取所有的评论
func GetAllComment(c echo.Context) error {
	if result, err := sql.Sql_dql("select ID,content,userID,date,good,commentType from comment"); err != nil || result[0][0] == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"查询数据失败"}`))
	} else {
		datas := *new([]map[string]string)
		for _, v := range result {
			data := make(map[string]string)
			data["id"] = v[0]
			data["content"] = v[1]
			data["userID"] = v[2]
			data["time"] = v[3]
			data["good"] = v[4]
			data["type"] = v[5]
			datas = append(datas, data)
		}
		str, _ := json.Marshal(datas)
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":`+string(str)+`,"msg":"查询数据成功"}`))
	}
}

//删除评论
func DeleteComment(c echo.Context) error {
	//获取相关的参数
	user := c.FormValue("commentID")
	if user == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	if sql.Sql_dml("delete from comment where ID=" + user) {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":"","msg":"删除评论成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":"","msg":"删除评论失败"}`))
	}
}

//添加板块
func AddPlate(c echo.Context) error {
	name := c.FormValue("name")
	content := c.FormValue("content")
	imgUrl := c.FormValue("imgUrl")
	if name == "" || content == "" || imgUrl == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	if sql.Sql_dml("insert into plate (name,imgUrl,content,userID,status,view,date,good,plateType) values ('" + name + "','" + imgUrl + "','" + content + "',0,1,0,now(),0,0)") {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":"","msg":"添加板块成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":"","msg":"添加板块失败"}`))
	}
}

//获取板块列表
func GetTopicalList(c echo.Context) error {
	if result, err := sql.Sql_dql("select a.ID,a.content,a.name,a.view,a.good,b.imgUrl,a.userID,a.date,a.good from plate a,user_info b where a.userID=b.ID and a.plateType=1"); err != nil || result[0][0] == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"查询数据失败"}`))
	} else {
		datas := *new([]map[string]string)
		for _, v := range result {
			data := make(map[string]string)
			data["topicalID"] = v[0]
			data["content"] = v[1]
			data["title"] = v[2]
			data["view"] = v[3]
			data["good"] = v[4]
			data["imgUrl"] = v[5]
			data["userID"] = v[6]
			data["time"] = v[7]
			datas = append(datas, data)
		}
		str, _ := json.Marshal(datas)
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":`+string(str)+`,"msg":"查询数据成功"}`))
	}
}

// 获取所有的经验
func GetAllShare(c echo.Context) error {
	//连表查询获取个人分享的内容
	result, err := sql.Sql_dql("select a.ID,a.date,a.title,u.ID,b.name,a.content,a.imgUrl,a.view,a.good,(select count(ID) from comment where commentType=1 and shareID=a.ID) from share a,user_info u,plate b where a.userID=u.ID and a.plateID=b.ID")
	if err != nil {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"读取数据库失败"}`))
	}
	datas := new([]map[string]string)
	//遍历result
	for _, v := range result {
		data := make(map[string]string)
		//判断类型获取标题和id
		data["id"] = v[0]
		data["time"] = common.TimeChange(v[1])
		data["title"] = v[2]
		data["userID"] = v[3]
		data["plate"] = v[4]
		data["content"] = v[5]
		data["img"] = tools.GetDefaultImg(v[6])
		data["rowImg"] = v[6]
		data["total"] = v[8]
		data["view"] = v[7]
		data["comments"] = v[9]
		*datas = append(*datas, data)
	}
	//解析为json数据
	str, _ := json.Marshal(datas)
	//判断数据是否为空
	if result[0][0] == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"获取数据失败"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":`+string(str)+`,"msg":"获取数据成功"}`))
	}
}

//获取各个学院的人数分布
func DataGetCollege(c echo.Context) error {
	if result, err := sql.Sql_dql("select college,count(ID) from user_info group by college"); err == nil && result[0][0] != "" {
		datas := new([]map[string]string)
		//遍历result
		for _, v := range result {
			data := make(map[string]string)
			//判断类型获取标题和id
			data["name"] = v[0]
			data["values"] = v[1]
			*datas = append(*datas, data)
		}
		str, _ := json.Marshal(datas)
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":`+string(str)+`,"msg":"获取数据成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"获取数据失败"}`))
	}
}

//获取经验分布数据
func DataGetShare(c echo.Context) error {
	if result, err := sql.Sql_dql("select a.title,a.view,a.good,(select count(ID)  from comment where shareID=a.ID and commentType=1) from share a"); err == nil && result[0][0] != "" {
		datas := *new([][]string)
		datas = append(datas, []string{"经验标题", "浏览数", "点赞数", "评论数"})
		//遍历result
		for _, v := range result {
			datas = append(datas, []string{v[0], v[1], v[2], v[3]})
		}
		str, _ := json.Marshal(datas)
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":`+string(str)+`,"msg":"获取数据成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"获取数据失败"}`))
	}
}

//获取话题的分布数据
func DataGetTopical(c echo.Context) error {
	if result, err := sql.Sql_dql("select a.name,a.view,a.good,(select count(ID)  from comment where shareID=a.ID and commentType=2) from plate a where a.plateType=1 order by a.view desc limit 0,10"); err == nil && result[0][0] != "" {
		datas := *new([][]string)
		topical := *new([]string)
		view := *new([]string)
		good := *new([]string)
		comment := *new([]string)
		//遍历result
		for _, v := range result {
			topical = append(topical, v[0])
			view = append(view, v[1])
			good = append(good, v[2])
			comment = append(comment, v[3])
		}
		datas = append(datas, topical)
		datas = append(datas, good)
		datas = append(datas, comment)
		datas = append(datas, view)
		str, _ := json.Marshal(datas)
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":`+string(str)+`,"msg":"获取数据成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"获取数据失败"}`))
	}
}
