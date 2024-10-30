package connection

import "net"

type Permissions rune

// FileConn is the struct having ability to share and recieve files.
// net.Conn is the connection.
// Permissions are multiplexer for int. will be always a unique code.
// recv = rune(r)
// send = rune(s)
// r = 114
// s = 115
type FileConn struct {
	net.Conn    // this will be the be connection
	Permissions // permissions where to just share, just recieve or both
}

// recieve a pointer connection which will be added to conn pool
// see conn pool doc for more details...
func NewFileConn(conn net.Conn, recv, send bool) *FileConn {
	var perm rune
	if recv {
		perm += 'r'
	}
	if send {
		perm += 's'
	}
	return &FileConn{
		Conn:        conn,
		Permissions: Permissions(perm),
	}
}
