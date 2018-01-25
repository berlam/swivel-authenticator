package pkg

import "testing"

func TestCalculateOTC(t *testing.T) {
	if calculateOTC("9876509432", "0189") != "2943" {
		t.Fail()
	}
	if calculateOTC("3456782210", "0189") != "0321" {
		t.Fail()
	}
}
