
# Compiler and flags
CC = gcc
CFLAGS = `pkg-config --cflags --libs gtk4`
SRC = gtk_gui/app.c
TARGET = gtk_gui/app
CXXFLAGS := -I/home/msandova/repos/go-llama.cpp/llama.cpp/common -I/home/msandova/repos/go-llama.cpp/llama.cpp -L/home/msandova/repos/go-llama.cpp
LDFLAGS := -L/home/msandova/repos/go-llama.cpp/libbinding.a
LIBS := -lbinding

build:
	@echo $(CXXFLAGS)
	GOOS=linux CGO_CPPFLAGS="$(subst $(space), ,$(CXXFLAGS))" CGO_LDFLAGS="$(LDFLAGS) $(LIBS)" go build 

protogen-go:
	protoc --proto_path=backend/proto --go_out=backend/proto --go_opt=paths=source_relative --go-grpc_out=backend/proto --go-grpc_opt=paths=source_relative backend.proto

gtk-app:
	$(CC) $(SRC) -o $(TARGET) $(CFLAGS) $(LDFLAGS)

clean:
	rm -f $(TARGET)
