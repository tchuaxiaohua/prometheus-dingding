package notify

import (
	"bytes"
	"github.com/gin-gonic/gin"
)

var Service *Sender

// 告警发送接口 所有告警渠道需要实现该接口
type Sender interface {
	Send(*gin.Context,*bytes.Buffer ) error
}
