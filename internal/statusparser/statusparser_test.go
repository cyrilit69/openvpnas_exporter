package statusparser

import (
	"testing"
)

var (
	validInput = `{
    "active_profile": "Default",
    "errors": {},
    "last_restarted": "Tue Aug 17 13:11:58 2021",
    "service_status": {
      "api": "on",
      "auth": "on",
      "bridge": "on",
      "client_query": "on",
      "crl": "on",
      "daemon_pre": "on",
      "db_push": "on",
      "ip6tables_live": "on",
      "ip6tables_openvpn": "on",
      "iptables_live": "on",
      "iptables_openvpn": "on",
      "iptables_web": "on",
      "log": "on",
      "openvpn_0": "on",
      "openvpn_1": "on",
      "subscription": "on",
      "user": "on",
      "web": "on"
    }
  }`

	validInputWithErrors = `{
    "active_profile": "Default",
    "errors": {"somethig": "wrong"},
    "last_restarted": "Tue Aug 17 13:11:58 2021",
    "service_status": {
      "api": "on",
      "auth": "on",
      "bridge": "on",
      "client_query": "on",
      "crl": "on",
      "daemon_pre": "on",
      "db_push": "on",
      "ip6tables_live": "on",
      "ip6tables_openvpn": "on",
      "iptables_live": "on",
      "iptables_openvpn": "on",
      "iptables_web": "on",
      "log": "on",
      "openvpn_0": "on",
      "openvpn_1": "on",
      "subscription": "on",
      "user": "on",
      "web": "on"
    }
  }`
)

func TestValidParse(t *testing.T) {
	in := []byte(validInput)
	res, err := Parse(in)
	if err != nil {
		t.Fail()
	}
	//t.Log("active_profile", res.ActiveProfile)
	t.Log("errors", res.ErrorsTotal)
	t.Log("last_restarted", res.LastRestarted)
	t.Log("service_status", res.ServiceStatus)
	// if res.ActiveProfile != "Default" {
	// 	t.Fail()
	// }
}

func TestValidWithErrorsParse(t *testing.T) {
	in := []byte(validInputWithErrors)
	res, err := Parse(in)
	if err != nil {
		t.Errorf("error from Parse method: %v", err)
	}
	//t.Log("active_profile", res.ActiveProfile)
	t.Log("errors", res.ErrorsTotal)
	t.Log("last_restarted", res.LastRestarted)
	t.Log("service_status", res.ServiceStatus)
	if res.ErrorsTotal == 0 {
		t.Errorf("Errors should be: get %v, expected: something > 0", res.ErrorsTotal)
	}
}
