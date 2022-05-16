package dingding

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/tchuaxiaohua/prometheus-dingding/config"

	"github.com/gin-gonic/gin"
)

func NewDingInfo(c *gin.Context) *DingInfo {
	TokenName := c.Param("channle-type")
	if TokenName == "" {
		return nil
	}
	return &DingInfo{
		Token: config.C().Dingding[TokenName].Token,
		AppName: TokenName,
	}
}
//  钉钉配置 对象
type DingInfo struct {
	Token 	string
	AppName string
}

func (d *DingInfo) GetToken() string {
	 d.Token = config.C().Dingding[d.AppName].Token
	 return d.Token
}

// 发送方法
func (d *DingInfo) Send(c *gin.Context,data *bytes.Buffer) error {
	// 告警数据内容
	content := fmt.Sprintf(`{
									"msgtype": "markdown",
									"markdown": {
    									"title": "prometheus_alert",
    									"text": "%s",
    								},
								}`,data.String())
	dingUrl := fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s",d.GetToken())
	// 构造POST请求
	req, err := http.NewRequest("POST", dingUrl,strings.NewReader(content))
	if err != nil {
		fmt.Println("构造请求失败:",err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送失败",err)
	}
	// 钉钉告警response
	io.Copy(os.Stdout,resp.Body)
	defer resp.Body.Close()
	return nil
}