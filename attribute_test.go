package stun_test

import (
  "bytes"
  . "stun"
  "testing"
)

func TestNewAttribute(t *testing.T) {
  attribute := NewAttribute(0, "value")

  if attributeType := attribute.Type; attributeType != 0 {
    t.Errorf(`attribute.Type  = %d, want 0`, attributeType)
  }

  if length := attribute.Length; length != 5 {
    t.Errorf(`attribute.Length = %d, want 5`, length)
  }

  if value := attribute.Value; value != "value" {
    t.Errorf(`attribute.Value = %s, want "value"`, value)
  }
}

func TestSerializeAttribute(t *testing.T) {
  attribute := NewAttribute(0, "This is a value")
  buffer := attribute.Serialize()

  matching_buffer := []byte{
      0,   0,   0,  15, // Type, Length
    115, 105, 104,  84, // Value
     32, 115, 105,  32,
     97, 118,  32,  97,
      0, 101, 117, 108,
  }

  if !bytes.Equal(buffer, matching_buffer) {
    t.Errorf(`attribute.Serialize() = %v, want %v`, buffer, matching_buffer)
  }
}
