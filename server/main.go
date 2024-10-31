package main

import (
	"fmt"
	"log"

	fileserver "github.com/toastsandwich/fileSharingSystem/server/file_server"
)

func main() {
	fserv := fileserver.NewFileServer(":15001")
	fserv.Info()

	fmt.Println("starting server in different go routine")
	go fserv.Start()

	for err := range fserv.ErrorCh {
		log.Panic(err)
	}
}
