package route

import (
	"../functions/topical"
	"../functions/user"
	"github.com/labstack/echo"
)

func Route(e *echo.Echo) {
	/*获取用户的appid*/
	e.GET("/api/user/get/openid", user.GetOpenId)
	/*用户注册*/
	e.POST("/api/user/registered", user.UserRegistered)
	/*获取用户信息*/
	e.GET("/api/user/get/userInfo", user.GetUserInfo)
	/*修改个人中心的用户信息*/
	e.POST("/api/user/update/userInfo", user.UpdateUserInfo)
	/*进入个人中心获取分享的经验列表*/
	e.GET("/api/user/get/shareList", user.GetShareList)
	/*进入个人中心获取打卡列表*/
	e.GET("/api/user/get/cardList", user.GetCardList)
	/*个人中心获取话题列表*/
	e.GET("/api/user/get/topicalList", user.GetTopicalList)
	/*获取个人收藏的话题*/
	e.GET("/api/user/get/collect/topicalList", user.GetCollectTopicalList)
	/*获取个人收藏的经验*/
	e.GET("/api/user/get/collect/shareList", user.GetCollectShareList)
	/*获取收到的评论消息*/
	e.GET("/api/user/get/comment/messageList", user.GetCommentMessage)
	/*获取收到的赞的消息*/
	e.GET("/api/user/get/good/messageList", user.GetGoodMessage)
	/*获取收到的系统消息*/
	e.GET("/api/user/get/system/messageList", user.GetSystemMessage)
	/*修改消息的状态*/
	e.POST("/api/user/set/messageStatus", user.SetMessageStatus)

	/*获取话题列表*/
	e.GET("/api/user/get/system/messageList", topical.GetTopicalList)
	/*发布话题*/
	e.POST("/api/topical/release", topical.ReleaseTopical)
	/*获取话题的内容*/
	e.GET("/api/topical/get/content", topical.GetTopicalContent)
	/*话题点赞或者取消赞*/
	e.POST("/api/topical/update/good", topical.TopicalGood)
	/*话题收藏或者取消收藏*/
	e.POST("/api/topical/update/collect", topical.UpdateCollect)
	/*删除话题*/
	e.POST("/api/topical/delete", topical.DeleteTopical)
	/*获取评论内容*/
	e.GET("/api/comment/get/commentList", topical.GetCommentList)
}
