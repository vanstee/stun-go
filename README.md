STUN implementation in Go
=========================

```go
message := &Message{
  Header: NewHeader(RequestClass),
  Attributes: []*Attribute{},
}

responseMessage, err := Request(message)
if err != nil {
  fmt.Fatal(err.Error())
}

attributeValue := responseMessage.Attributes[0].Value
mappedAddress, ok := attributeValue.(*MappedAddress)
if !ok {
  fmt.Fatal("Expected Attribute of type MappedAddress")
}

address := make([]byte, 4)
binary.BigEndian.PutUint32(address, mappedAddress.Address)
formattedAddress := net.IPv4(address[0], address[1], address[2], address[3]).String()

fmt.Printf("Your public IP address is %s\n", formattedAddress)

// Output:
// Your public IP address is 8.8.8.8
```
