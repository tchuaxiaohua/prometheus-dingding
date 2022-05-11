package config

import (
	"fmt"
	"testing"
)

func TestLoadFIle(t *testing.T) {
	err := LoadConfigFromToml("../etc/app.toml")
	if err != nil {
		fmt.Printf("文件加载失败:%s",err)
	}
}
