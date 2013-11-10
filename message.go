package stun

import (
  "bytes"
  "fmt"
)

type Message struct {
  Header *Header
  Attributes []*Attribute
}

func (message *Message) Serialize() []byte {
  bytes := []byte{}
  bytes = append(bytes, message.Header.Serialize()...)

  for _, attribute := range message.Attributes {
    bytes = append(bytes, attribute.Serialize()...)
  }

  return bytes
}

func ParseMessage(rawMessage []byte) (*Message, error) {
  header, err := ParseHeader(rawMessage[0:20])
  if (err != nil) { return nil, err }

  attributes, err := ParseAttributes(rawMessage[20:])
  if (err != nil) { return nil, err }

  message := &Message{
    Header: header,
    Attributes: attributes,
  }

  return message, nil
}

func (message *Message) String() string {
  var buffer bytes.Buffer
  buffer.WriteString("Message:\n")
  buffer.WriteString("Header:\n")
  buffer.WriteString(message.Header.String())

  for i, attribute := range message.Attributes {
    buffer.WriteString(fmt.Sprintf("Attributes[%d]:\n", i))
    buffer.WriteString(attribute.String())
  }

  return buffer.String()
}
