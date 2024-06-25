package network

import (
	"net"

	"github.com/KM911/fish/format"
)

func GetTCPListener(_url string) *net.TCPListener {
	TCPAddress, err := net.ResolveTCPAddr("tcp", _url)
	format.Must(err)
	listener, err := net.ListenTCP("tcp", TCPAddress)
	format.Must(err)
	return listener
}

func EstablishConnetion(_url string) *net.TCPConn {
	TCPAddress, err := net.ResolveTCPAddr("tcp", _url)
	format.Must(err)
	conn, err := net.DialTCP("tcp", nil, TCPAddress)
	format.Must(err)
	return conn
}
