package exporter

import (
	"log"
	"os"
	"os/exec"

	"github.com/cyrilit69/openvpnas_exporter/internal/statusparser"
	"github.com/cyrilit69/openvpnas_exporter/internal/substatusparser"
	"github.com/cyrilit69/openvpnas_exporter/internal/summaryparser"

	"github.com/prometheus/client_golang/prometheus"
)

type Exporter struct {
	sacliPath string
}

func NewExporter(sacliPath string) (*Exporter, error) {
	if _, err := os.Stat(sacliPath); err != nil {
		return &Exporter{sacliPath}, err
	}
	return &Exporter{sacliPath}, nil
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	for _, v := range OpenVpnASMetrics {
		for _, m := range v {
			ch <- m
		}
	}
}

func (e *Exporter) runSacli(arg string) ([]byte, error) {
	out, err := exec.Command(e.sacliPath, arg).Output()
	if err != nil {
		return out, err
	}
	return out, nil
}

func (e *Exporter) summaryParse(data []byte, ch chan<- prometheus.Metric) {
	s, err := summaryparser.Parse(data)
	if err != nil {
		log.Printf("error during VPNSummary parsing: %v", err)
		ch <- prometheus.MustNewConstMetric(summaryParsed, prometheus.GaugeValue, 0)
		return
	}
	ch <- prometheus.MustNewConstMetric(summaryParsed, prometheus.GaugeValue, 1)
	ch <- prometheus.MustNewConstMetric(summaryActiveProfile, prometheus.GaugeValue, 1, s.ActiveProfile)
	ch <- prometheus.MustNewConstMetric(summaryLastRestarted, prometheus.GaugeValue, s.LastRestarted)
	ch <- prometheus.MustNewConstMetric(summaryErrors, prometheus.GaugeValue, s.ErrorsTotal)
	for k, v := range s.ServiceStatusTotal {
		ch <- prometheus.MustNewConstMetric(summaryStatus, prometheus.GaugeValue, v, k)
	}
}

func (e *Exporter) subsParse(data []byte, ch chan<- prometheus.Metric) {
	s, err := substatusparser.Parse(data)
	if err != nil {
		log.Printf("error during SubscriptionStatus parsing: %v", err)
		ch <- prometheus.MustNewConstMetric(subsParsed, prometheus.GaugeValue, 0)
		return
	}
	ch <- prometheus.MustNewConstMetric(subsParsed, prometheus.GaugeValue, 1)
	// float64
	ch <- prometheus.MustNewConstMetric(subsAgentDisabled, prometheus.GaugeValue, s.AgentDisabled)
	ch <- prometheus.MustNewConstMetric(subsCCLimit, prometheus.GaugeValue, s.CcLimit)
	ch <- prometheus.MustNewConstMetric(subsCurrentCC, prometheus.GaugeValue, s.CurrentCc)
	ch <- prometheus.MustNewConstMetric(subsError, prometheus.GaugeValue, s.Error)
	ch <- prometheus.MustNewConstMetric(subsFallbackCC, prometheus.GaugeValue, s.FallbackCc)
	ch <- prometheus.MustNewConstMetric(subsGracePeriod, prometheus.GaugeValue, s.GracePeriod)
	ch <- prometheus.MustNewConstMetric(subsLastSuccessfulUpdate, prometheus.GaugeValue, s.LastSuccessfulUpdate)
	ch <- prometheus.MustNewConstMetric(subsLastSuccessfulUpdateAge, prometheus.GaugeValue, s.LastSuccessfulUpdateAge)
	ch <- prometheus.MustNewConstMetric(subsMaxCC, prometheus.GaugeValue, s.MaxCc)
	ch <- prometheus.MustNewConstMetric(subsNextUpdate, prometheus.GaugeValue, s.NextUpdate)
	ch <- prometheus.MustNewConstMetric(subsNextUpdateIn, prometheus.GaugeValue, s.NextUpdateIn)
	ch <- prometheus.MustNewConstMetric(subsOverdraft, prometheus.GaugeValue, s.Overdraft)
	ch <- prometheus.MustNewConstMetric(subsUpdatesFailed, prometheus.GaugeValue, s.UpdatesFailed)
	// string
	ch <- prometheus.MustNewConstMetric(subsName, prometheus.GaugeValue, 1, s.Name)
	ch <- prometheus.MustNewConstMetric(subsServer, prometheus.GaugeValue, 1, s.Server)
	ch <- prometheus.MustNewConstMetric(subsState, prometheus.GaugeValue, 1, s.State)
	ch <- prometheus.MustNewConstMetric(subsType, prometheus.GaugeValue, 1, s.Type)
	// []string
	if len(s.Notes) > 0 {
		for _, n := range s.Notes {
			ch <- prometheus.MustNewConstMetric(subsNotes, prometheus.GaugeValue, 1, n)
		}
	}
}

func (e *Exporter) statusParse(data []byte, ch chan<- prometheus.Metric) {
	s, err := statusparser.Parse(data)
	if err != nil {
		log.Printf("error during VPNStatus parsing: %v", err)
		ch <- prometheus.MustNewConstMetric(statusParsed, prometheus.GaugeValue, 0)
		return
	}
	ch <- prometheus.MustNewConstMetric(statusParsed, prometheus.GaugeValue, 1)
	for _, c := range s {
		// float64
		ch <- prometheus.MustNewConstMetric(statusClientBytesReceived, prometheus.GaugeValue, c.BytesReceived, c.ClientName)
		ch <- prometheus.MustNewConstMetric(statusClientBytesSend, prometheus.GaugeValue, c.BytesSend, c.ClientName)
		ch <- prometheus.MustNewConstMetric(statusClientConnectedSince, prometheus.GaugeValue, c.ConnectedSinceTs, c.ClientName)
		// string ("vpn", "common_name", "id", "peer_id", "real_addr", "vpn_addr")
		ch <- prometheus.MustNewConstMetric(statusClientInfo, prometheus.GaugeValue, 1,
			c.ClientVPN,
			c.ClientName,
			c.ClientId,
			c.ClientPeerId,
			c.RealAddr,
			c.VPNAddr,
		)
	}

}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	// Working with sacli VPNSummary
	summaryData, err := e.runSacli("VPNSummary")
	if err != nil {
		log.Printf("cannot execute 'sacli VPNSummary' command: %v", err)
		ch <- prometheus.MustNewConstMetric(summaryParsed, prometheus.GaugeValue, 0)
	} else {
		e.summaryParse(summaryData, ch)
	}
	// Working with sacli SubscriptionStatus
	subsData, err := e.runSacli("SubscriptionStatus")
	if err != nil {
		log.Printf("cannot execute 'sacli SubscriptionStatus' command: %v", err)
		ch <- prometheus.MustNewConstMetric(
			subsParsed, prometheus.GaugeValue, 0,
		)
	} else {
		e.subsParse(subsData, ch)
	}
	// Working with sacli VPNStatus
	statusData, err := e.runSacli("VPNStatus")
	if err != nil {
		log.Printf("cannot execute 'sacli VPNStatus' command: %v", err)
		ch <- prometheus.MustNewConstMetric(
			statusParsed, prometheus.GaugeValue, 0,
		)
	} else {
		e.statusParse(statusData, ch)
	}
}
