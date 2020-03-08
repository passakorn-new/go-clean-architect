package models

import (
	"time"
)

type Introduction struct {
	Birthday    time.Time `json:"birthday"`
	Information string    `json:"information"`
}
