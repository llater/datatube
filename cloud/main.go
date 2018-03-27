package main

import (
	"flag"
	"fmt"
	"encoding/json"
	"os"
)

type channelResponse struct {
	Channel *Channel `json:"channel"`
}

func main() {
	channelNumber := flag.Int("channel", 1, "Channel number")
	flag.Parse()
	var b = int32(*channelNumber)
	data, err := ChannelStrength(b)

	if err != nil {
		json.NewEncoder(os.Stdout).Encode(channelResponse{&Channel{b, data}})
	} else {
		fmt.Println("did not work")
	}
}