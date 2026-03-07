package metrics

import (
	"sync"
	"sync/atomic"
)

var activeConnections int64
var totalVisits int64

var endpointVisits = make(map[string]int64)
var mu sync.Mutex
var visitors = make(map[string]struct{})

func IncActive() {
	atomic.AddInt64(&activeConnections, 1)
}

func DecActive() {
	atomic.AddInt64(&activeConnections, -1)
}

func IncTotal() {
	atomic.AddInt64(&totalVisits, 1)
}

func IncEndpoint(endpoint string) {
	mu.Lock()
	endpointVisits[endpoint]++
	mu.Unlock()
}

func Active() int64 {
	return atomic.LoadInt64(&activeConnections)
}

func Total() int64 {
	return atomic.LoadInt64(&totalVisits)
}

func Endpoints() map[string]int64 {
	mu.Lock()
	defer mu.Unlock()

	copy := make(map[string]int64, len(endpointVisits))
	for k, v := range endpointVisits {
		copy[k] = v
	}

	return copy
}

func IncVisitor(hash string) {
	mu.Lock()
	visitors[hash] = struct{}{}
	mu.Unlock()
}

func Snapshot() (int64, map[string]int64, []string) {

	total := atomic.SwapInt64(&totalVisits, 0)

	mu.Lock()

	endpoints := make(map[string]int64, len(endpointVisits))
	for k, v := range endpointVisits {
		endpoints[k] = v
		endpointVisits[k] = 0
	}

	visitorList := make([]string, 0, len(visitors))
	for v := range visitors {
		visitorList = append(visitorList, v)
	}

	visitors = make(map[string]struct{})

	mu.Unlock()

	return total, endpoints, visitorList
}
