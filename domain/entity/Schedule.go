package entity

import "time"

type Schedule struct {
	Title       string    `json:"title"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	Done		bool
}
