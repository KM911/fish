package network

import (
	"fmt"
	"strings"
)

type Request struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
	Body    []byte
}

type Respond struct {
	Version string
	Code    int

	Headers map[string]string
	Body    []byte
}

func NewRespond() *Respond {
	return &Respond{
		Version: "HTTP/1.1",
		Code:    200,
		Headers: make(map[string]string),
		Body:    []byte{},
	}
}

func (r *Respond) GetString() string {
	sb := strings.Builder{}
	sb.WriteString(r.Version + " " + fmt.Sprintf("%d", r.Code) + " OK\r\n")
	r.Headers["Content-Length"] = fmt.Sprintf("%d", len(r.Body))
	for k, v := range r.Headers {

		sb.WriteString(k + ": " + v + "\r\n")
	}
	sb.WriteString("\r\n")
	sb.Write(r.Body)
	return sb.String()

}

func Parse(_data []byte) (r Request) {
	lines := strings.Split(string(_data), "\r\n")
	fmt.Println(lines[0])
	first_line_items := strings.Split(lines[0], " ")
	r.Method = first_line_items[0]
	r.Path = first_line_items[1]
	r.Version = first_line_items[2]

	r.Headers = make(map[string]string)
	for _, line := range lines[1:] {
		if line == "" {
			break
		}
		items := strings.Split(line, ": ")
		r.Headers[items[0]] = items[1]
	}
	// r.Body
	r.Body = []byte("")
	return r
}

func (r *Request) GetString() string {
	sb := strings.Builder{}
	sb.WriteString(r.Method + " " + r.Path + " " + r.Version + "\r\n")
	for k, v := range r.Headers {
		sb.WriteString(k + ": " + v + "\r\n")
	}
	sb.WriteString("\r\n")
	sb.Write(r.Body)
	return sb.String()
}
