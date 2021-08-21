package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/cyrilit69/openvpnas_exporter/internal/exporter"
)

var (
	listenAddress = flag.String("web.listen-address", ":9185",
		"Address to listen on for telemetry")
	metricsPath = flag.String("web.telemetry-path", "/metrics",
		"Path under which to expose metrics")
	sacliPath = flag.String("sacli-path", "/usr/local/openvpn_as/scripts/sacli",
		"Path to 'sacli' script")
)

func init() {
	flag.Parse()
}

func main() {
	exporter, err := exporter.NewExporter(*sacliPath)
	if err != nil {
		log.Fatalf("cannot start the exporter: %v", err)
	}
	prometheus.MustRegister(exporter)

	http.Handle(*metricsPath, promhttp.Handler())
	log.Fatal(http.ListenAndServe(*listenAddress, nil))

}
