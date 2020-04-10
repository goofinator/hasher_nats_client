package app

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

// Process performs main cycle
func Process() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input lines to hash.")

	for {
		msg := getLine(reader)
		fmt.Printf("ECHO: %s\n", string(msg))
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
