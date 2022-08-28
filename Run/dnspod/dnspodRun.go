package dnspod_tenSDK

import structConstant "dnspro/Run/yamlConfig/struct"

func StartDnsPodSync(domain structConstant.ADomain) {
	RecordReq := GetDNSPODRecordList_TencloudSDK(domain.Token[0], domain.Token[1], domain.Domain.Domain, domain.Domain.Subdomain)
	res := RecordReq.Response
	println(res.RecordList)
}
