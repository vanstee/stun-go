package stun_test

import (
  "fmt"
  . "stun"
)

func ExampleMessage() {
  header := Header{
    Method: "method",
    Class: "class",
    TransactionId: "transaction_id",
  }

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

  fmt.Printf("message.Header.Method: %s\n", message.Header.Method)
  fmt.Printf("message.Header.Class: %s\n", message.Header.Class)
  fmt.Printf("message.Header.TransactionId: %s\n", message.Header.TransactionId)

  for i := range attributes {
    fmt.Printf("message.Attributes[%d].Type: %s\n", i, message.Attributes[i].Type)
    fmt.Printf("message.Attributes[%d].Length: %d\n", i, message.Attributes[i].Length)
    fmt.Printf("message.Attributes[%d].Value: %s\n", i, message.Attributes[i].Value)
  }

  // Output:
  // message.Header.Method: method
  // message.Header.Class: class
  // message.Header.TransactionId: transaction_id
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
