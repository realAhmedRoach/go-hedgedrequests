# Go Hedged Requests
A simple library to execute hedged requests in Go

## Usage
```go
import . "hedgedrequests"

yourRequest := func() []byte {
// Do something...
}

tailLatency := 100 // Milliseconds
maxQueries := 3 // Max retries

result := HedgedRequest(yourRequest, tailLatency, maxQueries)
```

## Acknowledgements
[The Tail At Scale](https://research.google/pubs/pub40801/)