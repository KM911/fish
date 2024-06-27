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
func HttpsProxy(port int) {
	tcpAddress, err := net.ResolveTCPAddr("tcp", "0.0.0.0:"+strconv.Itoa(port))
	format.Must(err)
	tcpListner, err := net.ListenTCP("tcp", tcpAddress)
	format.Must(err)
	for {
		conn, err := tcpListner.AcceptTCP()
		format.Must(err)
		go HandleTCPConn(conn)
	}
}

func HandleTCPConn(conn *net.TCPConn) {
	buffer := make([]byte, 4096)
	conn.Read(buffer)
	if string(buffer[:7]) == http.MethodConnect {
		// https
		conn.Write([]byte("HTTP/1.1 200 Connection Established\r\n\r\n"))
		index := 10
		for {
			if buffer[index] == 32 {
				break
			}
			index++
		}
		host := string(buffer[8:index])
		remote, err := net.Dial("tcp", host)
		if err != nil {
			log.Println(err)
			conn.Close()
			return
		}
		// transfer
		go transfer(conn, remote)
		go transfer(remote, conn)
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
