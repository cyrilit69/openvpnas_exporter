/*
	Someone could kill me for this code
*/

package statusparser

import (
	"encoding/json"

	"github.com/cyrilit69/openvpnas_exporter/models"
)

/*
The input for each client will be like this, but in the list format:
   {
0       "Common Name": "ipupkin",
1       "Real Address": "17.18.19.20:51979",
2       "Virtual Address": "10.0.0.206",
3       "Virtual IPv6 Address": "",
4       "Bytes Received": "14722394",
5       "Bytes Sent": "12662809",
6       "Connected Since": "Wed Aug 18 07:46:28 2021",
7       "Connected Since (time_t)""1629272788",
8       "Username": "ipupkin",
9       "Client ID": "37960",
10      "Peer ID": "119"
   }
*/

type raw struct {
	Data map[string]vpnData
}

type vpnData struct {
	Clients []models.VPNStatus `json:"client_list"`
}

func (r *raw) UnmarshalJSON(data []byte) error {
	r.Data = make(map[string]vpnData)
	var data_v map[string]vpnData
	if err := json.Unmarshal(data, &data_v); err != nil {
		return err
	}
	for k, v := range data_v {
		r.Data[k] = v
	}
	return nil
}

func Parse(in []byte) ([]*models.VPNStatus, error) {
	resArr := make([]*models.VPNStatus, 0)
	r := new(raw)

	err := json.Unmarshal(in, r)
	if err != nil {
		return resArr, err
	}

	for k, v := range r.Data {
		for i := range v.Clients {
			v.Clients[i].ClientVPN = k
			resArr = append(resArr, &v.Clients[i])
		}
	}
	return resArr, nil
}
