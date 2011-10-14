package main

import (
	"os"
)

type Server struct {
}

func Protocols(protocols [4][2]string) os.Error{
	protocols = [4][2]string{
		[2]string{"tcp", ":1234"},
		[2]string{"udp", ":1234"},
		[2]string{"tls", ":1235"},
		[2]string{"http", ":1236"},
	}
	return nil
}
