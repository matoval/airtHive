package server

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"sync"

	pb "github.com/matoval/airtHive/backend/proto"
	"google.golang.org/grpc"
)

type AirtHive struct {
	pb.UnimplementedBackendServer
	sync.Mutex
}

func NewAirtHive() (*AirtHive) {
	return &AirtHive{pb.UnimplementedBackendServer{}, sync.Mutex{}}
}

func (h *AirtHive) ChatStream(options *pb.ChatOptions, stream grpc.ServerStreamingServer[pb.Reply]) error {
	cmd := exec.Command("llama-cli", "-m", options.Model, "--repeat-penalty", fmt.Sprint(options.RepeatPenalty), "-n", fmt.Sprint(options.Predict), "--rpc", options.Backends, "-ngl 99")


	for {
		stdin, err := cmd.StdinPipe()
		if err != nil {
			return err
		}
	
		go func() {
			defer stdin.Close()
			msg := ""
			stream.RecvMsg(msg)
			io.WriteString(stdin, msg)
		}()
		if stream.Context() == nil {
			break
		}
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(output)

	return nil
}

func (h *AirtHive) StartBackend(context.Context, *pb.BackendOptions) (*pb.BackendResponse, error) {
	cmd := exec.Command("rpc-server", "-p", "50052")

	output, err := cmd.Output()
	if err != nil {
		return &pb.BackendResponse{}, err
	}

	fmt.Println(output)

	res := pb.BackendResponse{
		State: 2,
		Memory: &pb.MemoryUsageData{
			Breakdown: map[string]uint64{"test": 311},
		},
	}

	return &res, nil
}
