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

}
