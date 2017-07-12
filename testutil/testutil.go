package testutil

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

// GetPort returns a free open port that is ready to use
func GetPort() int {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port
}

// NewTestingServer - testing server and teardown func
func NewTestingServer(port int, status int, body string) (*http.Server, func()) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		if body != "" {
			io.WriteString(w, body)
		}
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(handler),
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		log.Fatalf("could not start http stub server: %v", err)
	}

	stoppable := Handle(listener)
	go server.Serve(stoppable)

	teardown := func() {
		stoppable.Stop <- true
	}

	return &server, teardown
}
