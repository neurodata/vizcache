package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func test(c echo.Context) error {
	return c.JSON(http.StatusOK, "I am testing!")
}

func getNdstoreVolume(c echo.Context) error {

	cc := c.(*NDStoreContext)
	return cc.GetVolume(c.ParamValues()[0])
}

func proxyNdStoreRequest(c echo.Context) error {
	fmt.Println(c.ParamValues())
	return c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s"))
}

func main() {
	e := echo.New()

	// Custom context
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &NDStoreContext{c, nil}
			return h(cc)
		}
	})

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET},
	}))

	g := e.Group("/ndstore")

	g.GET("/volume*", getNdstoreVolume)
	g.GET("/sd/*", proxyNdStoreRequest)
	e.Logger.Fatal(e.Start(":8080"))

}
