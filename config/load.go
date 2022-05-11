package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

var (
	Global *Config
)

// 全局单例 保证配置文件只加载一次
func C() *Config {
	if Global == nil {
		//err := LoadConfigFromToml("./etc/app.toml")
		panic("配置文件未加载.....")
	}
	return Global
}

// 从配置文件加载
func LoadConfigFromToml(filePath string) error  {
	// 初始化全局配置文件 用于配置文件解析
	cfg := NewConfig()
	// 配置文件解析
	if _,err := toml.DecodeFile(filePath,cfg);err != nil {
		return fmt.Errorf("配置文件解析失败:%s",err)
	}
	// 解析后配置 赋值给Global
	Global = cfg
	return nil
}

//func init()  {
//	// 全局配置初始化
//	err := LoadConfigFromToml("etc/app.toml")
//	if err != nil {
//		fmt.Println(err)
//	}
//}