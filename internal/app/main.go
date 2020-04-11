package app

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/goofinator/hasher_nats_client/internal/api"
)

// Process performs main cycle
func Process(hasher api.Hasher) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input lines to hash.")

	for {
		msg := getLine(reader)
		hashes, err := hasher.RequestHashes(msg)
		display(hashes, err)
	}

}

func display(hashes [][]byte, err error) {
	if err != nil {
		fmt.Printf("error on hashes request: %s\n", err)
	}

	for i, hash := range hashes {
		fmt.Printf("hash %3d: %X\n", i+1, hash)
	}
}

func getLine(reader *bufio.Reader) []byte {
	fmt.Print("-> ")

	msg, err := reader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}

	return bytes.TrimRight(msg, "\n")
}
