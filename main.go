package main

import (
	"flag"
	args "github.com/alexflint/go-arg"
	"github.com/kardianos/service"
	"time"
)

type Args struct {
	ConfigPath           string   `arg:"-c,--config" help:"配置文件路径"`
	DomainList           []string `arg:"-d,--domain" help:"域名列表"`
	IDServiceProvider    string   `arg:"-i,--serviceID" help:"服务商用于身份验证的身份标识或编号"`
	TokenServiceProvider string   `arg:"-k,--serviceToken" help:"服务商用于身份验证的密钥或令牌"`
	RecordType           []string `arg:"-r,--recordtype" help:"域名记录的类型,请与domain参数顺序对齐"`
	ServiceProvider      []string `arg:"-s,--serviceprovider" help:"服务商名称, dnspod/aliyun/huawei/cf"`
}

var logger service.Logger

// Program structures.
//  Define Start and Stop methods.
type program struct {
	exit chan struct{}
}

func (p *program) Start(s service.Service) error {
	if service.Interactive() {
		logger.Info("Running in terminal.")
	} else {
		logger.Info("Running under service manager.")
	}
	p.exit = make(chan struct{})

	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() error {
	logger.Infof("I'm running %v.", service.Platform())
	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case tm := <-ticker.C:
			logger.Infof("Still running at %v...", tm)
		case <-p.exit:
			ticker.Stop()
			return nil
		}
	}
}
func (p *program) Stop(s service.Service) error {
	// Any work in Stop should be quick, usually a few seconds at most.
	logger.Info("I'm Stopping!")
	close(p.exit)
	return nil
}

func main() {
	arg := Args{}
	args.MustParse(&arg)
	svcFlag := flag.String("service", "", "Control the system service")
	flag.Parse()
	options := make(service.KeyValue)
	options["Restart"] = "on-success"
	options["SuccessExitStatus"] = "1 2 8 SIGKILL"
	svcConfig := &service.Config{
		Name:        "DDNSPro Golang",
		DisplayName: "DDNS Service for golang",
		Description: "动态域名程序,将你本机IP地址与你的域名记录同步",
		Dependencies: []string{
			"Requires=network.target",
			"After=network-online.target syslog.target"},
	}
	//yamlconfig_ddnspro.ReadYamlConfig()
	//dnspod_tencentcloud_sdk_api_ddnspro.GetDNSPODRecordList_TencloudSDK()
}
