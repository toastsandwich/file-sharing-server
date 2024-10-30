package connection

import "net"

// FileConn is the struct having ability to share and recieve files.
// net.Conn is the connection.
// Permissions are multiplexer for int. will be always a unique code.
// recv = rune(r)
// send = rune(s)
// r = 114
// s = 115

const (
	PermissionSend        = 'S'
	PermissionRecieve     = 'R'
	PermissionSendRecieve = 'B'
)

type FileConn struct {
	net.Conn // this will be the be connection
	rune     // permissions where to just share, just recieve or both
}

// recieve a pointer connection which will be added to conn pool
// see conn pool doc for more details...
func NewFileConn(conn net.Conn) (*FileConn, error) {
	var perm rune
	buf := make([]byte, 1)
	_, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}

	perm = rune(buf[0])
	if perm != PermissionRecieve && perm != PermissionSend && perm != PermissionSendRecieve {
		return nil, err
	}

	return &FileConn{
		Conn: conn,
		rune: perm,
	}, nil
}

// For getting the permission for
// a specific connection
func (f *FileConn) Perm() rune {
	return f.rune
}
