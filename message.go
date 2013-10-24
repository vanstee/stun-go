package stun

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

  attributes, err := ParseAttributes(rawMessage[20:], header.Length)
  if (err != nil) { return nil, err }

  message := &Message{
    Header: header,
    Attributes: attributes,
  }

  return message, nil
}
