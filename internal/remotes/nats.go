package remotes

import (
	"encoding/json"
	"fmt"

	"github.com/goofinator/hasher_nats_client/internal/api"
	"github.com/goofinator/hasher_nats_client/internal/init/startup"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"

	"github.com/goofinator/hasher_nats_server/pkg"
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

	c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	return h.useConnection(c, message)
}

func (h *hasher) useConnection(c *nats.EncodedConn, message []byte) ([][]byte, error) {
	msg, subjectBase := h.prepareMessage(message)
	if err := c.Publish(subjectBase+".out", msg); err != nil {
		return nil, err
	}

	ch := make(chan *pkg.Message)
	c.BindRecvChan(subjectBase+".in", ch)

	return decodeResult(<-ch)
}

func decodeResult(msg *pkg.Message) ([][]byte, error) {
	result := make([][]byte, 0)

	if err := json.Unmarshal(msg.Body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (h *hasher) prepareMessage(message []byte) (*pkg.Message, string) {
	msg := &pkg.Message{
		Sender: h.iniData.Sender,
		ID:     uuid.New(),
		Type:   pkg.DefaultMessageType,
		Body:   message,
	}
	subject := fmt.Sprintf("worker.%s", msg.Sender)
	return msg, subject
}
