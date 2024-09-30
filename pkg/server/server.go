package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/incubator4/webp-cloud-exporter/pkg/webpse"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/version"
	"github.com/prometheus/exporter-toolkit/web"
	"log"
	"net/http"
)

var (
	metricsPath   = "/metrics"
	landingConfig = web.LandingConfig{
		Name:        "Webp-cloud Exporter",
		Description: "A Prometheus exporter for webp-cloud",
		Version:     version.Info(),
		Links: []web.LandingLinks{
			{
				Address: metricsPath,
				Text:    "Metrics",
			},
		},
	}
)

type Server struct {
	mux *mux.Router
}

func New(client *webpse.Client) Server {
	r := mux.NewRouter()

	r.Use(loggingMiddleware)

	landingPage, err := web.NewLandingPage(landingConfig)
	if err != nil {
		panic(err)
	}
	prometheus.MustRegister(gaugeVecUserQuota, gaugeVecUserSendByteTotal)

	r.Handle("/", landingPage)
	r.HandleFunc("/healthz", Healthz).Methods(http.MethodGet)
	r.HandleFunc(metricsPath, Metrics(client)).Methods(http.MethodGet)

	return Server{mux: r}
}

func (s *Server) Start(port int) error {
	log.Printf("Starting server on port %d", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), s.mux)
}
