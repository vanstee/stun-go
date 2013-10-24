package stun

import (
  "bytes"
  "encoding/binary"
  "fmt"
  "net"
)

type Attribute struct {
  Type uint16
  Length uint16
  Value string
}

type MappedAddress struct {
  Family uint16
  Port uint16
  Address uint32
}

func NewAttribute(attributeType uint16, value string) *Attribute {
  return &Attribute{
    Type: attributeType,
    Length: uint16(len(value)),
    Value: value,
  }
}

func (attribute *Attribute) Serialize() []byte {
  buffer := new(bytes.Buffer)
  binary.Write(buffer, binary.BigEndian, attribute.Type)
  binary.Write(buffer, binary.BigEndian, attribute.Length)
  binary.Write(buffer, binary.BigEndian, attribute.ChunkedValue())
  return buffer.Bytes()
}

func ParseAttribute(rawAttribute []byte) (*Attribute, error) {
  return &Attribute{}, nil
}

func ParseAttributes(rawAttributes []byte, messageLength uint16) ([]*Attribute, error) {
  buffer := bytes.NewBuffer(rawAttributes)
  attributes := make([]*Attribute, 1)

  attribute := &Attribute{}
  binary.Read(buffer, binary.BigEndian, &attribute.Type)
  binary.Read(buffer, binary.BigEndian, &attribute.Length)

  attributes[0] = attribute

  return attributes, nil
}
