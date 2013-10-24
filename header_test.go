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

  matching_buffer := []byte{
      0,   1,   0,   0, // Class | Method, Length
     33,  18, 164,  66, // MagicCookie
      0,   0,   0,   1, // TransactionId
      0,   0,   0,   2,
      0,   0,   0,   3,
  }

  if !bytes.Equal(buffer, matching_buffer) {
    t.Errorf(`header.Serialize() = %X, want %X`, buffer, matching_buffer)
  }
}
