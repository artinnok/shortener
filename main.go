package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"shortener/web"
)

type Response struct {
	Success bool        `json:"success"`
	Payload interface{} `json:"payload"`
}

func hello(c echo.Context) error {
	username := c.Param("username")
	payload := fmt.Sprintf("Hello, %s!", username)
	response := Response{Success: true, Payload: payload}

	return c.JSON(http.StatusOK, &response)
}

func shortify(c echo.Context) (err error) {
	form := new(web.ShortifyForm)
	if err = c.Bind(form); err != nil {
		response := Response{Success: false, Payload: "invalid_data"}
		return c.JSON(http.StatusForbidden, &response)
	}

	db := web.GetDb()
	url := web.URLModel{LongLink: form.LongLink}

	db.Create(&url)
	defer db.Close()

	response := Response{Success: true, Payload: url}
	return c.JSON(http.StatusOK, &response)
}

func longify(c echo.Context) (err error) {
	response := Response{Success: false, Payload: "longify"}
	return c.JSON(http.StatusOK, &response)
}

func main() {
	e := echo.New()

	db := web.GetDb()

	db.AutoMigrate(&web.URLModel{})
	defer db.Close()

	e.GET("/hello/:username", hello)
	e.POST("/shortify", shortify)
	e.POST("/longify", longify)

	e.Logger.Fatal(e.Start("127.0.0.1:9000"))
}
