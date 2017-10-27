package mac

import (
	"fmt"
	"os/exec"
	"peaberry/config"
)

var conf = config.GetInstance().Mac

func Notify(title string, msg string) {
	args := []string{}
	args = append(args, "-title", title)
	args = append(args, "-message", msg)
	if conf.SoundFlag {
		args = append(args, "-sound", "default")
	}
	_, err := exec.Command("terminal-notifier",
		args...,
	).Output()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Installing terminal-notifier may solbe the errors. 'brew install terminal-notifier'")
	}
}
