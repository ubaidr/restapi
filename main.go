package main

import (
	"github.com/labstack/echo"
	"net/http"
)

type event struct {
	Title         string `json:""`

}

type allevent [] event
var events=allevent{
	{
		Title: "Hello",
	},
}

func main() {
	e := echo.New()
	e.GET("/greetings", func(c echo.Context) error {
		return c.JSON(http.StatusCreated,events)
	})
	e.Logger.Fatal(e.Start(":8081"))
}
