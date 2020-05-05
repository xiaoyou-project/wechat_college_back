package card

import (
	"../../common"
	"../sql"
	"../tools"
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

/**
获取我的打卡
*/
func GetCardList(c echo.Context) error {
	//获取参数
	userId := c.FormValue("userID")
	//判断参数是否为空
	if userId == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//查询我的打卡数据
	//连表查询获取个人分享的内容
	result, err := sql.Sql_dql("select a.ID,b.imgUrl,b.name,a.date,a.title,a.content,c.keepDay,a.totalDay,(select count(ID) from user_card where cardID=a.ID) from card a,user_info b,user_card c where a.userID=b.ID and a.ID=c.cardID and c.userID=" + userId)
	if err != nil {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"读取数据库失败"}`))
	}
	datas := new([]map[string]string)
	//遍历result
	for _, v := range result {
		data := make(map[string]string)
		//判断类型获取标题和id
		data["id"] = v[0]
		data["imgUrl"] = v[1]
		data["name"] = v[2]
		data["time"] = common.TimeChange(v[3])
		data["title"] = v[4]
		data["description"] = tools.GetDescription(v[5])
		data["keepDay"] = v[6]
		data["totalDay"] = v[7]
		data["peoples"] = v[8]
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
获取所有的打卡数据
*/
func GetAllCard(c echo.Context) error {
	userId := c.FormValue("userID")
	if userId == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//连表查询获取个人分享的内容
	result, err := sql.Sql_dql("select a.ID,b.imgUrl,b.name,a.date,a.title,a.content,a.keepDay,a.totalDay,(select count(ID) from user_card where cardID=a.ID),b.ID from card a,user_info b where a.userID=b.ID")
	if err != nil {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"读取数据库失败"}`))
	}
	datas := new([]map[string]string)
	//遍历result
	for _, v := range result {
		data := make(map[string]string)
		//判断类型获取标题和id
		data["id"] = v[0]
		data["imgUrl"] = v[1]
		data["name"] = v[2]
		data["time"] = common.TimeChange(v[3])
		data["title"] = v[4]
		data["description"] = v[5]
		//判断用户是否加了这个打卡
		if result, err := sql.Sql_dql("select ID,keepDay from user_card where userID=" + userId + " and cardID=" + v[0]); err == nil && result[0][0] != "" {
			data["keepDay"] = result[0][1]
		} else {
			data["keepDay"] = "-1"
		}
		data["totalDay"] = v[7]
		data["peoples"] = v[8]
		data["userID"] = v[9]
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
发起新的打卡
*/
func ReleaseCard(c echo.Context) error {
	userID := c.FormValue("userID")
	name := c.FormValue("name")
	content := c.FormValue("content")
	totalDay := c.FormValue("totalDay")
	//判断参数是否为空
	if userID == "" || name == "" || content == "" || totalDay == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//发起打卡
	//判断结构
	if result, id := sql.Sql_dml_id("insert into card (`title`,`content`,`totalDay`,`keepDay`,`userID`,`date`) values ('" + name + "','" + content + "'," + totalDay + ",0," + userID + ",now())"); result {
		//成功后把自己的打卡也放进去
		sql.Sql_dml("insert into user_card (`userID`,`cardID`,`keepDay`,`lastTime`) values (" + userID + "," + id + ",0,date_sub(now(),interval 1 day))")
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":{},"msg":"发起打卡成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"发起打卡失败"}`))
	}
}

/**
获取打卡内容
*/
func GetCardContent(c echo.Context) error {
	userID := c.FormValue("userID")
	id := c.FormValue("cardID")
	//判断参数是否为空
	if userID == "" || id == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//获取打卡的主要参数
	data := make(map[string]string)
	result, err := sql.Sql_dql("select a.name,a.imgUrl,b.title,b.userId,b.date,b.content,b.totalDay,b.keepDay,(select count(ID) from user_card where cardID=b.ID) from user_info a,card b where b.userID=a.ID and b.ID=" + id)
	if err != nil {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"读取数据库失败"}`))
	}
	//判断类型获取标题和id
	data["name"] = result[0][0]
	data["imgUrl"] = result[0][1]
	data["title"] = result[0][2]
	data["userID"] = result[0][3]
	data["time"] = result[0][4]
	data["description"] = result[0][5]
	data["totalDay"] = result[0][6]
	data["keepDay"] = result[0][7]
	data["peoples"] = result[0][8]
	//判断用户是否打卡
	if result, _ := sql.Sql_dql("select ID,keepDay,lastTime from user_card where userID=" + userID + " and cardID=" + id); result != nil && result[0][0] != "" {
		data["cardStatus"] = "1"
		data["keepDay"] = result[0][1]
		//判断今天是否打卡
		loc, _ := time.LoadLocation("Local")
		theTime, err := time.ParseInLocation("2006-01-02 15:04:05", result[0][2], loc)
		if err == nil {
			now := time.Now()
			if now.Month() == theTime.Month() && now.Day() == theTime.Day() {
				data["nowStatus"] = "1"
			} else {
				data["nowStatus"] = "0"
			}
		} else {
			data["nowStatus"] = "0"
		}
	} else {
		data["cardStatus"] = "0"
		data["nowStatus"] = "0"
	}
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
完成打卡
*/
func FinishCard(c echo.Context) error {
	userID := c.FormValue("userID")
	id := c.FormValue("cardID")
	content := c.FormValue("content")
	//判断参数是否为空
	if userID == "" || id == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//更新打卡天数
	if sql.Sql_dml("update user_card set keepDay=keepDay+1,lastTime=now() where cardID=" + id + " and userID=" + userID) {
		//发表发布后的感想
		if content != "" {
			sql.Sql_dml("insert into comment (`shareID`,`commentType`,`userID`,`content`,`date`,`good`) values (" + id + ",3," + userID + ",'" + content + "',now(),0)")
		}
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":[],"msg":"打卡成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"打卡失败"}`))
	}
}

/**
加入打卡
*/
func JoinCard(c echo.Context) error {
	userID := c.FormValue("userID")
	id := c.FormValue("cardID")
	//判断参数是否为空
	if userID == "" || id == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//更新打卡天数
	if sql.Sql_dml("insert into user_card (`userID`,`cardID`,`keepDay`,`lastTime`) values (" + userID + "," + id + ",0,date_sub(now(),interval 1 day))") {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":[],"msg":"加入打卡成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"加入打卡失败"}`))
	}
}

/**
删除打卡
*/
func DeleteCard(c echo.Context) error {
	userID := c.FormValue("userID")
	id := c.FormValue("cardID")
	//判断参数是否为空
	if userID == "" || id == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//更新打卡天数
	if sql.Sql_dml("delete from card where ID=" + id + " and userID=" + userID) {
		//删除所有的用户打卡记录
		sql.Sql_dml("delete from user_card where cardID=" + id)
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":[],"msg":"删除打卡成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"删除打卡失败"}`))
	}
}

/**
退出打卡
*/
func AbortCard(c echo.Context) error {
	userID := c.FormValue("userID")
	id := c.FormValue("cardID")
	//判断参数是否为空
	if userID == "" || id == "" {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":{},"msg":"参数错误"}`))
	}
	//删除用户的打卡记录
	if sql.Sql_dml("delete from user_card where cardID=" + id + " and userID=" + userID) {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":1,"data":[],"msg":"退出打卡成功"}`))
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"code":0,"data":[],"msg":"退出打卡失败"}`))
	}
}
