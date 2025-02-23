package server

import (
	"log"
	"net"

	pb "github.com/matoval/airtHive/backend/proto"
	"google.golang.org/grpc"
)

func StartServer() error {
	listen, err := net.Listen("tcp", ":50111")
	if err != nil {
		log.Fatalf("Cannot create listener: %s", err)
	}

	s := grpc.NewServer()
	newHive := NewAirtHive()
	if err != nil {
		return err
	}
	pb.RegisterBackendServer(s, newHive)
	log.Printf("gRPC server listening at %v", listen.Addr())
	if err = s.Serve(listen); err != nil {
		return err
	}

	return nil
}
