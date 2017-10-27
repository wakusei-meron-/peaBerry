package entity

import "time"

type Schedule struct {
	Id    string
	Title string
	Start time.Time
	End   time.Time
	Done  bool
}
