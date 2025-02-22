package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/matoval/airtHive/backend/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedBackendServer
	llm LLM
}

func (s *server) Health(ctx context.Context, in *pb.HealthMessage) (*pb.Reply, error) {
	return newReply("OK"), nil
}

func (s *server) LoadModel(ctx context.Context, in *pb.ModelOptions) (*pb.Result, error) {
	if s.llm.Locking() {
		s.llm.Lock()
		defer s.llm.Unlock()
	}
	err := s.llm.Load(in)
	if err != nil {
		return &pb.Result{Message: fmt.Sprintf("Error loading model: %s", err.Error()), Success: false}, err
	}
	return &pb.Result{Message: "Loading succeeded", Success: true}, nil
}

func (s *server) Predict(ctx context.Context, in *pb.PredictOptions) (*pb.Reply, error) {
	if s.llm.Locking() {
		s.llm.Lock()
		defer s.llm.Unlock()
	}
	result, err := s.llm.Predict(in)
	return newReply(result), err
}

func (s *server) PredictStream(in *pb.PredictOptions, stream pb.Backend_PredictStreamServer) error {
	if s.llm.Locking() {
		s.llm.Lock()
		defer s.llm.Unlock()
	}
	resultChan := make(chan string)

	done := make(chan bool)
	go func() {
		for result := range resultChan {
			stream.Send(newReply(result))
		}
		done <- true
	}()

	err := s.llm.PredictStream(in, resultChan)
	<-done

	return err
}

func (s *server) Status(ctx context.Context, in *pb.HealthMessage) (*pb.StatusResponse, error) {
	res, err := s.llm.Status()
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func StartServer(address string, model LLM) error {
	listen, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("Cannot create listener: %s", err)
	}

	s := grpc.NewServer()
	pb.RegisterBackendServer(s, &server{llm: model})
	log.Printf("gRPC server listening at %v", listen.Addr())
	if err = s.Serve(listen); err != nil {
		return err
	}

	return nil
}
