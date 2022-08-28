package DDNS_Run

import (
	dnspod_tenSDK "dnspro/Run/dnspod"
	structConstant "dnspro/Run/yamlConfig/struct"
)

func StartSync(domain structConstant.ADomain) {
	switch domain.ServiceProvider {
	case 0:
		dnspod_tenSDK.StartDnsPodSync(domain)
		break
	}
}
