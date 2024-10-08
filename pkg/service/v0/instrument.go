package svc

import (
	"context"

	"github.com/anaswaratrajan/ocis-jupyter/pkg/metrics"
	v0proto "github.com/anaswaratrajan/ocis-jupyter/pkg/proto/v0"
	"github.com/prometheus/client_golang/prometheus"
)

// NewInstrument returns a service that instruments metrics.
func NewInstrument(next v0proto.JupyterNotebookSupportHandler, metrics *metrics.Metrics) v0proto.JupyterNotebookSupportHandler {
	return instrument{
		next:    next,
		metrics: metrics,
	}
}

type instrument struct {
	next    v0proto.JupyterNotebookSupportHandler
	metrics *metrics.Metrics
}

// Greet implements the HelloHandler interface.
func (i instrument) GenerateHTML(ctx context.Context, req *v0proto.JupyterNotebookJSON, rsp *v0proto.JupyterNotebookHTML) error {
	timer := prometheus.NewTimer(prometheus.ObserverFunc(func(v float64) {
		us := v * 1000000

		i.metrics.Latency.WithLabelValues().Observe(us)
		i.metrics.Duration.WithLabelValues().Observe(v)
	}))

	defer timer.ObserveDuration()

	err := i.next.GenerateHTML(ctx, req, rsp)

	if err == nil {
		i.metrics.Counter.WithLabelValues().Inc()
	}

	return err
}
