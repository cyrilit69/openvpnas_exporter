package models

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
	Package contains structs for the metrics building. Fields could be:
	* float64 (just a value)
	* string or []string (will be a label and its value will be '1')
	* map[string]float64 (label + value)
*/

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

func (c *SubscriptionStatus) UnmarshalJSON(data []byte) error {
	var v map[string]interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	failedToParse := []string{}
	var ok bool

	// boolean
	adis, ok := v["agent_disabled"].(bool)
	if !ok {
		failedToParse = append(failedToParse, "agent_disabled")
	}
	if adis {
		c.AgentDisabled = 1
	}

	od, ok := v["overdraft"].(bool)
	if !ok {
		failedToParse = append(failedToParse, "overdraft")
	}
	if od {
		c.Overdraft = 1
	}

	// ints
	c.CcLimit, ok = v["cc_limit"].(float64)
	if !ok {
		failedToParse = append(failedToParse, "cc_limit")
	}
	c.CurrentCc, ok = v["current_cc"].(float64)
	if !ok {
		failedToParse = append(failedToParse, "current_cc")
	}
	c.FallbackCc, ok = v["fallback_cc"].(float64)
	if !ok {
		failedToParse = append(failedToParse, "fallback_cc")
	}
	c.GracePeriod, ok = v["grace_period"].(float64)
	if !ok {
		failedToParse = append(failedToParse, "grace_period")
	}
	c.LastSuccessfulUpdate, ok = v["last_successful_update"].(float64)
	if !ok {
		failedToParse = append(failedToParse, "last_successful_update")
	}
	c.LastSuccessfulUpdateAge, ok = v["last_successful_update_age"].(float64)
	if !ok {
		failedToParse = append(failedToParse, "last_successful_update_age")
	}
	c.MaxCc, ok = v["max_cc"].(float64)
	if !ok {
		failedToParse = append(failedToParse, "max_cc")
	}
	c.NextUpdate, ok = v["next_update"].(float64)
	if !ok {
		failedToParse = append(failedToParse, "next_update")
	}
	c.NextUpdateIn, ok = v["next_update_in"].(float64)
	if !ok {
		failedToParse = append(failedToParse, "next_update_in")
	}
	c.UpdatesFailed, ok = v["updates_failed"].(float64)
	if !ok {
		failedToParse = append(failedToParse, "updates_failed")
	}
	// strings
	c.Name = v["name"].(string)
	if !ok {
		failedToParse = append(failedToParse, "name")
	}
	c.Server = v["server"].(string)
	if !ok {
		failedToParse = append(failedToParse, "server")
	}
	c.State = v["state"].(string)
	if !ok {
		failedToParse = append(failedToParse, "state")
	}
	c.Type = v["type"].(string)
	if !ok {
		failedToParse = append(failedToParse, "type")
	}
	// errors
	if v["error"] != nil {
		c.Error = 1
	}
	for _, n := range v["notes"].([]interface{}) {
		c.Notes = append(c.Notes, fmt.Sprint(n))
	}
	if len(failedToParse) > 0 {
		log.Printf("cannot get '%v' from the sacli output: %v", failedToParse, string(data))
	}
	return nil
}
