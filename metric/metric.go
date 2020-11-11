package metric

import "time"

// Metric struct
type Metric struct {
	Count      int
	UpdateTime time.Time
}

// HandleMetrics Interface for handle metric saveing in memory or in database
type HandleMetrics interface {
	AddMetric(key string, t time.Time)
	GetMetricCount(key string, t time.Time) (int, error)
}

// PostMetricResponse empty response
type PostMetricResponse struct {
}

// GetMetricCountResponse success response
type GetMetricCountResponse struct {
	Value int
}
