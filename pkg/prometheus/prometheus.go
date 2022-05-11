package prometheus

import (
	"time"
)

// 告警结构体构造函数
func NewMsgInfo() *MsgInfo {
	return &MsgInfo{
		Receiver: "",
		Status:   "",
		Alerts:   nil,
	}
}

// 告警信息 结构体
type MsgInfo struct {
	Receiver string     `json:"receiver"`
	Status   string     `json:"status"`
	Alerts   []AlertMsg `json:"alerts"`
}

// alerts 信息内容 对象
type AlertMsg struct {
	Status       string                 `json:"status"`
	Labels       map[string]string 		`json:"labels"`
	Annotations  AnnotaTions            `json:"annotations"`
	Startsat     string                 `json:"startsAt"`
	Endsat       string                 `json:"endsAt"`
	Generatorurl string                 `json:"generatorURL"`
	Fingerprint  string                 `json:"fingerprint"`
}

// 时间转换 UTC --> CST
// 模式是UTC，需要转化为CST时间
func (a *AlertMsg) TimeFormat() {
	// 告警触发时间
	layout := "2006-01-02 15:04:05"
	//loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.Parse(time.RFC3339, a.Startsat)
	a.Startsat = t.In(time.Local).Format(layout)
	// 告警结束时间
	if len(a.Endsat) > 0 {
		t1, _ := time.Parse(time.RFC3339, a.Endsat)
		a.Endsat =  t1.In(time.Local).Format(layout)
	}
}


// annotations 告警信息 对象
type AnnotaTions struct {
	Description string `json:"description"`
	RunbookUrl  string `json:"runbook_url"`
	Summary     string `json:"summary"`
}