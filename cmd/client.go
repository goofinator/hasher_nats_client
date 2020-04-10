package main

import (
	"github.com/goofinator/hasher_nats_client/internal/app"
	"github.com/goofinator/hasher_nats_client/internal/init/startup"
	"github.com/goofinator/hasher_nats_client/internal/remotes"
)

func main() {
	iniData := startup.Configuration()
	hasher := remotes.NewHaser(iniData)
	app.Process(hasher)
}
