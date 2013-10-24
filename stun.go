package stun

import (
  "bytes"
  "encoding/binary"
  "errors"
  "math"
  "math/rand"
  "net"
)

const (
  ProtocolMask = 0xC000
  ClassMask = 0x0110
  MethodMask = 0x3EEF

  RequestClass = 0x0000
  IndicationClass = 0x0010
  SuccessResponseClass = 0x0100
  FailureResponseClass = 0x0110

  BindingMethod = 0x0001

  MagicCookie = 0x2112A442

  GoogleStunServer = "stun.l.google.com:19302"
  MaxResponseLength = 548
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
  Header *Header
  Attributes []*Attribute
}

func NewHeader(class uint16) *Header {
  return &Header{
    Class: class,
    Method: BindingMethod,
    Length: 0,
    MagicCookie: MagicCookie,
    TransactionId: generateTransactionId(),
  }
}

func (header *Header) SetType(headerType uint16) {
  header.Class = headerType & ClassMask
  header.Method = headerType & MethodMask
}

func (header *Header) Type() uint16 {
  return header.Class | header.Method
}

func generateTransactionId() [3]uint32 {
  return [3]uint32{
    rand.Uint32(),
    rand.Uint32(),
    rand.Uint32(),
  }
}

func (header *Header) Serialize() []byte {
  buffer := new(bytes.Buffer)
  binary.Write(buffer, binary.BigEndian, header.Type())
  binary.Write(buffer, binary.BigEndian, header.Length)
  binary.Write(buffer, binary.BigEndian, header.MagicCookie)
  binary.Write(buffer, binary.BigEndian, header.TransactionId)
  return buffer.Bytes()
}

func ParseHeader(rawHeader []byte) (*Header, error) {
  buffer := bytes.NewBuffer(rawHeader)
  header := &Header{}

  var headerType uint16
  binary.Read(buffer, binary.BigEndian, &headerType)
  if headerType & ProtocolMask != 0x0000 {
    return nil, errors.New("Protocol is invalid")
  }
  header.SetType(headerType)

  binary.Read(buffer, binary.BigEndian, &header.Length)

  binary.Read(buffer, binary.BigEndian, &header.MagicCookie)
  if header.MagicCookie != MagicCookie {
    return nil, errors.New("MagicCookie is invalid")
  }

  binary.Read(buffer, binary.BigEndian, &header.TransactionId)

  return header, nil
}

func NewAttribute(attributeType uint16, value string) *Attribute {
  return &Attribute{
    Type: attributeType,
    Length: uint16(len(value)),
    Value: value,
  }
}

func (attribute *Attribute) ChunkedValue() []uint32 {
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

func (attribute *Attribute) Serialize() []byte {
  buffer := new(bytes.Buffer)
  binary.Write(buffer, binary.BigEndian, attribute.Type)
  binary.Write(buffer, binary.BigEndian, attribute.Length)
  binary.Write(buffer, binary.BigEndian, attribute.ChunkedValue())
  return buffer.Bytes()
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

  attributes, err := ParseAttributes(rawMessage[20:], header.Length)
  if (err != nil) { return nil, err }

  message := &Message{
    Header: header,
    Attributes: attributes,
  }

  return message, nil
}

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
