package remotes

import (
	"github.com/goofinator/hasher_nats_client/internal/api"
	"github.com/goofinator/hasher_nats_client/internal/init/startup"
	"github.com/nats-io/nats.go"
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
	nc, err := nats.Connect(h.iniData.URL)
	if err != nil {
		return nil, err
	}
	defer nc.Close()

	if err := nc.Publish(h.iniData.ChanelName, message); err != nil {
		return nil, err
	}
	
	return [][]byte{[]byte("PLUG"), []byte("plug")}, nil
}
