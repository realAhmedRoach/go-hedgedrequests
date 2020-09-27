package hedgedrequests_test

import (
	. "hedgedrequests"
	"testing"
)

// Test if the result of a simple function
// is returned
func TestHedgedRequest_Simple(t *testing.T) {
	req := func() []byte {
		return []byte("test")
	}

	res := HedgedRequest(req, 100, 1)

	if string(res) != "test" {
		t.Fatal("simple request failed")
	}
}
