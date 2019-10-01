package main

import (
	"command"
	"model"
)

func main() {
	model.Init()
	command.Execute()
}
