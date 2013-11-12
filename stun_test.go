package stun

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"testing"
)

func TestRequestPublicAddress(t *testing.T) {
	ipAddress, err := RequestPublicIPAddress()
	if err != nil {
		t.Errorf(err.Error())
	}

	matchingIPAddress, err := publicIPAddress()
	if err != nil {
		t.Errorf(err.Error())
	}

	if !ipAddress.Equal(matchingIPAddress) {
		t.Errorf(`ipAddress = "%s", want "%s"`, ipAddress, matchingIPAddress)
	}
}

func TestRequest(t *testing.T) {
	responseMessage, err := Request()
	if err != nil {
		t.Errorf(err.Error())
	}

	attributeValue := responseMessage.Attributes[0].Value
	mappedAddress, ok := attributeValue.(*MappedAddress)
	if !ok {
		t.Errorf(`Attribute was expected to be of type MappedAddress`)
	}

	ipAddress := mappedAddress.IPAddress()

	matchingIPAddress, err := publicIPAddress()
	if err != nil {
		t.Errorf(err.Error())
	}

	if !ipAddress.Equal(matchingIPAddress) {
		t.Errorf(`ipAddress = "%s", want "%s"`, ipAddress, matchingIPAddress)
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

func publicIPAddress() (net.IP, error) {
	response, err := http.Get("http://icanhazip.com")
	if err != nil {
		return nil, err
	}

	address, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}

	cleanAddress := strings.TrimSpace(string(address))
	ip := net.ParseIP(cleanAddress)
	return ip, nil
}

func fakeTransactionId() [3]uint32 {
	return [3]uint32{1, 2, 3}
}
