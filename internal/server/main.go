package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/njdaniel/dnd/internal/list"
	grpc "google.golang.org/grpc"
)

// run grpc server that interacts with fileserver

type fileServer struct{}

func main() {
	//fmt.Println("in main")
	srv := grpc.NewServer()
	var files fileServer
	list.RegisterFilesServer(srv, files)
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("could not listen to port :8888: %v", err)
	}
	log.Fatal(srv.Serve(l))
}

func (fileServer) List(ctx context.Context, void *list.Path) (fileList *list.FileList, err error) {

	return nil, fmt.Errorf("List not setup yet %v", err)
}
