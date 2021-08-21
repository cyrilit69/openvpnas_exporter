package substatusparser

import (
	"encoding/json"

	"github.com/cyrilit69/openvpnas_exporter/models"
)

func Parse(in []byte) (*models.SubscriptionStatus, error) {
	res := new(models.SubscriptionStatus)
	err := json.Unmarshal(in, res)
	if err != nil {
		return res, err
	}
	return res, nil
}
