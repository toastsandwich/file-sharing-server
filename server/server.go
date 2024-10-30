package server

import (
	"io/fs"

	"github.com/toastsandwich/fileSharingSystem/connection"
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



example:
	fileshare init -s -r
	this will connect machine to server with specific flags as permission


*/

type FileServer struct {
	addr string
	fs   *fs.FS

	ConnPool map[*connection.FileConn]struct{} // ConnPool for active users
}

func NewFileServer(addr string) *FileServer {
	// logic for file system
	var fs *fs.FS
	// pending logic

	
	return &FileServer{
		addr:     addr,
		fs:       fs,
		ConnPool: make(map[*connection.FileConn]struct{}),
	}
}
