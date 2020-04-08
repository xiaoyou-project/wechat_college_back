package route

import (
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
}
