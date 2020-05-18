package main

import (
	"github.com/labstack/echo"
	"github.com/rommel96/api-golang-mariadb-echo-jwt/routes"
)

func main() {
	e := echo.New()
	routes.Home(e)
	routes.Users(e)
	routes.Recetas(e)
	e.Logger.Fatal(e.Start(":3000"))
}
