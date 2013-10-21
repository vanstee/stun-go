package stun

import (
  "bytes"
  "encoding/binary"
  "math"
  "math/rand"
)

const (
  RequestClass = 0x0000
  IndicationClass = 0x0010
  SuccessResponseClass = 0x0100
  FailureResponseClass = 0x0110

  BindingMethod = 0x0001

  MagicCookie = 0x2112A442
)

type Header struct {
  Class uint16
  Method uint16
  Length uint16
  MagicCookie uint32
  TransactionId [3]uint32
}

type Attribute struct {
  Type uint16
  Length uint16
  Value string
}

type Message struct {
  Header Header
  Attributes []Attribute
}

func NewHeader(class uint16) Header {
  return Header{
    Class: class,
    Method: BindingMethod,
    Length: 0,
    MagicCookie: MagicCookie,
    TransactionId: generateTransactionId(),
  }
}

func (header Header) Type() uint16 {
  return header.Class | header.Method
}

func generateTransactionId() [3]uint32 {
  return [3]uint32{
    rand.Uint32(),
    rand.Uint32(),
    rand.Uint32(),
  }
}

func (header Header) Serialize() []byte {
  buffer := new(bytes.Buffer)
  binary.Write(buffer, binary.BigEndian, header.Type())
  binary.Write(buffer, binary.BigEndian, header.Length)
  binary.Write(buffer, binary.BigEndian, header.MagicCookie)
  binary.Write(buffer, binary.BigEndian, header.TransactionId)
  return buffer.Bytes()
}

func NewAttribute(attributeType uint16, value string) Attribute {
  return Attribute{
    Type: attributeType,
    Length: uint16(len(value)),
    Value: value,
  }
}

func (attribute Attribute) ChunkedValue() []uint32 {
  sizeOfUint32 := 4
  words := int(math.Ceil(float64(len(attribute.Value)) / float64(sizeOfUint32)))
  chunks := make([]uint32, words)
  bytes := []byte(attribute.Value)
  bytes = append(bytes, make([]byte, (len(chunks) * sizeOfUint32) - len(bytes))...)

  for i := range chunks {
    start := sizeOfUint32 * i
    end := sizeOfUint32 * (i + 1)
    chunks[i] = binary.LittleEndian.Uint32(bytes[start:end])
  }
  return chunks
}

func (attribute Attribute) Serialize() []byte {
  buffer := new(bytes.Buffer)
  binary.Write(buffer, binary.BigEndian, attribute.Type)
  binary.Write(buffer, binary.BigEndian, attribute.Length)
  binary.Write(buffer, binary.BigEndian, attribute.ChunkedValue())
  return buffer.Bytes()
}
