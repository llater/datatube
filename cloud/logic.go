package main

import (
	"context"
	"fmt"
	"github.com/llater/datatube/signals"
	"github.com/twitchtv/twirp"
	"math"
	"net/http"
)

type Channel struct {
	Number   int32   `json:"number"`
	Strength float64 `json:"strength"`
}

func ChannelsBusierThan(val float64) ([]*Channel, error) {
	client := signals.NewBusierThanChannelsServiceProtobufClient("http://10.10.0.10:4040", &http.Client{})
	resp, err := client.BusierThanChannels(context.Background(), &signals.BusierThanChannelsRequest{Lowerbound: 0.1})
	if err != nil {
		if twerr, ok := err.(twirp.Error); ok {
			switch twerr.Code() {
			case twirp.InvalidArgument:
				fmt.Println("invalid argument")
				return nil, err.(twirp.Error)
			default:
				return nil, err.(twirp.Error)
			}
		}
	}
	res := []*Channel{}
	for _, r := range resp.Channels {
		res = append(res, &Channel{r.Number, r.Strength})
	}
	return res, nil
}

func ChannelsQuieterThan(val float64) ([]*Channel, error) {
	client := signals.NewQuieterThanChannelsServiceProtobufClient("http://10.10.0.10:4040", &http.Client{})
	resp, err := client.QuieterThanChannels(context.Background(), &signals.QuieterThanChannelsRequest{Upperbound: 3.01})
	if err != nil {
		if twerr, ok := err.(twirp.Error); ok {
			switch twerr.Code() {
			case twirp.InvalidArgument:
				fmt.Println("invalid argument")
				return nil, err.(twirp.Error)
			default:
				return nil, err.(twirp.Error)
			}
		}
	}
	res := []*Channel{}
	for _, r := range resp.Channels {
		res = append(res, &Channel{r.Number, r.Strength})
	}
	return res, nil
}

func ChannelStrength(channel int32) (float64, error) {
	client := signals.NewChannelStrengthServiceProtobufClient("http://10.10.0.10:4040", &http.Client{})
	resp, err := client.ChannelStrength(context.Background(), &signals.ChannelStrengthRequest{channel})
	if err != nil {
		if twerr, ok := err.(twirp.Error); ok {
			switch twerr.Code() {
			case twirp.InvalidArgument:
				fmt.Println("invalid argument")
				return math.MaxFloat64, err.(twirp.Error)
			default:
				return math.MaxFloat64, err.(twirp.Error)
			}
		}
	}
	fmt.Println(resp.Strength)
	return resp.Strength, nil
}
