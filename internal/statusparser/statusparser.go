/*
This package contains a struct for the 'sacli status' output marshaling and
the only Parse([]byte) func to marshal and cast some data into metrics format
*/

package statusparser

import (
	"encoding/json"

	"github.com/cyrilit69/openvpnas_exporter/models"
)

func Parse(in []byte) (*models.Status, error) {
	res := new(models.Status)
	err := json.Unmarshal(in, res)
	if err != nil {
		return res, err
	}
	return res, nil
}
