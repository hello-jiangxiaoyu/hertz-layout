package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/tracer"
	"github.com/cloudwego/hertz/pkg/common/tracer/stats"
	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"hertz/demo/internal/utils"
	"net/http"
	"strconv"
	"time"
)

const (
	labelMethod     = "method"
	labelStatusCode = "statusCode"
	labelPath       = "path"

	unknownLabelValue = "unknown"
)

type serverTracer struct {
	serverHandledCounter   *prom.CounterVec
	serverHandledHistogram *prom.HistogramVec
}

// Start record the beginning of server handling request from client.
func (s *serverTracer) Start(ctx context.Context, c *app.RequestContext) context.Context {
	return ctx
}

// Finish record the ending of server handling request from client.
func (s *serverTracer) Finish(ctx context.Context, c *app.RequestContext) {
	if c.GetTraceInfo().Stats().Level() == stats.LevelDisabled {
		return
	}

	httpStart := c.GetTraceInfo().Stats().GetEvent(stats.HTTPStart)
	httpFinish := c.GetTraceInfo().Stats().GetEvent(stats.HTTPFinish)
	if httpFinish == nil || httpStart == nil {
		return
	}
	cost := httpFinish.Time().Sub(httpStart.Time())
	_ = counterAdd(s.serverHandledCounter, 1, genLabels(c))
	_ = histogramObserve(s.serverHandledHistogram, cost, genLabels(c))
}

// counterAdd wraps Add of prom.Counter.
func counterAdd(counterVec *prom.CounterVec, value int, labels prom.Labels) error {
	counter, err := counterVec.GetMetricWith(labels)
	if err != nil {
		return err
	}
	counter.Add(float64(value))
	return nil
}

// histogramObserve wraps Observe of prom.Observer.
func histogramObserve(histogramVec *prom.HistogramVec, value time.Duration, labels prom.Labels) error {
	histogram, err := histogramVec.GetMetricWith(labels)
	if err != nil {
		return err
	}
	histogram.Observe(float64(value.Microseconds()))
	return nil
}

// genLabels make labels values.
func genLabels(ctx *app.RequestContext) prom.Labels {
	labels := make(prom.Labels)
	labels[labelMethod] = utils.DefaultValIfEmpty(string(ctx.Request.Method()), unknownLabelValue)
	labels[labelStatusCode] = utils.DefaultValIfEmpty(strconv.Itoa(ctx.Response.Header.StatusCode()), unknownLabelValue)
	labels[labelPath] = utils.DefaultValIfEmpty(string(ctx.Request.Path()), unknownLabelValue)

	return labels
}

func PrometheusHandler(addr, path string) tracer.Tracer {
	registry := prom.NewRegistry()
	enableGoCollector := false
	http.Handle(path, promhttp.HandlerFor(registry, promhttp.HandlerOpts{ErrorHandling: promhttp.ContinueOnError}))
	go func() {
		if err := http.ListenAndServe(addr, nil); err != nil {
			hlog.Fatal("HERTZ: Unable to start a prom http server, err: " + err.Error())
		}
	}()

	serverHandledCounter := prom.NewCounterVec(
		prom.CounterOpts{
			Name: "hertz_server_throughput",
			Help: "Total number of HTTPs completed by the server, regardless of success or failure.",
		},
		[]string{labelMethod, labelStatusCode, labelPath},
	)
	registry.MustRegister(serverHandledCounter)
	serverHandledHistogram := prom.NewHistogramVec(
		prom.HistogramOpts{
			Name:    "hertz_server_latency_us",
			Help:    "Latency (microseconds) of HTTP that had been application-level handled by the server.",
			Buckets: []float64{5000, 10000, 25000, 50000, 100000, 250000, 500000, 1000000},
		},
		[]string{labelMethod, labelStatusCode, labelPath},
	)
	registry.MustRegister(serverHandledHistogram)
	if enableGoCollector {
		registry.MustRegister(collectors.NewGoCollector())
	}

	return &serverTracer{
		serverHandledCounter:   serverHandledCounter,
		serverHandledHistogram: serverHandledHistogram,
	}
}
