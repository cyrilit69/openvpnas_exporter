package exporter

import "github.com/prometheus/client_golang/prometheus"

const (
	namespace = "openvpnas"
)

var (
	// Status metrics
	summaryParsed = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "status", "parsed"),
		"Is 'sacli status' output parsed succesfully",
		nil, nil,
	)

	summaryErrors = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "status", "errors"),
		"Count of errors in sacli status output",
		nil, nil,
	)

	summaryLastRestarted = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "status", "last_restarted"),
		"Timestamp of the last restart",
		nil, nil,
	)

	summaryStatus = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "status", "service_state"),
		"States of all services from sacli status output",
		[]string{"service"}, nil,
	)

	// SubscriptionStatus metrics
	subsParsed = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "subscription", "parsed"),
		"Is 'sacli SubscriptionStatus' output parsed succesfully",
		nil, nil,
	)
	subsAgentDisabled = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "subscription", "agent_disabled"),
		"If subscription agent disabled", nil, nil,
	)
	subsCCLimit = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "subscription", "cc_limit"),
		"The limit of clients connections (CCs)", nil, nil,
	)
	subsCurrentCC = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "subscription", "current_cc"),
		"Client connected (from subscription status)", nil, nil,
	)
	subsError = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "subscription", "error"),
		"Subscription error", nil, nil,
	)
	subsFallbackCC = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "subscription", "fallback_cc"),
		"Fallback clients connections", nil, nil,
	)
	subsGracePeriod = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "subscription", "grace_period"),
		"Grace period", nil, nil,
	)
	subsLastSuccessfulUpdate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "subscription", "last_successful_update"),
		"Timestamp of the last successful update", nil, nil,
	)
	subsLastSuccessfulUpdateAge = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "subscription", "last_successful_update_age"),
		"Last successful update age (wtf?)", nil, nil,
	)
	subsMaxCC = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "subscription", "max_cc"),
		"max clients connections (what is the difference between this and limit cc?)", nil, nil,
	)
	subsName = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "subscription", "name"),
		"Subscription name", []string{"name"}, nil,
	)
	subsNextUpdate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "subscription", "next_update"),
		"Timestamp fo the next update", nil, nil,
	)
	subsNextUpdateIn = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "subscription", "next_update_in"),
		"Seconds (?) before the next update", nil, nil,
	)
	subsNotes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "subscription", "notes"),
		"Some notes", []string{"note"}, nil,
	)
	subsOverdraft = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "subscription", "overdraft"),
		"Subscription Overdraft", nil, nil,
	)
	subsServer = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "subscription", "server"),
		"Subscription Server", []string{"server"}, nil,
	)
	subsState = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "subscription", "state"),
		"Subscription State", []string{"state"}, nil, // TODO find all possible values
	)
	subsType = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "subscription", "type"),
		"Subscription type", []string{"type"}, nil,
	)
	subsUpdatesFailed = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "subscription", "updates_failed"),
		"Subscription updates failed", nil, nil,
	)
	// VPNStatus metrics
	statusParsed = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "vpnstatus", "parsed"),
		"Is 'sacli VPNStatus' output parsed succesfully",
		nil, nil,
	)
	statusClientInfo = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "vpnstatus", "client_info"),
		"All Clients info", []string{"vpn", "common_name", "id", "peer_id", "real_addr", "vpn_addr"}, nil,
	)
	statusClientBytesReceived = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "vpnstatus", "client_bytes_received"),
		"Bytes received by client", []string{"vpn", "common_name", "id", "peer_id", "real_addr", "vpn_addr"}, nil,
	)
	statusClientBytesSend = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "vpnstatus", "client_bytes_send"),
		"Bytes sent by client", []string{"vpn", "common_name", "id", "peer_id", "real_addr", "vpn_addr"}, nil,
	)
	statusClientConnectedSince = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "vpnstatus", "client_connected_since"),
		"Timestamp of the last clients connect", []string{"vpn", "common_name", "id", "peer_id", "real_addr", "vpn_addr"}, nil,
	)
)

// This map is used just to simplify the Describe() func and create
// a slice of metrics with specific length
var OpenVpnASMetrics = map[string][]*prometheus.Desc{
	// The list of exporter metrics from 'sacli status' command
	"status": []*prometheus.Desc{
		summaryParsed,
		summaryErrors,
		summaryLastRestarted,
		summaryStatus,
	},
	"SubscriptionStatus": []*prometheus.Desc{
		subsParsed,
		subsAgentDisabled,
		subsCCLimit,
		subsCurrentCC,
		subsError,
		subsFallbackCC,
		subsGracePeriod,
		subsLastSuccessfulUpdate,
		subsLastSuccessfulUpdateAge,
		subsMaxCC,
		subsName,
		subsNextUpdate,
		subsNextUpdateIn,
		subsNotes,
		subsOverdraft,
		subsServer,
		subsState,
		subsType,
		subsUpdatesFailed,
	},
	"VPNStatus": []*prometheus.Desc{
		statusParsed,
		statusClientBytesReceived,
		statusClientBytesSend,
		statusClientConnectedSince,
		statusClientInfo,
	},
}
