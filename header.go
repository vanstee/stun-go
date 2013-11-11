package stun

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math/rand"
)

const (
	ProtocolMask = 0xC000
	ClassMask    = 0x0110
	MethodMask   = 0x3EEF

	RequestClass         = 0x0000
	IndicationClass      = 0x0010
	SuccessResponseClass = 0x0100
	FailureResponseClass = 0x0110

	BindingMethod = 0x0001

	MagicCookie = 0x2112A442
)

type Header struct {
	Class         uint16
	Method        uint16
	Length        uint16
	MagicCookie   uint32
	TransactionId [3]uint32
}

func NewHeader(class uint16) *Header {
	return &Header{
		Class:         class,
		Method:        BindingMethod,
		Length:        0,
		MagicCookie:   MagicCookie,
		TransactionId: generateTransactionId(),
	}
}

func (header *Header) Serialize() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.BigEndian, header.Type())
	binary.Write(buffer, binary.BigEndian, header.Length)
	binary.Write(buffer, binary.BigEndian, header.MagicCookie)
	binary.Write(buffer, binary.BigEndian, header.TransactionId)
	return buffer.Bytes()
}

func ParseHeader(rawHeader []byte) (*Header, error) {
	buffer := bytes.NewBuffer(rawHeader)
	header := &Header{}

	var headerType uint16
	binary.Read(buffer, binary.BigEndian, &headerType)
	if headerType&ProtocolMask != 0x0000 {
		return nil, errors.New("Protocol is invalid")
	}
	header.SetType(headerType)

	binary.Read(buffer, binary.BigEndian, &header.Length)

	binary.Read(buffer, binary.BigEndian, &header.MagicCookie)
	if header.MagicCookie != MagicCookie {
		return nil, errors.New("MagicCookie is invalid")
	}

	binary.Read(buffer, binary.BigEndian, &header.TransactionId)

	return header, nil
}

func (header *Header) SetType(headerType uint16) {
	header.Class = headerType & ClassMask
	header.Method = headerType & MethodMask
}

func (header *Header) Type() uint16 {
	return header.Class | header.Method
}

func (header *Header) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Class: %d\n", header.Class))
	buffer.WriteString(fmt.Sprintf("Method: %d\n", header.Method))
	buffer.WriteString(fmt.Sprintf("Length: %d\n", header.Length))
	buffer.WriteString(fmt.Sprintf("MagicCookie: %d\n", header.MagicCookie))
	buffer.WriteString(fmt.Sprintf("TransactionId: %d\n", header.TransactionId))
	return buffer.String()
}

func generateTransactionId() [3]uint32 {
	return [3]uint32{
		rand.Uint32(),
		rand.Uint32(),
		rand.Uint32(),
	}
}
