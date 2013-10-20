package stun_test

import (
  "bytes"
  "fmt"
  . "stun"
  "testing"
)

func TestNewHeader(t *testing.T) {
  header := NewHeader(RequestClass)

  if class := header.Class; class != RequestClass {
    t.Errorf(`header.Class = %d, want 0`, class)
  }

  if method := header.Method; method != BindingMethod {
    t.Errorf(`header.Method = %d, want 1`, method)
  }

  if magicCookie := header.MagicCookie; magicCookie != MagicCookie {
    t.Errorf(`header.magicCookie = %d, want 554869826`, magicCookie)
  }

  if transactionIdLength := len(header.TransactionId); transactionIdLength != 3 {
    t.Errorf(`len(header.transactionId) = %d, want 3`, transactionIdLength)
  }
}

func TestSerializeHeader(t *testing.T) {
  header := NewHeader(RequestClass)
  header.TransactionId = fakeTransactionId()
  buffer := header.Serialize()

  matching_buffer := []byte{
      0,   1,   0,   0, // Class ^ Method, Length
     33,  18, 164,  66, // MagicCookie
      0,   0,   0,   1, // TransactionId
      0,   0,   0,   2,
      0,   0,   0,   3,
  }

  if !bytes.Equal(buffer, matching_buffer) {
    t.Errorf(`header.Serialize() = %X, want %X`, buffer, matching_buffer)
  }
}

func ExampleMessage() {
  header := NewHeader(RequestClass)
  header.TransactionId = fakeTransactionId()

  attributes := make([]Attribute, 3)

  for i := range attributes {
    attributes[i] = Attribute{
      Type: "type",
      Length: i,
      Value: "value",
    }
  }

  message := &Message{
    Header: header,
    Attributes: attributes,
  }

  fmt.Printf("message.Header.Method: %d\n", message.Header.Method)
  fmt.Printf("message.Header.Class: %d\n", message.Header.Class)
  fmt.Printf("message.Header.MagicCookie: %d\n", message.Header.MagicCookie)
  fmt.Printf("message.Header.TransactionId: %d\n", message.Header.TransactionId)

  for i := range attributes {
    fmt.Printf("message.Attributes[%d].Type: %s\n", i, message.Attributes[i].Type)
    fmt.Printf("message.Attributes[%d].Length: %d\n", i, message.Attributes[i].Length)
    fmt.Printf("message.Attributes[%d].Value: %s\n", i, message.Attributes[i].Value)
  }

  // Output:
  // message.Header.Method: 1
  // message.Header.Class: 0
  // message.Header.MagicCookie: 554869826
  // message.Header.TransactionId: [1 2 3]
  // message.Attributes[0].Type: type
  // message.Attributes[0].Length: 0
  // message.Attributes[0].Value: value
  // message.Attributes[1].Type: type
  // message.Attributes[1].Length: 1
  // message.Attributes[1].Value: value
  // message.Attributes[2].Type: type
  // message.Attributes[2].Length: 2
  // message.Attributes[2].Value: value
}

func fakeTransactionId() [3]uint32 {
  return [3]uint32{1, 2, 3}
}
