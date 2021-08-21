package models

/*
	Package contains structs for the metrics building. Fields could be:
	* float64 (just a value)
	* string or []string (will be a label and its value will be '1')
	* map[string]float64 (label + value)
*/

import (
	"encoding/json"
	"strconv"
	"strings"
)

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

func (c *VPNStatus) UnmarshalJSON(data []byte) error {
	var v []interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	c.ClientName = v[0].(string)
	c.ClientId = v[9].(string)
	c.ClientPeerId = v[10].(string)
	c.RealAddr = strings.Split(v[1].(string), ":")[0]
	c.VPNAddr = v[2].(string)
	res, err := strconv.ParseFloat(v[4].(string), 64)
	if err != nil {
		c.BytesReceived = 0
	} else {
		c.BytesReceived = res
	}
	c.BytesSend, err = strconv.ParseFloat(v[5].(string), 64)
	if err != nil {
		c.BytesSend = 0
	}
	c.ConnectedSinceTs, err = strconv.ParseFloat(v[7].(string), 64)
	if err != nil {
		c.ConnectedSinceTs = 0
	}
	return nil
}
