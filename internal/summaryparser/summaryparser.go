/*
This package contains a struct for the 'sacli VPNSummary' output marshaling and
the only Parse([]byte) func to marshal and cast some data into metrics format
*/

package summaryparser

import (
	"encoding/json"

	"github.com/cyrilit69/openvpnas_exporter/models"
)

func Parse(in []byte) (*models.VPNSummary, error) {
	res := new(models.VPNSummary)
	err := json.Unmarshal(in, res)
	if err != nil {
		return res, err
	}
	return res, nil
}
