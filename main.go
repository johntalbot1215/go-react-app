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
	e.POST("/login", handleLookupAccount)

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
		fmt.Println("New Account Error", err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	fmt.Println("Query")
	err := insertAccount(a)
	if err != nil {
		return c.String(500, err.Error())
	}
	return c.String(http.StatusOK, "ok")
}

func handleLookupAccount(c echo.Context) error {
	a := new(Account)
	if err := c.Bind(a); err != nil {
		fmt.Println(err)
		fmt.Println("Account Lookup Error")
		return c.String(http.StatusBadRequest, err.Error())
	}
	err := lookupAccount(a)
	if err != nil {
		return c.String(500, err.Error())
	}
	return c.String(http.StatusOK, "ok")
}
