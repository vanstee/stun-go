package stun_test

import (
  "fmt"
  . "stun"
)

func ExampleMessage() {
  header := Header{
    Class: RequestClass,
    Method: BindingMethod,
    MagicCookie: MagicCookie,
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

  fmt.Printf("message.Header.Method: %d\n", message.Header.Method)
  fmt.Printf("message.Header.Class: %d\n", message.Header.Class)
  fmt.Printf("message.Header.TransactionId: %s\n", message.Header.TransactionId)

  for i := range attributes {
    fmt.Printf("message.Attributes[%d].Type: %s\n", i, message.Attributes[i].Type)
    fmt.Printf("message.Attributes[%d].Length: %d\n", i, message.Attributes[i].Length)
    fmt.Printf("message.Attributes[%d].Value: %s\n", i, message.Attributes[i].Value)
  }

  // Output:
  // message.Header.Method: 1
  // message.Header.Class: 0
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
