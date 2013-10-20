package stun

import (
  "math/rand"
)

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
  TransactionId []uint32
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

func NewHeader(class uint16) Header {
  return Header{
    Class: class,
    Method: BindingMethod,
    MagicCookie: MagicCookie,
    TransactionId: generateTransactionId(),
  }
}

func generateTransactionId() []uint32 {
  return []uint32{
    rand.Uint32(),
    rand.Uint32(),
    rand.Uint32(),
  }
}
