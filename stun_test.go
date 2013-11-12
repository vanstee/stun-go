package stun

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"testing"
)

func TestRequest(t *testing.T) {
	header := NewHeader(RequestClass)

	message := &Message{
		Header:     header,
		Attributes: []*Attribute{},
	}

	responseMessage, err := Request(message)
	if err != nil {
		t.Errorf(err.Error())
	}

	attributeValue := responseMessage.Attributes[0].Value
	mappedAddress, ok := attributeValue.(*MappedAddress)

	if !ok {
		t.Errorf(`responseMessage.Attributes[0] was expected to be of type MappedAddress`)
	}

	address := make([]byte, 4)
	binary.BigEndian.PutUint32(address, mappedAddress.Address)

	formattedAddress := net.IPv4(address[0], address[1], address[2], address[3]).String()

	response, err := http.Get("http://icanhazip.com")
	if err != nil {
		t.Errorf(err.Error())
	}

	matchingAddress, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		t.Errorf(err.Error())
	}

	formattedMatchingAddress := strings.TrimSpace(string(matchingAddress))

	if formattedAddress != formattedMatchingAddress {
		t.Errorf(`formattedAddress = "%s", want "%s"`, formattedAddress, formattedMatchingAddress)
	}
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
