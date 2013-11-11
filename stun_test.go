package stun_test

import (
	"fmt"
	. "stun"
)

func ExampleRequest() {
	header := NewHeader(RequestClass)

	message := &Message{
		Header:     header,
		Attributes: []*Attribute{},
	}

	responseMessage, err := Request(message)

	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf(responseMessage.String())
	}

	// Output:
	// Message:
	// Header:
	// Class: 0
	// Method: 1
	// Length: 0
	// MagicCookie: 554869826
	// TransactionId: [1 2 3]
	// Attributes[0]:
	// Type: 1
	// Length: 5
	// Value:
	// Family: 1
	// Port: 19302
	// Address: 134744072
}

func ExampleMessage() {
	header := NewHeader(RequestClass)
	header.TransactionId = fakeTransactionId()

	mappedAddress := NewMappedAddress(IPv4, 19302, 134744072)
	attribute := NewAttribute(1, mappedAddress)

	attributes := make([]*Attribute, 1)
	attributes[0] = attribute

	message := &Message{
		Header:     header,
		Attributes: attributes,
	}

	fmt.Printf(message.String())

	// Output:
	// Message:
	// Header:
	// Class: 0
	// Method: 1
	// Length: 0
	// MagicCookie: 554869826
	// TransactionId: [1 2 3]
	// Attributes[0]:
	// Type: 1
	// Length: 5
	// Value:
	// Family: 1
	// Port: 19302
	// Address: 134744072
}

func fakeTransactionId() [3]uint32 {
	return [3]uint32{1, 2, 3}
}
