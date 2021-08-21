package vpnstatusparser

import (
	"testing"
)

var validInput string = `
{
	"openvpn_0": {
	  "client_list": [
		[
		  "vpupkin", 
		  "32.1.1.2:61173", 
		  "10.0.0.69", 
		  "", 
		  "46473652", 
		  "117695232", 
		  "Wed Aug 18 06:30:16 2021", 
		  "1629268216", 
		  "vpupkin", 
		  "10872", 
		  "0"
		], 
		[
		  "ccyrilkin", 
		  "8.8.8.9:60094", 
		  "10.0.0.62", 
		  "", 
		  "11332614", 
		  "7056514", 
		  "Wed Aug 18 08:33:29 2021", 
		  "1629275609", 
		  "ccyrilkin", 
		  "10899", 
		  "0"
		]
	  ], 
	  "global_stats": {
		"Max bcast/mcast queue length": "10"
	  }, 
	  "time": [
		"Wed Aug 18 08:49:02 2021", 
		"1629276542"
	  ], 
	  "title": "OpenVPN 2.4.9as2 x86_64-pc-linux-gnu [SSL (OpenSSL)] [LZO] [LZ4] [EPOLL] [MH/PKTINFO] [AEAD] built on Apr 13 2021"
	}, 
	"openvpn_1": {
	  "client_list": [
		[
		  "ipupkin", 
		  "17.18.19.20:51979", 
		  "10.0.0.206", 
		  "", 
		  "14722394", 
		  "12662809", 
		  "Wed Aug 18 07:46:28 2021", 
		  "1629272788", 
		  "ipupkin", 
		  "37960", 
		  "119"
		], 
		[
		  "alkonost", 
		  "77.7.7.7:55493", 
		  "10.0.0.81", 
		  "", 
		  "7196204", 
		  "8116344", 
		  "Wed Aug 18 07:54:01 2021", 
		  "1629273241", 
		  "alkonost", 
		  "37972", 
		  "121"
		]
	  ], 
	  "global_stats": {
		"Max bcast/mcast queue length": "193"
	  }, 
	  "time": [
		"Wed Aug 18 08:49:02 2021", 
		"1629276542"
	  ], 
	  "title": "OpenVPN 2.4.9as2 x86_64-pc-linux-gnu [SSL (OpenSSL)] [LZO] [LZ4] [EPOLL] [MH/PKTINFO] [AEAD] built on Apr 13 2021"
	}
  }`

func TestValidInput(t *testing.T) {
	in := []byte(validInput)
	res, err := Parse(in)
	if err != nil {
		t.Fail()
	}
	for _, i := range res {
		t.Logf("Client %v has IP %v", i.ClientName, i.RealAddr)
	}
	if res[0].ClientName != "vpupkin" {
		t.Errorf("Wrong username '%v', expected: 'vpupkin'", res[0].ClientName)
	}
}
