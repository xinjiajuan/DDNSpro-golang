package yamlconfig_ddnspro

import (
	"dnspro/Run/yamlConfig/struct"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func ReadYamlConfig(configPath string) structConstant.DDNS_config {
	var conf structConstant.DDNS_config
	yamlfile, err := ioutil.ReadFile(configPath)
	if err != nil {
		println(err)
		return conf
	}
	if err = yaml.Unmarshal(yamlfile, &conf); err != nil {
		println(err)
		return conf
	}
	return conf
}
func CreateDomainList(config structConstant.DDNS_config) []structConstant.AllDomainList {
	domainlist := []structConstant.AllDomainList{}
	for i, user := range config.UserConfig {
		if user.ServiceProvider < 0 || user.ServiceProvider > 2 {
			fmt.Printf("错误:第%d个账号的服务商编号配置输入错误!\a", i+1)
			return nil
		}
		for _, domain := range user.DomainList {
			domainA := structConstant.AllDomainList{}
			domainA.Domain = domain
			domainA.ServiceProvider = user.ServiceProvider
			domainA.Token = getservicetoken(user.ServiceProvider, user)
			domainlist = append(domainlist, domainA)
		}
	}
	return domainlist
}
func getservicetoken(pro int, user structConstant.User) []string {
	var token []string
	if pro == 0 {
		token = append(token, user.Dnspod.TenSecretID)
		token = append(token, user.Dnspod.TenSecretKey)
		return token
	} else if pro == 1 {
		token = append(token, user.Cloudflare.AuthKey)
		return token
	}
	return nil
}
