package server

import (
	pb "github.com/matoval/airtHive/backend/proto"
)

type LLM interface {
	Lock()
	Unlock()
	Locking() bool
	Predict(*pb.PredictOptions) (string, error)
	PredictStream(*pb.PredictOptions, chan string) error
	Load(*pb.ModelOptions) error
	Status() (pb.StatusResponse, error)
}

func newReply(s string) *pb.Reply {
	return &pb.Reply{Message: []byte(s)}
}
