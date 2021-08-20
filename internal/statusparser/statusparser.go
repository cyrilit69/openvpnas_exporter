/*
	Someone will kill me for this code
*/

package statusparser

import (
	"encoding/json"
	"strconv"
	"strings"

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
	Clients []ClientData `json:"client_list"`
}

type ClientData struct {
	CommonName       string
	BytesReceived    float64
	BytesSend        float64
	ConnectedSinceTs float64
	ClientId         string
	PeerId           string
	RealAddr         string
	VPNAddr          string
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

func (c *ClientData) UnmarshalJSON(data []byte) error {
	var v []interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	c.CommonName = v[0].(string)
	c.ClientId = v[9].(string)
	c.PeerId = v[10].(string)
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

func Parse(in []byte) ([]*models.VPNStatus, error) {
	resArr := make([]*models.VPNStatus, 0)
	r := new(raw)

	err := json.Unmarshal(in, r)
	if err != nil {
		return resArr, err
	}

	for k, v := range r.Data {
		for _, cl := range v.Clients {
			res := new(models.VPNStatus)

			res.ClientVPN = k
			res.ClientName = cl.CommonName
			res.ClientId = cl.ClientId
			res.ClientPeerId = cl.PeerId
			res.RealAddr = cl.RealAddr
			res.VPNAddr = cl.VPNAddr
			res.BytesReceived = cl.BytesReceived
			res.BytesSend = cl.BytesSend
			res.ConnectedSinceTs = cl.ConnectedSinceTs

			resArr = append(resArr, res)
		}
	}
	return resArr, nil
}
