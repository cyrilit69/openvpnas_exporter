package substatusparser

import (
	"encoding/json"
	"fmt"

	"github.com/cyrilit69/openvpnas_exporter/models"
)

type raw struct {
	AgentDisabled           bool          `json:"agent_disabled"`
	Overdraft               bool          `json:"overdraft"`
	CcLimit                 int           `json:"cc_limit"`
	CurrentCc               int           `json:"current_cc"`
	FallbackCc              int           `json:"fallback_cc"`
	GracePeriod             int           `json:"grace_period"`
	LastSuccessfulUpdate    int           `json:"last_successful_update"`
	LastSuccessfulUpdateAge int           `json:"last_successful_update_age"`
	MaxCc                   int           `json:"max_cc"`
	NextUpdate              int           `json:"next_update"`
	NextUpdateIn            int           `json:"next_update_in"`
	UpdatesFailed           int           `json:"updates_failed"`
	Name                    string        `json:"name"`
	Server                  string        `json:"server"`
	State                   string        `json:"state"`
	Type                    string        `json:"type"`
	Error                   interface{}   `json:"error"`
	Notes                   []interface{} `json:"notes"`
}

func Parse(in []byte) (*models.SubscriptionStatus, error) {
	res := new(models.SubscriptionStatus)
	r := new(raw)
	err := json.Unmarshal(in, r)
	if err != nil {
		return res, err
	}
	// boolean
	if r.AgentDisabled {
		res.AgentDisabled = 1
	} // else default (0)
	if r.Overdraft {
		res.Overdraft = 1
	}
	// ints
	res.CcLimit = float64(r.CcLimit)
	res.CurrentCc = float64(r.CurrentCc)
	res.FallbackCc = float64(r.FallbackCc)
	res.GracePeriod = float64(r.GracePeriod)
	res.LastSuccessfulUpdate = float64(r.LastSuccessfulUpdate)
	res.LastSuccessfulUpdateAge = float64(r.LastSuccessfulUpdateAge)
	res.MaxCc = float64(r.MaxCc)
	res.NextUpdate = float64(r.NextUpdate)
	res.NextUpdateIn = float64(r.NextUpdateIn)
	res.UpdatesFailed = float64(r.UpdatesFailed)
	// strings
	res.Name = r.Name
	res.Server = r.Server
	res.State = r.State
	res.Type = r.Type
	// errors
	if r.Error != nil {
		res.Error = 1
	}
	for _, n := range r.Notes {
		res.Notes = append(res.Notes, fmt.Sprint(n))
	}
	return res, nil
}
