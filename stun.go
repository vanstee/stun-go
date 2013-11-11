package stun

import (
	"net"
	"time"
)

const (
	GoogleStunServer           = "stun.l.google.com:19302"
	MaxResponseLength          = 548
	RequestTimeoutMilliseconds = 500
)

func Request(request *Message) (*Message, error) {
	connection, err := net.DialTimeout("udp", GoogleStunServer, RequestTimeout())
	if err != nil {
		return nil, err
	}

	defer connection.Close()

	_, err = connection.Write(request.Serialize())
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, MaxResponseLength)
	_, err = connection.Read(buffer)
	if err != nil {
		return nil, err
	}

	return ParseMessage(buffer)
}

func RequestTimeout() time.Duration {
	return time.Duration(RequestTimeoutMilliseconds) * time.Millisecond
}
