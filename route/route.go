package route

import (
	"../functions/user"
	"github.com/labstack/echo"
)

func Route(e *echo.Echo) {
	/*用户注册*/
	e.POST("/api/user/registered", user.UserRegistered)

}
