package cmd

import (
	"dingding-alert/config"
	"dingding-alert/pkg/notify/dingding"
	"dingding-alert/pkg/prometheus"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func SendDing(c *gin.Context) {
	// 初始化 Prometheus实例
	msgObj := prometheus.NewMsgInfo()
	// 初始化钉钉实例
	send := dingding.NewDingInfo(c)
	fmt.Printf("告警渠道:%s\n", send.AppName)
	// 告警信息 解析
	if err := c.Bind(msgObj); err != nil {
		fmt.Printf("数据绑定失败:%s", err)
	}
	// 处理告警分组
	for _, msg := range msgObj.Alerts {
		msg.TimeFormat()
		data, _ := prometheus.ParseTemplate("template/alert.tmpl", msg)
		err := send.Send(c, data)
		if err != nil {
			fmt.Printf("告警发送失败", err)
		}
	}
}

var RootCmd = &cobra.Command{
	Use:   "promentheus-dingtalk",
	Short: "prometheus 告警渠道管理",
	Long:  "prometheus 告警渠道管理",
	RunE: func(cmd *cobra.Command, args []string) error {

		// 1 初始话全局配置
		if err := config.LoadConfigFromToml(configName); err != nil {
			fmt.Println("配置文件初始话失败", err)
			return err
		}
		// 2 参数判断
		ParamsInit()
		// 3 gin 启动
		g := gin.Default()
		g.POST("/dingding/:channle-type", SendDing)
		g.Run(":" + appPort)
		return fmt.Errorf("没有参数")
	},
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&configName, "conf", "f", configName, "项目启动配置文件")
	RootCmd.PersistentFlags().StringVarP(&appPort, "port", "p", appPort, "项目监听端口")
	RootCmd.PersistentFlags().BoolVarP(&help, "help", "h", false, "帮助信息")
}
