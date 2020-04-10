package main

import (
	"github.com/goofinator/hasher_nats_client/internal/app"
	"github.com/goofinator/hasher_nats_client/internal/nats"
)

func main() {
	hasher := nats.NewHaser()
	app.Process(hasher)
}
