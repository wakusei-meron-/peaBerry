package slack

import (
	"testing"
)

func TestNotify(t *testing.T) {
	Notify("title1", "hoge")
}
