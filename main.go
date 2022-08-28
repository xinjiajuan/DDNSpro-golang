package main

import (
	DDNS_Run "dnspro/Run"
	"dnspro/Run/yamlConfig"
	"dnspro/Run/yamlConfig/struct"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

func main() {
	var pathConfig string

	app := &cli.App{
		Name:     "DDNS Pro",
		Version:  structConstant.Version,
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "xinjiajuan",
				Email: "itpours@qq.com",
			},
		},
		Copyright: "(c) 2022 xinjiajuan",
		Usage:     "多服务商动态 IP 地址解析",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "config",
				//Value:       "config.yaml",
				Usage:       "指定YAML配置文件路径",
				Aliases:     []string{"c"},
				Required:    true,
				Destination: &pathConfig,
			},
		},
		Action: func(cCtx *cli.Context) error {
			//判断文件夹是否存在
			_, err := os.Stat(cCtx.String("config")) //os.Stat获取文件信息
			if err != nil {
				return cli.Exit("配置文件不存在!", 86)
			} else {
				config := yamlconfig_ddnspro.ReadYamlConfig(cCtx.String("config"))
				for i, domain := range yamlconfig_ddnspro.CreateDomainList(config) {
					go DDNS_Run.RunTicker(i, domain)
				}
				select {}
			}
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
