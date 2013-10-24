package stun_test

import (
  "bytes"
  "fmt"
  "reflect"
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
      0,   1,   0,   0, // Class | Method, Length
     33,  18, 164,  66, // MagicCookie
      0,   0,   0,   1, // TransactionId
      0,   0,   0,   2,
      0,   0,   0,   3,
  }

  if !bytes.Equal(buffer, matching_buffer) {
    t.Errorf(`header.Serialize() = %X, want %X`, buffer, matching_buffer)
  }
}

func TestNewAttribute(t *testing.T) {
  attribute := NewAttribute(0, "value")

  if attributeType := attribute.Type; attributeType != 0 {
    t.Errorf(`attribute.Type  = %d, want 0`, attributeType)
  }

  if length := attribute.Length; length != 5 {
    t.Errorf(`attribute.Length = %d, want 5`, length)
  }

  if value := attribute.Value; value != "value" {
    t.Errorf(`attribute.Value = %s, want "value"`, value)
  }
}

func TestChunkedValue(t *testing.T) {
  attribute := NewAttribute(0, "This is a chunked value")
  chunks := attribute.ChunkedValue()

  matching_chunks := []uint32{
    0x73696854, 0x20736920, 0x68632061, 0x656B6E75, 0x61762064, 0x0065756C,
  }

  if !reflect.DeepEqual(chunks, matching_chunks) {
    t.Errorf(`attribute.ChunkedValue() = %X, want %X`, chunks, matching_chunks)
  }
}

func TestSerializeAttribute(t *testing.T) {
  attribute := NewAttribute(0, "This is a value")
  buffer := attribute.Serialize()

  matching_buffer := []byte{
      0,   0,   0,  15, // Type, Length
    115, 105, 104,  84, // Value
     32, 115, 105,  32,
     97, 118,  32,  97,
      0, 101, 117, 108,
  }

  if !bytes.Equal(buffer, matching_buffer) {
    t.Errorf(`attribute.Serialize() = %v, want %v`, buffer, matching_buffer)
  }
}

func ExampleMessage() {
  header := NewHeader(RequestClass)
  header.TransactionId = fakeTransactionId()

  attributes := make([]*Attribute, 3)

  for i := range attributes {
    attributes[i] = &Attribute{
      Type: 0,
      Length: 0,
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
    fmt.Printf("message.Attributes[%d].Type: %d\n", i, message.Attributes[i].Type)
    fmt.Printf("message.Attributes[%d].Length: %d\n", i, message.Attributes[i].Length)
    fmt.Printf("message.Attributes[%d].Value: %s\n", i, message.Attributes[i].Value)
  }

  // Output:
  // message.Header.Method: 1
  // message.Header.Class: 0
  // message.Header.MagicCookie: 554869826
  // message.Header.TransactionId: [1 2 3]
  // message.Attributes[0].Type: 0
  // message.Attributes[0].Length: 0
  // message.Attributes[0].Value: value
  // message.Attributes[1].Type: 0
  // message.Attributes[1].Length: 0
  // message.Attributes[1].Value: value
  // message.Attributes[2].Type: 0
  // message.Attributes[2].Length: 0
  // message.Attributes[2].Value: value
}

func fakeTransactionId() [3]uint32 {
  return [3]uint32{1, 2, 3}
}
