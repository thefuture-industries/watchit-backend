package configuration

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	RequestDuration *prometheus.HistogramVec
	RequestTotal    *prometheus.CounterVec
	RequestErrors   *prometheus.CounterVec

	QueryDuration prometheus.Histogram
	ErrorCount    prometheus.Counter
	DBConnections prometheus.Gauge

	ActiveUsers prometheus.Gauge
}

// NewMetrics создание метрик
// --------------------------
func NewMetrics() *Metrics {
	reg := prometheus.NewRegistry()

	// Создание метрик
	m := &Metrics{
		// Метрика для измерения длительности запросов
		RequestDuration: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
		}, []string{"handler", "method", "status"}),
		// Метрика для подсчета количества запросов
		RequestTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		}, []string{"handler", "method", "status"}),

		// Метрика для измерения длительности запросов
		QueryDuration: prometheus.NewHistogram(prometheus.HistogramOpts{
			Name: "recommendation_query_duration_seconds",
			Help: "Duration of recommendation queries in seconds",
		}),
		// Метрика для подсчета количества ошибок
		ErrorCount: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "recommendation_errors_total",
			Help: "Total number of recommendation errors",
		}),
		// Метрика для подсчета количества открытых соединений с БД
		DBConnections: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "db_open_connections",
			Help: "Number of open database connections",
		}),

		ActiveUsers: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "active_user_count",
			Help: "Number of active users",
		}),
	}

	// Регистрация метрик в Prometheus
	reg.MustRegister(
		m.RequestDuration,
		m.RequestTotal,
		m.QueryDuration,
		m.ErrorCount,
		m.DBConnections,
		m.ActiveUsers,
	)

	return m
}
