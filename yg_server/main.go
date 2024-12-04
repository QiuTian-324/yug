package main

import (
	"yug_server/cmd"
	"yug_server/internal/config"
)

func main() {
	cmd.StartServer()
}

func init() {
	config.Viperinternal()
}
