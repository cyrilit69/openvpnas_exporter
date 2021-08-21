# OpenVPN-AS Prometheus Exporter

This is a pretty raw OpenVPN AS Prometheus Exporter that uses `sacli` util to get metrics.

For now it parses the output of the following commands:

* sacli VPNStatus
* sacli SubscriptionStatus
* sacli status

## Test data

Due to the lack of information about sacli I cannot be sure if exporter can parse any possible outputs right. For example, I cannot find any example of the output with errors and cannot reproduce it. So, shortly:

**I need more test data!**

And I need this data as is (without personal data, ofc). For example, `sacli SubscriptionStatus` output is not a valid JSON, it's Python's dict. I'm almost sure that exporter cannot process some possible cases.

## Usage
CLI:
```
  -sacli-path string
        Path to 'sacli' script (default "/usr/local/openvpn_as/scripts/sacli")
  -web.listen-address string
        Address to listen on for telemetry (default ":9185")
  -web.telemetry-path string
        Path under which to expose metrics (default "/metrics")
```
You should run exporter as a *openvpn* user (or as root). It's required to use `sacli`. For this reason I suggest to run it as systemd unit instead of using docker.

## Metrics
Metric | Labels | Value
--- | --- | ---
openvpnas_vpnstatus_parsed | | 0 or 1
openvpnas_vpnstatus_client_bytes_received | common_name, id, peer_id, real_addr, vpn_addr, vpn| int
openvpnas_vpnstatus_client_bytes_send | common_name, id, peer_id, real_addr, vpn_addr, vpn| int
openvpnas_vpnstatus_client_connected_since | common_name, id, peer_id, real_addr, vpn_addr, vpn| timestamp
openvpnas_vpnstatus_client_info | common_name, id, peer_id, real_addr, vpn_addr, vpn | always 1 (or absent)
openvpnas_vpnstatus_parsed | | 0 or 1
openvpnas_subscription_agent_disabled | | 0 or 1
openvpnas_subscription_cc_limit | | int
openvpnas_subscription_current_cc | | int
openvpnas_subscription_error | | 0 or 1
openvpnas_subscription_fallback_cc | | int
openvpnas_subscription_grace_period | | int
openvpnas_subscription_last_successful_update | | timestamp
openvpnas_subscription_last_successful_update_age | | int
openvpnas_subscription_max_cc | | int
openvpnas_subscription_name | name | 1
openvpnas_subscription_next_update | | timestamp
openvpnas_subscription_next_update_in | | int
openvpnas_subscription_overdraft | | 0 or 1
openvpnas_subscription_server | server | always 1
openvpnas_subscription_state | state | always 1
openvpnas_subscription_type | type | always 1
openvpnas_subscription_updates_failed | | int
openvpnas_subscription_parsed | | 0 or 1
openvpnas_status_errors | | int
openvpnas_status_last_restarted | | timestamp
openvpnas_status_service_state | service | 0 or 1
openvpnas_status_parsed | | 0 or 1

## Comparison with Alternatives
[lfdominguez/openvpn-access-exporter](https://github.com/lfdominguez/openvpn-access-exporter)

It uses `log.db` file and doesn't provide license and services info

All others work with common OpenVPN only :(