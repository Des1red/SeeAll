package devmode

import (
	"sync"
	"time"
)

type sourceMetric struct {
	Duration time.Duration
	Posts    int
}

var (
	sourceMetrics   = make(map[string]sourceMetric)
	sourceMetricsMu sync.Mutex
)

func RecordSource(name string, duration time.Duration, posts int) {
	if isPrewarming.Load() {
		return
	}

	sourceMetricsMu.Lock()
	defer sourceMetricsMu.Unlock()
	sourceMetrics[name] = sourceMetric{
		Duration: duration,
		Posts:    posts,
	}
}

func getSourceMetrics() map[string]sourceMetric {
	sourceMetricsMu.Lock()
	defer sourceMetricsMu.Unlock()

	copy := make(map[string]sourceMetric, len(sourceMetrics))
	for k, v := range sourceMetrics {
		copy[k] = v
	}
	return copy
}

func resetSourceMetrics() {
	sourceMetricsMu.Lock()
	defer sourceMetricsMu.Unlock()
	sourceMetrics = make(map[string]sourceMetric)
}
