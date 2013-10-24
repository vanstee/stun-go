package stun_test

import (
  "reflect"
  . "stun"
  "testing"
)

func TestChunkedValue(t *testing.T) {
  attribute := NewAttribute(0, "This is a chunked value")
  chunks := attribute.ChunkedValue()

  matching_chunks := []uint32{
    0x73696854, 0x20736920, 0x68632061, 0x656B6E75, 0x61762064, 0x0065756C,
  }

  if !reflect.DeepEqual(chunks, matching_chunks) {
    t.Errorf(`attribute.ChunkedValue() = %X, want %X`, chunks, matching_chunks)
  }
}
