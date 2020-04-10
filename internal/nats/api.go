package nats

import "github.com/goofinator/hasher_nats_client/internal/api"

// NewHaser returns new Hasher intity
func NewHaser() api.Hasher {
	return &hasher{}
}

type hasher struct{}

// RequestHashes gets hashes of message
// from the server via nats
func (h *hasher) RequestHashes(message []byte) ([][]byte, error) {
	return [][]byte{[]byte("PLUG"), []byte("plug")}, nil
}
