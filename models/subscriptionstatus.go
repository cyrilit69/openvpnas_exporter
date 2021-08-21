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

	var ok bool
	// boolean
	adis, ok := v["agent_disabled"].(bool)
	if !ok {
		log.Printf("cannot get 'agent_disabled' from the sacli output: %v", string(data))
	}
	if adis {
		c.AgentDisabled = 1
	}

	od, ok := v["overdraft"].(bool)
	if !ok {
		log.Printf("cannot get 'overdraft' from the sacli output: %v", string(data))
	}
	if od {
		c.Overdraft = 1
	}

	// ints
	c.CcLimit, ok = v["cc_limit"].(float64)
	if !ok {
		log.Printf("cannot get 'cc_limit' from the sacli output: %v", string(data))
	}
	c.CurrentCc, ok = v["current_cc"].(float64)
	if !ok {
		log.Printf("cannot get 'current_cc' from the sacli output: %v", string(data))
	}
	c.FallbackCc, ok = v["fallback_cc"].(float64)
	if !ok {
		log.Printf("cannot get 'fallback_cc' from the sacli output: %v", string(data))
	}
	c.GracePeriod, ok = v["grace_period"].(float64)
	if !ok {
		log.Printf("cannot get 'grace_period' from the sacli output: %v", string(data))
	}
	c.LastSuccessfulUpdate, ok = v["last_successful_update"].(float64)
	if !ok {
		log.Printf("cannot get 'last_successful_update' from the sacli output: %v", string(data))
	}
	c.LastSuccessfulUpdateAge, ok = v["last_successful_update_age"].(float64)
	if !ok {
		log.Printf("cannot get 'last_successful_update_age' from the sacli output: %v", string(data))
	}
	c.MaxCc, ok = v["max_cc"].(float64)
	if !ok {
		log.Printf("cannot get 'max_cc' from the sacli output: %v", string(data))
	}
	c.NextUpdate, ok = v["next_update"].(float64)
	if !ok {
		log.Printf("cannot get 'next_update' from the sacli output: %v", string(data))
	}
	c.NextUpdateIn, ok = v["next_update_in"].(float64)
	if !ok {
		log.Printf("cannot get 'next_update_in' from the sacli output: %v", string(data))
	}
	c.UpdatesFailed, ok = v["updates_failed"].(float64)
	if !ok {
		log.Printf("cannot get 'updates_failed' from the sacli output: %v", string(data))
	}
	// strings
	c.Name = v["name"].(string)
	if !ok {
		log.Printf("cannot get 'name' from the sacli output: %v", string(data))
	}
	c.Server = v["server"].(string)
	if !ok {
		log.Printf("cannot get 'server' from the sacli output: %v", string(data))
	}
	c.State = v["state"].(string)
	if !ok {
		log.Printf("cannot get 'state' from the sacli output: %v", string(data))
	}
	c.Type = v["type"].(string)
	if !ok {
		log.Printf("cannot get 'type' from the sacli output: %v", string(data))
	}
	// errors
	if v["error"] != nil {
		c.Error = 1
	}
	for _, n := range v["notes"].([]interface{}) {
		c.Notes = append(c.Notes, fmt.Sprint(n))
	}
	return nil
}
