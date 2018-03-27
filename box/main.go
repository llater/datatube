package main

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	pb "github.com/llater/datatube/signals"
	"net/http"
	"sync"
)

type Channel struct {
	Number      int32
	Strength    float64
	Quality     float64
	MinLastHour float64
	MaxLastHour float64
}

type Store struct {
	Channels []*Channel
	sync.Mutex
}

var store Store

type ChannelDataServer struct{}

func (s *ChannelDataServer) BusierThanChannels(ctx context.Context, req *pb.BusierThanChannelsRequest) (*pb.BusierThanChannelsResponse, error) {
	l := req.Lowerbound
	res := []*pb.BusierThanChannelsResponse_Channel{}
	// iterate through stored channels
	for _, c := range store.Channels {
		if c.Strength > l {
			// return the channel if it's strength is over the lowerbound
			res = append(res, &pb.BusierThanChannelsResponse_Channel{c.Number, c.Strength})
		}
	}
	return &pb.BusierThanChannelsResponse{res}, nil
}

func (s *ChannelDataServer) QuieterThanChannels(ctx context.Context, req *pb.QuieterThanChannelsRequest) (*pb.QuieterThanChannelsResponse, error) {
	u := req.Upperbound
	res := []*pb.QuieterThanChannelsResponse_Channel{}
	for _, c := range store.Channels {
		if c.Strength < u {
			res = append(res, &pb.QuieterThanChannelsResponse_Channel{c.Number, c.Strength})
		}
	}
	return &pb.QuieterThanChannelsResponse{res}, nil
}

func (s *ChannelDataServer) ChannelStrength(ctx context.Context, req *pb.ChannelStrengthRequest) (*pb.ChannelStrengthResponse, error) {
	num := req.Channel
	for _, c := range store.Channels {
		if num == c.Number {
			return &pb.ChannelStrengthResponse{c.Strength, c.Quality}, nil
		}
	}
	return nil, errors.New("channel requested not present")
}

func newChannelDataRouter() *mux.Router {
	server := &ChannelDataServer{}
	busierHandler := pb.NewBusierThanChannelsServiceServer(server, nil)
	quieterHandler := pb.NewQuieterThanChannelsServiceServer(server, nil)
	channelStrengthHandler := pb.NewChannelStrengthServiceServer(server, nil)
	r := mux.NewRouter()
	r.Handle("/busier-than", busierHandler)
	r.Handle("/quieter-than", quieterHandler)
	r.Handle("/channel-strength", channelStrengthHandler)
	return r
}

func main() {
	r := newChannelDataRouter()
	// in a working example this data would come from SDR
	c := []*Channel{}
	c = append(c, &Channel{1, 0.299999999, 4.2, 0.2233333333, 0.3555555555})
	c = append(c, &Channel{2, 1.2472572573, 5.6, 1.000002, 1.250000000000})
	c = append(c, &Channel{3, -0.111222222, 10.0, -2.777777777777, 1.1100000})
	c = append(c, &Channel{4, 3.9999999999, 5.5, 0.0000000000000, 16.0})
	http.ListenAndServe(":4040", r)
}
