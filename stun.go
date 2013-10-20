package stun

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
  MagicCookie uint32
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
