package main

import (
	"peaberry/domain"
)

func main() {
	schedMgr := domain.GetInstance()
	schedMgr.StartApplication()
}
