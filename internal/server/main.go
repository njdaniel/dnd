package main

import (
	"log"
	"net"

	"github.com/njdaniel/dnd/internal/list"
	grpc "google.golang.org/grpc"
)

// run grpc server that interacts with fileserver

func main() {
	//fmt.Println("in main")
	srv := grpc.NewServer()
	list.RegisterFilesServer(srv, files)
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("could not listen to port :8888: %v", err)
	}
	log.Fatal(srv.Serve(l))
}

type fileServer struct{}
