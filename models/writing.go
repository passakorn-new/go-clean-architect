package models

type Writing struct {
	Topic       string `json:"topic" validate:"required"`
	Information string `json:"information"`
}
