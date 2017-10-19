package mac

import (
	"fmt"
	"os/exec"
)

func Notify(title string, msg string) {
	_, err := exec.Command("terminal-notifier",
		"-title", title,
		"-message", msg,
	).Output()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Installing terminal-notifier may solbe the errors. 'brew install terminal-notifier'")
	}
}
