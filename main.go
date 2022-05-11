package main

import (
	"dingding-alert/cmd"
	"fmt"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Printf("启动失败:%s", err)
		os.Exit(1)
	}
}
