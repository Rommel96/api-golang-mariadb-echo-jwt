package routes

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/rommel96/api-golang-mariadb-echo-jwt/models"
)

var Receta models.Recetas

func Recetas(e *echo.Echo) {
	e.GET("/recetas", getAllRecetas)
	e.GET("/recetas/:id", getReceta)
	e.POST("/recetas", nuevaReceta)
	e.PUT("/recetas/:id", editReceta)
	e.DELETE("/recetas/:id", deleteReceta)
}

func getAllRecetas(c echo.Context) error {
	recetas := models.GetAllRecetas()
	return c.JSON(http.StatusOK, recetas)
}

func getReceta(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	r := models.GetRecetaID(id)
	return c.JSON(http.StatusOK, r)
}

func nuevaReceta(c echo.Context) error {
	if err := c.Bind(&Receta); err != nil {
		return err
	}
	r, err := Receta.CrearReceta()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, r)
}

func editReceta(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.Bind(&Receta); err != nil {
		return err
	}
	r, err := Receta.UpdateReceta(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, r)
	}
	return c.JSON(http.StatusOK, r)
}

func deleteReceta(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	r := models.DeleteReceta(id)
	return c.JSON(http.StatusOK, r)
}
