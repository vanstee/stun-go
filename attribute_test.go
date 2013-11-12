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
		0x00, 0x01, 0x00, 0x05, // Type, Length
		0x01, 0x4B, 0x66, 0x08, // Value
		0x08, 0x08, 0x08,
	}

	if !bytes.Equal(buffer, matchingBuffer) {
		t.Errorf(`attribute.Serialize() = %v, want %v`, buffer, matchingBuffer)
	}
}
