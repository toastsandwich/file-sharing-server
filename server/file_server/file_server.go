package file_server

import (
	"fmt"
	"net"
	"sync"

	"github.com/toastsandwich/fileSharingSystem/server/connection"
	idgenerator "github.com/toastsandwich/fileSharingSystem/server/idGenerator"
)

/*
how to allow file share

server will create a map, with permissions
somewhat like this

map ::
	key						value
	ip1						recv, send
	ip2 					recv
	ip3						send

scratch this map

new map ::
	conn struct{}


example:
	fileshare init -s -r
	this will connect machine to server with specific flags as permission


*/

const maxconnections = 10

type FileServer struct {
	iD      string
	addr    string
	counter int // counter to have server limited connections
	mu      sync.Mutex

	ConnPool map[*connection.FileConn]struct{} // ConnPool for active users
	ErrorCh  chan error                        // For better error handelling
}

// create a new FileServer
func NewFileServer(addr string) *FileServer {
	return &FileServer{
		iD:      idgenerator.GenerateID("file-server"),
		addr:    addr,
		counter: 0,

		ConnPool: make(map[*connection.FileConn]struct{}),
		ErrorCh:  make(chan error),
	}
}

func (f *FileServer) Start() {
	// listener for accepting incoming connections
	ln, err := net.Listen("tcp", f.addr)
	if err != nil {
		f.ErrorCh <- err
		return
	}
	defer ln.Close() // close the listner

	// wait for connections
	for {
		// check if connection limit is reached.
		if f.counter >= maxconnections {
			f.mu.Unlock()
			continue // skip connection
		}

		// accept the connection
		conn, err := ln.Accept()
		if err != nil {
			f.ErrorCh <- err
			continue
		}

		// get the file connection.
		fc, err := connection.NewFileConn(conn)
		if err != nil {
			f.ErrorCh <- err
			continue
		}
		// start a seperate go rountine for each connetion
		go f.handleConnection(fc)
	}
}

func (f *FileServer) handleConnection(fc *connection.FileConn) {
	defer fc.Close()

	perm := fc.Perm()
	buf := []byte(string(perm))
	_, err := fc.Write(buf)
	if err != nil {
		f.ErrorCh <- err
	}
}

func (f *FileServer) Info() {
	info := fmt.Sprintf("server id: %s\nserver addr: %s", f.iD, f.addr)
	fmt.Println(info)
}
