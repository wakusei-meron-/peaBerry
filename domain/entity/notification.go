package entity

import (
	"peaberry/config"
	"peaberry/adaptor/mac"
	"peaberry/adaptor/slack"
)

type Notification struct {
	Title   string
	Message string
}

var conf = config.GetInstance()

func (n *Notification) Fire() {
	slack.Notify(n.Title, n.Message)
	mac.Notify(n.Title, n.Message)
}
