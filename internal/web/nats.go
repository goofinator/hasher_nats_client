package web

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/goofinator/hasher_nats_client/internal/api"
	"github.com/goofinator/hasher_nats_client/internal/init/startup"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"

	"github.com/goofinator/hasher_nats_server/pkg"
)

const (
	waitTime = time.Second * 5
)

// NewHaser returns new Hasher intity
func NewHaser(natsSettings startup.NatsSettings) api.Hasher {
	return hasher{natsSettings: natsSettings}
}

type hasher struct {
	natsSettings startup.NatsSettings
}

// RequestHashes gets hashes of message
// from the server via nats
func (h hasher) RequestHashes(message []byte) ([][]byte, error) {
	if len(message) == 0 {
		return nil, fmt.Errorf("skip empty line")
	}
	nc, err := nats.Connect(h.natsSettings.URL)
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

func (h hasher) useConnection(c *nats.EncodedConn, message []byte) ([][]byte, error) {
	msg, subjectBase := h.prepareMessage(message)
	if err := c.Publish(subjectBase+".out", msg); err != nil {
		return nil, err
	}

	ch := make(chan pkg.Message)
	sub, err := c.BindRecvChan(subjectBase+".in", ch)
	if err != nil {
		return nil, err
	}
	defer sub.Unsubscribe()

	select {
	case res := <-ch:
		return decodeResult(res)
	case <-time.After(waitTime):
		break
	}

	return nil, fmt.Errorf("Exit by timeout")
}

func decodeResult(msg pkg.Message) ([][]byte, error) {
	result := make([][]byte, 0)

	if err := json.Unmarshal(msg.Body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (h hasher) prepareMessage(message []byte) (pkg.Message, string) {
	msg := pkg.Message{
		Sender: h.natsSettings.UUID,
		ID:     uuid.New(),
		Type:   pkg.DefaultMessageType,
		Body:   message,
	}
	subject := fmt.Sprintf("worker.%s", msg.Sender)
	return msg, subject
}
