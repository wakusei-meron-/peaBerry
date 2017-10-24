package entity

import (
	"peaberry/config"
	"peaberry/adaptor/mac"
)

type Notification struct {
	Title string
	Message string
}

var conf = config.GetInstance().Notification

func (n *Notification) Fire() {
	mac.Notify(n.Title, n.Message, conf.SoundFlag)
}
