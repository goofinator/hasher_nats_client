package startup

import (
	"flag"
)

// Default values of IniData fields
const (
	DefaultURL        = "localhost:4222"
	DefaultChanelName = "hasher"
)

// IniData structure stores initial data to start a app
type IniData struct {
	URL        string
	ChanelName string
}

// Configuration returns port to use obtained from user or DefaultPort
func Configuration() *IniData {
	iniData := &IniData{}
	flag.StringVar(&iniData.URL, "url", DefaultURL, "url of nats service")
	flag.StringVar(&iniData.ChanelName, "chanel", DefaultChanelName, "chanel to use")

	flag.Parse()
	return iniData
}
