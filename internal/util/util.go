// internal/util/util.go
package util

import (
	"context"
	"time"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

func QueryPrometheus(ctx context.Context, apiClient v1.API, query string, start, end time.Time) (model.Value, []string, error) {
	// Logic to query Prometheus
	// ...

	return nil, nil, nil
}

func ParseTimeRange(timeRange string) (time.Time, time.Time, error) {
	// Logic to parse time range
	// ...

	return time.Time{}, time.Time{}, nil
}
