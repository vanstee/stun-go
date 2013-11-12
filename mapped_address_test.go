package stun

import (
	"net"
	"testing"
)

func TestNewMappedAddress(t *testing.T) {
	mappedAddress := NewMappedAddress(IPv4, 19302, 134744072)

	if mappedAddress.Family != IPv4 {
		t.Errorf(`mappedAddress.Family = %d, want 1`, mappedAddress.Family)
	}

	if mappedAddress.Port != 19302 {
		t.Errorf(`mappedAddress.Port = %d, want 19302`, mappedAddress.Port)
	}

	if mappedAddress.Address != 134744072 {
		t.Errorf(`mappedAddress.Address = %d, want 134744072`, mappedAddress.Address)
	}
}

func TestMappedAddressIPAddress(t *testing.T) {
	mappedAddress := NewMappedAddress(IPv4, 19302, 134744072)

	ipAddress := mappedAddress.IPAddress()
	matchingIPAddress := net.ParseIP("8.8.8.8")

	if !ipAddress.Equal(matchingIPAddress) {
		t.Errorf(`ipAddress = "%s", want "%s"`, ipAddress, matchingIPAddress)
	}
}
