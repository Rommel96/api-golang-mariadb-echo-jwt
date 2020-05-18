package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/rommel96/api-golang-mariadb-echo-jwt/auth"
	"github.com/rommel96/api-golang-mariadb-echo-jwt/models"
	"github.com/rommel96/api-golang-mariadb-echo-jwt/utils"
)

type userLogin struct {
	Username string `json:"user"`
	Password string `json:"pass"`
}

func Home(e *echo.Echo) {
	e.POST("/signup", signup)
	e.POST("/login", login)
	r := e.Group("/api")
	r.Use(auth.SetMiddleware())
	r.GET("/auth", welcome)
}

//redirect route createUser
func signup(c echo.Context) error {
	return createUser(c)
}

func login(c echo.Context) error {
	u := userLogin{}
	if err := c.Bind(&u); err != nil {
		return err
	}
	userValid := models.VerifyUser(u.Username)
	if userValid == nil {
		return c.JSON(http.StatusNotFound, "{User not Found}")
	}
	if err := utils.Verify(userValid.Password, u.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, "{Inavlid password!!}")
	}
	token := auth.CreateToken(u.Username)
	return c.JSON(http.StatusOK, echo.Map{
		"auth":  true,
		"token": token,
	})
}

func welcome(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	decode, err := auth.DecodeToken(token)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, decode)
}
