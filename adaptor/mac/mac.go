package mac

import (
	"fmt"
	"os/exec"
)

func Notify(title string, msg string, soundFlag bool) {
	args := []string{}
	args = append(args, "-title", title)
	args = append(args, "-message", msg)
	if soundFlag {
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
