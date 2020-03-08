package models

type Project struct {
	Name        string `json:"name" validate:"required"`
	Information string `json:"information"`
}
