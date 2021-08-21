package models

/*
	Package contains structs for the metrics building. Fields could be:
	* float64 (just a value)
	* string or []string (will be a label and its value will be '1')
	* map[string]float64 (label + value)
*/

import (
	"encoding/json"
	"log"
	"time"
)

type Status struct {
	// ActiveProfile string
	ErrorsTotal   float64
	LastRestarted float64
	ServiceStatus map[string]float64
}

// type raw struct {
// 	ActiveProfile string                 `json:"active_profile"`
// 	Errors        map[string]interface{} `json:"errors"`
// 	LR            string                 `json:"last_restarted"`
// 	ServiceStatus map[string]string      `json:"service_status"`
// }

func (c *Status) UnmarshalJSON(data []byte) error {
	var v map[string]interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	var ok bool

	// c.ActiveProfile, ok = v["active_profile"].(string)
	// if !ok {
	// 	log.Printf("cannot get 'active_profile' from the sacli output: %v", string(data))
	// }

	errs, ok := v["errors"].(map[string]interface{})
	if !ok {
		log.Printf("cannot get 'errors' from the sacli output: %v", string(data))
	}
	c.ErrorsTotal = float64(len(errs))

	lr, ok := v["last_restarted"].(string)
	if !ok {
		log.Printf("cannot get 'last_restarted' from the sacli output: %v", string(data))
	}
	lrt, err := time.Parse(time.ANSIC, lr)
	if err != nil {
		log.Printf("cannot parse Last Restart time to time.Time: %v", err)
	}
	c.LastRestarted = float64(lrt.Unix())

	c.ServiceStatus = make(map[string]float64)

	ss, ok := v["service_status"].(map[string]interface{})
	if !ok {
		log.Printf("cannot get 'service_status' from the sacli output: %v", string(data))
	}
	for k, v := range ss {
		if v.(string) == "on" {
			c.ServiceStatus[k] = 1
		} else {
			c.ServiceStatus[k] = 0
		}
	}

	return nil
}
