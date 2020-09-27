package hedgedrequests_test

import (
	. "hedgedrequests"
	"testing"
	"time"
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

// Test if a new request is queries when its
// latency is over the tail latency.
func TestHedgedRequest_WithLatency(t *testing.T) {
	tailLatency := 20
	reqsCalled := 0

	req := func() []byte {
		reqsCalled++
		if reqsCalled == 1 {
			select {
			case <-time.After(time.Duration(tailLatency+1) * time.Millisecond): // Trigger subsequent request
				return []byte("test")
			}
		} else if reqsCalled == 2 { // Second request
			return []byte("test") // Return immediately
		}
		return []byte("fail, return called outside of if")
	}

	res := HedgedRequest(req, tailLatency, 2)

	if string(res) != "test" {
		t.Fatal("latency test failed")
	}
	if reqsCalled != 2 {
		t.Fatal(reqsCalled, "reqs were called instead of 2")
	}
}
