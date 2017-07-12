package testutil

import "testing"

func TestGetPort(t *testing.T) {
	p := GetPort()
	if p == 0 {
		t.Errorf("port should not be 0")
	}
}
