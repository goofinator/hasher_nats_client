package nats

import (
	"github.com/goofinator/hasher_nats_client/internal/api"
	"github.com/goofinator/hasher_nats_client/internal/init/startup"
)

// NewHaser returns new Hasher intity
func NewHaser(iniData *startup.IniData) api.Hasher {
	return &hasher{iniData: iniData}
}

type hasher struct {
	iniData *startup.IniData
}

// RequestHashes gets hashes of message
// from the server via nats
func (h *hasher) RequestHashes(message []byte) ([][]byte, error) {
	return [][]byte{[]byte("PLUG"), []byte("plug")}, nil
}
