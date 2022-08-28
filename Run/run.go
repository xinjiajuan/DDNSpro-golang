package DDNS_Run

import (
	structConstant "dnspro/Run/yamlConfig/struct"
	"fmt"
	"time"
)

func RunTicker(serialNumber int, domain structConstant.AllDomainList) {
	duration, er := time.ParseDuration(domain.Domain.SyncTime)
	if er != nil {
		fmt.Printf("配置文件顺序中第%d个域名的同步IP地址时间解析错误!%s", serialNumber+1, er.Error())
		return
	}
	ticker := time.NewTicker(duration)
	for {
		<-ticker.C
		println(666)
	}

}
