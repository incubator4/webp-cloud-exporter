package server

import (
	"fmt"
	"github.com/incubator4/webp-cloud-exporter/pkg/webpse"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var (
	userLabel    = []string{"uuid"}
	metricPrefix = "webp_cloud"

	handler = promhttp.Handler()

	userQuotaLabel    = []string{"uuid", "name", "email", "type"}
	gaugeVecUserQuota = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: fmt.Sprintf("%s_user_quota", metricPrefix),
			Help: "User quota from webp-cloud",
		}, userQuotaLabel)

	userSendByteTotalLabel    = []string{"uuid"}
	gaugeVecUserSendByteTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: fmt.Sprintf("%s_user_sent_total", metricPrefix),
			Help: "User sent total from webp-cloud",
		}, userSendByteTotalLabel)

	GaugeVecProxy = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{}, []string{})
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func Healthz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

func Metrics(client *webpse.Client) http.Handler {
	if resp, err := client.GetUserInfo(); err == nil {
		gaugeVecUserQuota.
			WithLabelValues(resp.Data.UUID, resp.Data.Name, resp.Data.Email, "daily").
			Set(float64(resp.Data.DailyQuota))
		gaugeVecUserQuota.
			WithLabelValues(resp.Data.UUID, resp.Data.Name, resp.Data.Email, "daily_limit").
			Set(float64(resp.Data.DailyQuotaLimit))
		gaugeVecUserQuota.
			WithLabelValues(resp.Data.UUID, resp.Data.Name, resp.Data.Email, "permanent").
			Set(float64(resp.Data.PermanentQuota))

	}

	if resp, err := client.GetUserStats(); err == nil {
		gaugeVecUserSendByteTotal.WithLabelValues(resp.Data.UUID).Set(float64(resp.Data.TotalBytesSent))
	}

	//if resp, err := client.GetProxiesStats(); err == nil {
	//	for _, proxy := range resp.Data {
	//		GaugeVecProxy.WithLabelValues(proxy.UUID).Set(float64(proxy.CacheSize))
	//	}
	//}

	return handler
}
