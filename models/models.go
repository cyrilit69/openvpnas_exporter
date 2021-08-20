package models

/*
	Package contains structs for the metrics building. Fields could be:
	* float64 (just a value)
	* string or []string (will be a label and its value will be '1')
	* map[string]float64 (label + value)
*/

type VPNSummary struct {
	ActiveProfile      string
	ErrorsTotal        float64
	LastRestarted      float64
	ServiceStatusTotal map[string]float64
}

type SubscriptionStatus struct {
	AgentDisabled           float64
	CcLimit                 float64
	CurrentCc               float64
	Error                   float64
	FallbackCc              float64
	GracePeriod             float64
	LastSuccessfulUpdate    float64
	LastSuccessfulUpdateAge float64
	MaxCc                   float64
	NextUpdate              float64
	NextUpdateIn            float64
	Overdraft               float64
	UpdatesFailed           float64
	Name                    string
	Server                  string
	State                   string
	Type                    string
	Notes                   []string
}

type VPNStatus struct {
	ClientVPN        string
	ClientName       string
	ClientId         string
	ClientPeerId     string
	RealAddr         string
	VPNAddr          string
	BytesReceived    float64
	BytesSend        float64
	ConnectedSinceTs float64
}
