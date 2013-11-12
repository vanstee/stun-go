package stun_test

import (
	"bytes"
	. "stun"
	"testing"
)

func TestNewHeader(t *testing.T) {
	header := NewHeader(RequestClass)

	if class := header.Class; class != RequestClass {
		t.Errorf(`header.Class = %d, want 0`, class)
	}

	if method := header.Method; method != BindingMethod {
		t.Errorf(`header.Method = %d, want 1`, method)
	}

	if magicCookie := header.MagicCookie; magicCookie != MagicCookie {
		t.Errorf(`header.magicCookie = %d, want 554869826`, magicCookie)
	}

	if transactionIdLength := len(header.TransactionId); transactionIdLength != 3 {
		t.Errorf(`len(header.transactionId) = %d, want 3`, transactionIdLength)
	}
}

func TestSerializeHeader(t *testing.T) {
	header := NewHeader(RequestClass)
	header.TransactionId = fakeTransactionId()
	buffer := header.Serialize()

	matchingBuffer := []byte{
		0x00, 0x01, 0x00, 0x00, // Class | Method, Length
		0x21, 0x12, 0xA4, 0x42, // MagicCookie
		0x00, 0x00, 0x00, 0x01, // TransactionId
		0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x03,
	}

	if !bytes.Equal(buffer, matchingBuffer) {
		t.Errorf(`header.Serialize() = %X, want %X`, buffer, matchingBuffer)
	}
}
