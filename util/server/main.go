package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"github.com/njdaniel/dnd/util/list"
	grpc "google.golang.org/grpc"
)

// run grpc server that interacts with fileserver

type fileServer struct{}

func main() {
	fmt.Println("in main")
	srv := grpc.NewServer()
	var files fileServer
	list.RegisterFilesServer(srv, files)
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("could not listen to port :8888: %v", err)
	}
	log.Fatal(srv.Serve(l))
}

func (fileServer) List(ctx context.Context, path *list.Path) (fileList *list.FileList, err error) {
	// return list of directories/files directly under
	fmt.Println("Func list called")
	var files list.FileList

	dirName := fmt.Sprintf("./data/default-house-example%v", path)
	fs, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	for _, f := range fs {
		//fmt.Println(f.Name())
		var file list.File
		file.Text = f.Name()
		files.Files = append(files.Files, &file)
	}
	return &files, nil

	//items in json
	//loop through jsons in /items directory

	//open json file
	//jsonFile, err := os.Open("../data/pf2playtest/weapons.json")
	//if err != nil {
	//	//TODO: return err? use RunE instead of Run?
	//	fmt.Println(err)
	//}
	//defer jsonFile.Close()

	//parse into struct

	//print out each name}
}
