package main

import (
	"net/http"
	"fmt"

	"github.com/labstack/echo"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type ShortifyForm struct {
	Link string `json:"link"`
}

type Response struct {
	Success bool `json:"success"`
	Payload interface{} `json:"payload"`
}

type URLModel struct {
	gorm.Model
	Link string
	ShortLink string
}

func getDb() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5435 user=short dbname=short password=short sslmode=disable")
	if err != nil {
		panic(err)
	}

	return db
}

func hello(c echo.Context) error {
	username := c.Param("username")
	payload := fmt.Sprintf("Hello, %s!", username)
	response := Response{Success: true, Payload: payload}

	return c.JSON(http.StatusOK, &response)
}

func shortify(c echo.Context) (err error) {
	form := new(ShortifyForm)
	if err = c.Bind(form); err != nil {
		response := Response{Success: false, Payload: "invalid_data"}
		return c.JSON(http.StatusForbidden, &response)
	}

	db := getDb()
	url := URLModel{Link: form.Link}

	db.Create(&url)
	defer db.Close()

	response := Response{Success: true, Payload: url}
	return c.JSON(http.StatusOK, &response)
}

func main() {
	e := echo.New()

	db := getDb()

	db.AutoMigrate(&URLModel{})
	defer db.Close()

	e.GET("/hello/:username", hello)
	e.POST("/shortify", shortify)

	e.Logger.Fatal(e.Start("127.0.0.1:9000"))
}
