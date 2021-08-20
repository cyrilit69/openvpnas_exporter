/*
This package contains a struct for the 'sacli VPNSummary' output marshaling and
the only Parse([]byte) func to marshal and cast some data into metrics format
*/

package summaryparser

import (
	"encoding/json"
	"log"
	"time"

	"github.com/cyrilit69/openvpnas_exporter/models"
)

// Struct is used for unmarshaling only. VPNSummary will be created based on it
type raw struct {
	ActiveProfile string                 `json:"active_profile"`
	Errors        map[string]interface{} `json:"errors"`
	LR            string                 `json:"last_restarted"`
	ServiceStatus map[string]string      `json:"service_status"`
}

// Struct contains all

func Parse(in []byte) (*models.VPNSummary, error) {
	res := new(models.VPNSummary)
	r := new(raw)
	err := json.Unmarshal(in, r)
	if err != nil {
		return res, err
	}
	res.ActiveProfile = r.ActiveProfile
	lrt, err := time.Parse(time.ANSIC, r.LR)
	if err != nil {
		log.Printf("cannot parse Last Restart time to time.Time: %v", err)
		res.LastRestarted = 0
	}
	res.LastRestarted = float64(lrt.Unix())
	res.ServiceStatusTotal = make(map[string]float64)
	for k, v := range r.ServiceStatus {
		if v == "on" {
			res.ServiceStatusTotal[k] = 1
		} else {
			res.ServiceStatusTotal[k] = 0
		}
	}
	res.ErrorsTotal = float64(len(r.Errors))
	return res, nil
}
