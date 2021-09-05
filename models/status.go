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
	ErrorsTotal   float64
	LastRestarted float64
	ServiceStatus map[string]float64
}

func (c *Status) UnmarshalJSON(data []byte) error {
	var v map[string]interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	failedToParse := []string{}
	var ok bool

	errs, ok := v["errors"].(map[string]interface{})
	if !ok {
		failedToParse = append(failedToParse, "errors")
	}
	c.ErrorsTotal = float64(len(errs))

	lr, ok := v["last_restarted"].(string)
	if !ok {
		failedToParse = append(failedToParse, "last_restarted")
	}
	lrt, err := time.Parse(time.ANSIC, lr)
	if err != nil {
		log.Printf("cannot parse Last Restart '%v' time to time.Time: %v", lr, err)
	}
	c.LastRestarted = float64(lrt.Unix())

	c.ServiceStatus = make(map[string]float64)

	ss, ok := v["service_status"].(map[string]interface{})
	if !ok {
		failedToParse = append(failedToParse, "service_status")
	}
	for k, v := range ss {
		if v.(string) == "on" {
			c.ServiceStatus[k] = 1
		} else {
			c.ServiceStatus[k] = 0
		}
	}
	if len(failedToParse) > 0 {
		log.Printf("cannot get '%v' from the sacli output: %v", failedToParse, string(data))
	}

	return nil
}
