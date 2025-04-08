package uuid

import (
	"testing"
)

func TestGetMac(t *testing.T) {

	mac := GetMac()
	length := len(mac)

	if length != 6 {
		t.Errorf("MAC address should be 6 bytes long, %d given", length)
	}
}
