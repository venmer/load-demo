package metrics

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	HttpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	HttpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)

	UsersCreated = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "users_created_total",
			Help: "Total number of created users",
		},
	)

	UsersDeleted = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "users_deleted_total",
			Help: "Total number of deleted users",
		},
	)

	DatabaseErrors = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "database_errors_total",
			Help: "Total number of database errors",
		},
	)
)

// Middleware для сбора HTTP метрик
func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		path := r.URL.Path

		// Создаем ResponseWriter для перехвата статуса
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(rw, r)

		duration := time.Since(start).Seconds()

		// Сохраняем метрики
		HttpRequestsTotal.WithLabelValues(
			r.Method,
			path,
			http.StatusText(rw.statusCode),
		).Inc()

		HttpRequestDuration.WithLabelValues(
			r.Method,
			path,
		).Observe(duration)
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// Handler для Prometheus
func MetricsHandler() http.Handler {
	return promhttp.Handler()
}
