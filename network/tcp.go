package network

import (
	"net"

	"github.com/KM911/fish/format"
)

func GetTCPListener(url string) *net.TCPListener {
	TCPAddress, err := net.ResolveTCPAddr("tcp", url)
	format.Must(err)
	listener, err := net.ListenTCP("tcp", TCPAddress)
	format.Must(err)
	return listener
}

func EstablishConnection(url string) *net.TCPConn {
	TCPAddress, err := net.ResolveTCPAddr("tcp", url)
	format.Must(err)
	conn, err := net.DialTCP("tcp", nil, TCPAddress)
	format.Must(err)
	return conn
}
