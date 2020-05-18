package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/rommel96/api-golang-mariadb-echo-jwt/models"
	"github.com/rommel96/api-golang-mariadb-echo-jwt/utils"
)

func Users(e *echo.Echo) {
	e.POST("/users", createUser)
	e.GET("/users", getAllUsers)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.GetUserID(id)
	return c.JSON(http.StatusOK, user)
}

func getAllUsers(c echo.Context) error {
	users := models.GetAllUsers()
	return c.JSON(http.StatusOK, users)
}

func createUser(c echo.Context) error {
	newUser := &models.Users{}
	if err := c.Bind(newUser); err != nil {
		log.Println(err)
		return err
	}
	hashedPasswd, err := utils.Encrypt((newUser.Password))
	if err != nil {
		return err
	}
	newUser.Password = string(hashedPasswd)
	newUser.CreateUser()
	return c.JSON(http.StatusCreated, newUser)
}

func updateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := &models.Users{}
	if err := c.Bind(user); err != nil {
		return err
	}
	err := user.UpdateUser(id)
	return c.JSON(http.StatusOK, err)
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result := models.DeleteUser(id)
	return c.JSON(http.StatusOK, result)
}
