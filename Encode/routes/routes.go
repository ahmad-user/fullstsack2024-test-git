package routes

import (
	"Encode/controller/user"

	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, ctl user.ClientController) {
	userRoute(c, ctl)
}

func userRoute(c *echo.Echo, ctl user.ClientController) {
	c.POST("/add", ctl.AddClient())

}
