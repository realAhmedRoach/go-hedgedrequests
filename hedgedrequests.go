// This package provides functionality for hedged requests,
// which are used to reduce latency in distributed systems.
//
// "One such approach isto defer sending a secondary
// request untilthe first request has been outstanding
// for more than the 95th-percentile expected latency
// for this class of requests. This approach limits the
// additional load to approximately 5% while
// substantially shortening the latency tail." - 2013 Google Paper
package hedgedrequests

import "time"

func HedgedRequest(req func() []byte, tailLatency int, maxQueries int) []byte {
	ch := make(chan []byte, maxQueries)
	var queriesDone = 1            // Count how many queries have been sent
	go func() { ch <- req() }()    // Make first request
	for queriesDone < maxQueries { // Stop the loop once maxQueries is hit
		select {
		case res := <-ch: // A result is available
			return res
		case <-time.After(time.Millisecond * time.Duration(tailLatency)): // Defer sending the subsequent requests
			queriesDone++
			go func() { ch <- req() }() // Make subsequent requests
		}
	}
	return <-ch // Blocks until a result is available
}
