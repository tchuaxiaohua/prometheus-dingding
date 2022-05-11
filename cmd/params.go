package cmd

import (
	"dingding-alert/config"
)

// 定义所有 启动参数
var (
	configName string = "etc/app.toml"
	appPort    string
	appHost    string
	help       bool
)

func ParamsInit() {
	// 1 初始话全局配置
	//if err := config.LoadConfigFromToml(configName); err != nil {
	//	fmt.Println("配置文件初始话失败",err)
	//}

	if appPort == "" {
		appPort = config.C().App.Port
	}
	if appHost == "" {
		appHost = config.C().App.Host
	}
}
