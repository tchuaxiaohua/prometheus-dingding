package main

import (
	"fmt"
	"os"

	"github.com/tchuaxiaohua/prometheus-dingding/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Printf("启动失败:%s", err)
		os.Exit(1)
	}
}
