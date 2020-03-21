package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.POST("/new-account", handleNewAccount)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

type Account struct {
	Username string `json:"username" query:"name"`
	Password string `json:"password"`
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func handleNewAccount(c echo.Context) error {
	a := new(Account)
	if err := c.Bind(a); err != nil {
		fmt.Println("Err", err)
		return c.String(http.StatusBadRequest, "Error")
	}
	fmt.Println(a)
	return c.String(http.StatusOK, "ok")
}
