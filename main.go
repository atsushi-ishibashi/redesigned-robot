package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	port := flag.String("port", "8080", "port to listen")
	flag.Parse()

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			req := c.Request()
			return req.URL.Path == "/healthcheck"
		},
	}))

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	e.GET("/users/:user_id", func(c echo.Context) error {
		userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "invalid param",
			})
		}
		return c.JSON(http.StatusOK, generateUserInfo(userID))
	})

	e.Logger.Fatal(e.Start(":" + *port))
}

type UserInfo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func generateUserInfo(userID int64) UserInfo {
	return UserInfo{
		ID:   userID,
		Name: fmt.Sprintf("ユーザー%d", userID),
	}
}
