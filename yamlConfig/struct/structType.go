package structConstant

type DDNS_config struct {
	UserConfig []Username `yaml:"userConfig"`
}

type Domain struct {
	DomainName string `yaml:"domain"`
	RecordType string `yaml:"recordType"`
	SyncTime   string `yaml:"syncTime"`
}

type Username struct {
	ServiceProvider string   `yaml:"serviceProvider"`
	ApiType         string   `yaml:"api_type"`
	DnspodId        string   `yaml:"dnspod_id"`
	DnspodToken     string   `yaml:"dnspod_token"`
	SecretId        string   `yaml:"tencent_secretId"`
	SecretKey       string   `yaml:"tencent_secretKey"`
	DomainList      []Domain `yaml:"domainList"`
}
