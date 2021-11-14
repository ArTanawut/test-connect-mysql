package main

import (
	"fmt"
	"net/http"

	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("No Port In Heroku :" + port)
	}
	return ":" + port
}

func main() {
	e := echo.New()

	InitialDb()
	e.GET("/helpcheck/:name", func(c echo.Context) error {
		name := c.Param("name")
		return c.String(http.StatusOK, name)
	})
	//e.POST("/getuserlogin", getuserlogin)
	e.POST("/insertuserlogin", insertuserlogin)
	//e.POST("/insertuserlog", insertuserlog)
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.CORSWithConfig(
			middleware.CORSConfig{
				AllowOrigins:     []string{"*"},
				AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions, http.MethodHead},
				AllowCredentials: true,
				AllowHeaders:     []string{echo.HeaderAuthorization, echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
			},
		),
	)
	// e.Logger.Fatal(e.Start(":8080"))
	e.Logger.Fatal(e.Start(getPort()))

}
