package structConstant

//配置文件
type DDNS_config struct {
	UserConfig []User `yaml:"userConfig"`
}

type Domain struct {
	DomainName     string `yaml:"domain"`
	NetworkAdapter string `yaml:"networkAdapter"`
	RecordType     string `yaml:"recordType"`
	SyncTime       string `yaml:"syncTime"`
}

type User struct {
	ServiceProvider int `yaml:"serviceProvider"`
	Dnspod          struct {
		TenSecretID  string `yaml:"tencent_secretId"`
		TenSecretKey string `yaml:"tencent_secretKey"`
	} `yaml:"dnspod"`
	Cloudflare struct {
		AuthKey string `yaml:"authKey"`
	} `yaml:"cloudflare"`
	NetworkAdapter string   `yaml:"networkAdapter"`
	DomainList     []Domain `yaml:"domainList"`
}

//修改版本到此修改
const Version string = "0.1.0"

type AllDomainList struct {
	ServiceProvider int
	Token           []string
	Domain          Domain
}
