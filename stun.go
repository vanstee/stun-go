package stun

type Header struct {
  Method string
  Class string
  TransactionId string
}

type Attribute struct {
  Type string
  Length int
  Value string
}

type Message struct {
  Header Header
  Attributes []Attribute
}
