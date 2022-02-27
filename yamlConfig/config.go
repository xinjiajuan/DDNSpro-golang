package yamlconfig_ddnspro

import (
	ddnsConfig "dnspro/yamlConfig/struct"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func getUserConfigDefault() ddnsConfig.Username {
	var defaultConfigUserName ddnsConfig.Username
	defaultConfigUserName.SecretId = ""
	defaultConfigUserName.SecretKey = ""
	defaultConfigUserName.DnspodId = ""
	defaultConfigUserName.DnspodToken = ""
	if defaultConfigUserName.ServiceProvider == "dnspod" {
		defaultConfigUserName.ApiType = "tencent"
	} else {
		defaultConfigUserName.ApiType = ""
	}
	return defaultConfigUserName
}

func getDomainConfigDefault() ddnsConfig.Domain {
	var defaultConfigDomainList ddnsConfig.Domain
	defaultConfigDomainList.SyncTime = "auto"
	defaultConfigDomainList.RecordType = "A"
	return defaultConfigDomainList
}

func ReadYamlConfig(configPath string) ddnsConfig.DDNS_config {
	viper.SetConfigName("config")   // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(configPath) // path to look for the config file in
	//viper.AddConfigPath("$HOME/.appname")  // call multiple times to add many search paths
	//viper.AddConfigPath(".")               // optionally look for config in the working directory
	viper.SetDefault("userConfig", getUserConfigDefault())
	viper.SetDefault("domainList", getDomainConfigDefault())
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("Config file not found: %w \n", err))
		} else {
			panic(fmt.Errorf("Config file was found but another error was produced: %w \n", err))
		}
	}
	c := ddnsConfig.DDNS_config{}
	er := viper.Unmarshal(&c)
	if er != nil {
		log.Fatalf("viper.Unmarshal: %v", er)
	}
	return c
	/*
		c := ddnsConfig.DDNS_config{}
		yaml_byte, io_er := ioutil.ReadFile("ddns.config.yaml")
		if io_er != nil {
			log.Fatalf(`"ioutil.ReadFile("ddns.config.yaml")`+": %v", io_er)
		}
		er := yaml.Unmarshal(yaml_byte, &c)
		if er != nil {
			log.Fatalf("yaml.Unmarshal: %v", er)
		}
		println(c.UserConfig[0].SecretId + "\n" + c.UserConfig[0].SecretKey)
	*/
}
