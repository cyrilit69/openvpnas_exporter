package substatusparser

import (
	"testing"
)

var validInput string = `
{
	"agent_disabled": false,
	"cc_limit": 250,
	"current_cc": 168,
	"error": null,
	"fallback_cc": 2,
	"grace_period": 30,
	"last_successful_update": 1629457445,
	"last_successful_update_age": 65,
	"max_cc": 250,
	"name": "Subscription 1",
	"next_update": 1629457645,
	"next_update_in": 134,
	"notes": [],
	"overdraft": false,
	"server": "asb.sts.openvpn.net",
	"state": "SUBSCRIPTION_OK",
	"type": "-",
	"updates_failed": 0
}`

func TestValidParse(t *testing.T) {
	in := []byte(validInput)
	res, err := Parse(in)
	if err != nil {
		t.Fail()
	}
	t.Log("name", res.Name)
	t.Log("MaxCC", res.MaxCc)
	t.Log("CurrentCC", res.CurrentCc)
	if res.MaxCc != 250 {
		t.Errorf("Wrong MaxCC: %v, expected: 250", res.MaxCc)
	}
}
