package substatusparser

import (
	"encoding/json"
	"strings"

	"github.com/cyrilit69/openvpnas_exporter/models"
)

func Parse(in []byte) (*models.SubscriptionStatus, error) {
	s := string(in)
	s = strings.ReplaceAll(s, "'", "\"")
	s = strings.ReplaceAll(s, "True", "true")
	s = strings.ReplaceAll(s, "False", "false")
	s = strings.ReplaceAll(s, "None", "null")
	in = []byte(s)
	res := new(models.SubscriptionStatus)
	err := json.Unmarshal(in, res)
	if err != nil {
		return res, err
	}
	return res, nil
}
