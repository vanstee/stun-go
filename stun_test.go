package stun_test

import (
  "fmt"
  . "stun"
)

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
