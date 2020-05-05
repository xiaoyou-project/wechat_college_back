package route

import (
	"../functions/admin"
	"../functions/card"
	"../functions/comment"
	"../functions/plate"
	"../functions/topical"
	"../functions/upload"
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
	e.GET("/api/topical/get/topicalList", topical.GetTopicalList)
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
	/*搜索话题*/
	e.GET("/api/topical/get/search/topicalList", topical.SearchTopical)

	/*获取评论内容*/
	e.GET("/api/comment/get/commentList", comment.GetCommentList)
	/*删除评论*/
	e.POST("/api/comment/delete", comment.DeleteComment)
	/*发表评论*/
	e.POST("/api/comment/publish", comment.PublishComment)
	/*评论点赞或者取消赞*/
	e.POST("/api/comment/update/good", comment.UpdateGood)
	/*判断评论类型*/
	e.GET("/api/comment/type", comment.CommentType)

	/*获取所有板块*/
	e.GET("/api/plate/get/plateList", plate.GetPlateList)
	/*获取经验列表*/
	e.GET("/api/plate/get/shareList", plate.GetShareList)
	/*获取经验内容*/
	e.GET("/api/plate/get/shareContent", plate.GetShareContent)
	/*修改经验内容*/
	e.POST("/api/plate/edit/shareContent", plate.EditShareContent)
	/*点赞或者取消点赞*/
	e.POST("/api/plate/update/good", plate.UpdateGood)
	/*收藏或者取消收藏经验*/
	e.POST("/api/plate/update/collect/share", plate.UpdateCollect)
	/*删除经验内容*/
	e.POST("/api/plate/delete/shareContent", plate.DeleteShare)
	/*获取我收藏的板块*/
	e.GET("/api/plate/get/collect/plateList", plate.GetCollectPlate)
	/*收藏板块*/
	e.POST("/api/plate/update/collect", plate.CollectPlate)
	/*申请板块*/
	e.POST("/api/plate/application/plate", plate.ApplicationPlate)
	/*发布新的经验*/
	e.POST("/api/plate/release/share", plate.ReleaseNewShare)
	/*判断板块是否收藏*/
	e.GET("/api/plate/status/collect", plate.StatusCollect)
	/*搜索经验*/
	e.GET("/api/plate/get/search/shareList", plate.SearchShare)

	/*获取我的打卡*/
	e.GET("/api/card/get/me/cardList", card.GetCardList)
	/*获取所有打卡*/
	e.GET("/api/card/get/cardList", card.GetAllCard)
	/*发起新的打卡*/
	e.POST("/api/card/release", card.ReleaseCard)
	/*获取打卡内容*/
	e.GET("/api/card/get/cardContent", card.GetCardContent)
	/*完成打卡*/
	e.POST("/api/card/finish", card.FinishCard)
	/*加入打卡*/
	e.POST("/api/card/join", card.JoinCard)
	/*删除打卡*/
	e.POST("/api/card/delete", card.DeleteCard)
	/*退出打卡*/
	e.POST("/api/card/abort", card.AbortCard)

	/*上传图片并保存到本地*/
	e.POST("/api/img/upload", upload.UploadFile)

	/*注册静态文件路由*/
	e.Static("/static", "static")

	/*和admin相关的路由*/
	//判断用户是否合法
	e.POST("/api/admin/token", admin.AdminToken)
	//获取用户数，话题数，经验数还有打卡数
	e.GET("/api/admin/visualization/overall", admin.VisualizationOverall)
	//获取所有的板块列表
	e.GET("/api/admin/plate/get/plateList", admin.GetAllPlate)
	//获取所有用户
	e.GET("/api/admin/user/get/userList", admin.GetAllUser)
	//删除用户
	e.GET("/api/admin/user/delete", admin.DeleteUser)
	//删除板块
	e.GET("/api/admin/plate/delete", admin.DeletePlate)
	//修改板块信息
	e.POST("/api/admin/plate/change/plateInfo", admin.ChangePlateInfo)
	//修改板块的状态
	e.POST("/api/admin/plate/change/plateStatus", admin.ChangePlateStatus)
	//添加板块
	e.POST("/api/admin/plate/add", admin.AddPlate)
	//获取所有经验
	e.GET("/api/admin/plate/get/shareList", admin.GetAllShare)

	//修改话题内容
	e.POST("/api/admin/topical/edit", admin.ChangeTopicalContent)
	//获取话题列表
	e.GET("/api/admin/topical/get/topicalList", admin.GetTopicalList)
	//修改打卡信息
	e.POST("/api/admin/card/change/cardInfo", admin.ChangeCardInfo)
	//获取所有评论
	e.GET("/api/admin/getAllComment", admin.GetAllComment)
	//删除评论
	e.POST("/api/admin/comment/delete", admin.DeleteComment)

	//数据可视化
	e.GET("/api/admin/data/getCollege", admin.DataGetCollege)
	//获取经验分布数据
	e.GET("/api/admin/data/getShare", admin.DataGetShare)
	//获取热门话题(只显示10条)
	e.GET("/api/admin/data/getTopical", admin.DataGetTopical)
}
