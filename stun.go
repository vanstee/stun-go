package stun

import (
  "net"
)

const (
  GoogleStunServer = "stun.l.google.com:19302"
  MaxResponseLength = 548
)

func Request(request *Message) (*Message, error) {
  connection, err := net.Dial("udp", GoogleStunServer)
  if (err != nil) { return nil, err }

  defer connection.Close()

  _, err = connection.Write(request.Serialize())
  if (err != nil) { return nil, err }

  buffer := make([]byte, MaxResponseLength)
  _, err = connection.Read(buffer)
  if (err != nil) { return nil, err }

  response, err := ParseMessage(buffer)
  return response, err
}
