package main

import (
	"github.com/njdaniel/dnd/internal/list"
	"google.golang.org/grpc"
)

// run grpc server that interacts with fileserver

func main() {
	//fmt.Println("in main")
	srv := grpc.NewServer()
	list.RegisterFilesServer(srv, files)
}

type fileServer struct{}
