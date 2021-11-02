package skio

import (
	"net"
	"net/http"
	"net/url"
)

type Conn interface {
	ID() string
	Close() error
	URL() url.URL
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	RemoteHeader() http.Header

	Context() interface{}
	SetContext(v interface{})
	Namespace() string
	Emit(msg string,v ...interface{})

	Join(room string)
	Leave(room string)
	LeaveAll()
	Rooms() []string
}

type AppSocket interface {
	Conn
}
type appSocket struct {
	Conn
}
func NewAppSocket(conn Conn) *appSocket{
	return  &appSocket{conn}
}