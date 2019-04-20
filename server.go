package main

import (
	"github.com/serhio83/api/app"
)

func main() {
	a := app.App{}
	a.Initialize()
	a.Run()
}
