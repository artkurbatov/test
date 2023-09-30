package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	e.GET("/status", Handler)
	e.Use(MW)

	err := e.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func Handler(c echo.Context) error {
	date := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	daysLeft := int(time.Until(date).Hours()) / 24
	return c.String(http.StatusOK, fmt.Sprintf("days left: %d", daysLeft))
}

func MW(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		role := c.Request().Header.Get("User-Role")

		if role == "admin" {
			log.Println("red button user detected")
		}

		err := next(c)
		if err != nil {
			return err
		}

		return nil
	}
}
