package network

import (
	"io"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/KM911/fish/format"
)

// will create a https proxy
func HttpsProxy(_port int) {
	tcpAddress, err := net.ResolveTCPAddr("tcp", "0.0.0.0:"+strconv.Itoa(_port))
	format.Must(err)
	tcpListner, err := net.ListenTCP("tcp", tcpAddress)
	format.Must(err)
	for {
		conn, err := tcpListner.AcceptTCP()
		format.Must(err)
		go HandleTCPConn(conn)
	}
}

func HandleTCPConn(_conn *net.TCPConn) {
	buffer := make([]byte, 4096)
	_conn.Read(buffer)
	if string(buffer[:7]) == http.MethodConnect {
		// https
		_conn.Write([]byte("HTTP/1.1 200 Connection Established\r\n\r\n"))
		index := 10
		for {
			if buffer[index] == 32 {
				break
			}
			index++
		}
		host := string(buffer[8:index])
		// conn remote
		// fmt.Println("host is ", host)
		remote, err := net.Dial("tcp", host)
		if err != nil {
			log.Println(err)
			_conn.Close()
			return
		}
		// transfer
		go transfer(_conn, remote)
		go transfer(remote, _conn)
	} else {
		// deny http
		return
	}
}
func transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
}
