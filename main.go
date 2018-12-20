package main

import (
	"net/http"
	"fmt"
	
	"github.com/labstack/echo"
)

type LongURL struct {
	URL string `json:"url"`
}

func hello(c echo.Context) error {
	username := c.Param("username")
	res := fmt.Sprintf("Hello, %s!", username)

	return c.JSON(http.StatusOK, res)
}

func shortify(c echo.Context) (err error) {
	url := &LongURL{}
	if err = c.Bind(url); err != nil {
		return echo.NewHTTPError(http.StatusForbidden, "Invalid data")
	}

	return c.JSON(http.StatusOK, url)
}

func main() {
	e := echo.New()

	e.GET("/hello/:username", hello)
	e.POST("/shortify", shortify)

	e.Logger.Fatal(e.Start(":9000"))
}
