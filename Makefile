
# Compiler and flags
CC = gcc
CFLAGS = `pkg-config --cflags --libs gtk4`
SRC = gtk_gui/app.c
TARGET = gtk_gui/app

build:
	go build 

protogen-go:
	protoc --proto_path=backend/proto --go_out=backend/proto --go_opt=paths=source_relative --go-grpc_out=backend/proto --go-grpc_opt=paths=source_relative backend.proto

gtk-app:
	$(CC) $(SRC) -o $(TARGET) $(CFLAGS) $(LDFLAGS)

clean:
	rm -f $(TARGET)
