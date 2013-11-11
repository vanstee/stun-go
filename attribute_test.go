package stun_test

import (
	"bytes"
	. "stun"
	"testing"
)

func TestNewAttribute(t *testing.T) {
	mappedAddress := NewMappedAddress(IPv4, 19302, 134744072)
	attribute := NewAttribute(1, mappedAddress)

	if attributeType := attribute.Type; attributeType != 1 {
		t.Errorf(`attribute.Type  = %d, want 1`, attributeType)
	}

	if length := attribute.Length; length != 5 {
		t.Errorf(`attribute.Length = %d, want 5`, length)
	}

	if value := attribute.Value; value != mappedAddress {
		t.Errorf(`attribute.Value = %s, want %s`, value, mappedAddress.String())
	}
}

func TestSerializeAttribute(t *testing.T) {
	mappedAddress := NewMappedAddress(IPv4, 19302, 134744072)
	attribute := NewAttribute(1, mappedAddress)
	buffer := attribute.Serialize()

	matchingBuffer := []byte{
		0, 1, 0, 5, // Type, Length
		1, 75, 102, 8, // Value
		8, 8, 8,
	}

	if !bytes.Equal(buffer, matchingBuffer) {
		t.Errorf(`attribute.Serialize() = %v, want %v`, buffer, matchingBuffer)
	}
}
