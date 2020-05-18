package models

import (
	"encoding/json"
	"time"
)

type Recetas struct {
	Id           int             `json:"id,primary_key"`
	Nombre       string          `json:"nombre"`
	Descripcion  string          `json:"descripcion"`
	Ingredientes json.RawMessage `json:"ingredientes"`
	Preparacion  json.RawMessage `json:"preparacion"`
	Images       json.RawMessage `json:"images"`
	Votos        int             `json:"votos"`
	Comentarios  int             `json:"comentarios"`
	Created_at   time.Time       `json:"created_at"`
	CategoriaID  int             `json:"categoriaID,foreign_key"`
	UserID       int             `json:"userID,foreign_key"`
}

func (r *Recetas) CrearReceta() (*Recetas, error) {
	//db.NewRecord(&r)
	err := db.Create(&r).Error
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetAllRecetas() []Recetas {
	var recetas []Recetas
	db.Find(&recetas)
	return recetas
}

func GetRecetaID(id int) *Recetas {
	var receta Recetas
	err := db.First(&receta, id).Error
	if err != nil {
		return nil
	}
	return &receta
}

func (r *Recetas) UpdateReceta(id int) (*Recetas, error) {
	var receta Recetas
	err := db.First(&receta, id).Error
	if err != nil {
		return nil, err
	}
	db.Model(&receta).Update(r)
	return &receta, nil
}

func DeleteReceta(id int) *Recetas {
	var receta Recetas
	err := db.First(&receta, id).Error
	if err != nil {
		return nil
	}
	db.Delete(&receta)
	return &receta
}
