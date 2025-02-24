build:
	go build 

protogen-go:
	protoc --proto_path=backend/proto --go_out=backend/proto --go_opt=paths=source_relative --go-grpc_out=backend/proto --go-grpc_opt=paths=source_relative backend.proto
