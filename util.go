package stun

import (
  "encoding/binary"
  "math"
)

const (
  sizeOfUint32 = 4
)

func (attribute *Attribute) ChunkedValue() []uint32 {
  words := int(math.Ceil(float64(len(attribute.Value)) / float64(sizeOfUint32)))
  chunks := make([]uint32, words)
  bytes := []byte(attribute.Value)
  bytes = append(bytes, make([]byte, (len(chunks) * sizeOfUint32) - len(bytes))...)

  for i := range chunks {
    start := sizeOfUint32 * i
    end := sizeOfUint32 * (i + 1)
    chunks[i] = binary.LittleEndian.Uint32(bytes[start:end])
  }
  return chunks
}
