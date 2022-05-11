package models

import (
	"fmt"
	"text/template"
	"time"
)


// 告警结构体构造函数
func NewMsgInfo() *MsgInfo1 {
	return &MsgInfo1{
		Receiver: "",
		Status:   "",
		Alerts:   nil,
	}
}

// 告警信息 结构体
type MsgInfo1 struct {
	Receiver string     `json:"receiver"`
	Status   string     `json:"status"`
	Alerts   []alertMsg `json:"alerts"`
}

// 告警发送信息
// 发送的告警信息内容
type alertMsg struct {
	Status       string                 `json:"status"`
	Labels       map[string]interface{} `json:"labels"`
	Annotations  desMsg                 `json:"annotations"`
	Startsat     string                 `json:"startsAt"`
	Endsat       string                 `json:"endsAt"`
	Generatorurl string                 `json:"generatorURL"`
	Fingerprint  string                 `json:"fingerprint"`
}
// 模板解析判断
// 不同告警场景解析模板需要对应
// 告警触发 告警恢复
func (a alertMsg) ParseTemplate() (*template.Template, error) {
	if a.Status == "firing" {
		t, err := template.ParseFiles("./templates/docker/0fire.tmpl")
		if err != nil {
			fmt.Printf("模板文件加载失败%v\n", err)
		}
		return t, err
	}
	if a.Status == "resolved" {
		t, err := template.ParseFiles("./templates/docker/1resove.tmpl")
		if err != nil {
			fmt.Printf("模板文件加载失败%v\n", err)
		}
		return t, err
	}
	return nil, nil
}

// 时间转换 UTC --> CST
// 模式是UTC，需要转化为CST时间
func (a *alertMsg) TimeFormat() {
	// 告警触发时间
	t, _ := time.Parse(time.RFC3339, a.Startsat)
	a.Startsat = t.In(time.Local).String()
	// 告警结束时间
	if len(a.Endsat) > 0 {
		t1, _ := time.Parse(time.RFC3339, a.Endsat)
		a.Endsat = t1.In(time.Local).String()
	}
}

// Labels 告警信息标签
// 暂未使用
type labelMsg struct {
	Alertname  string `json:"alertname"`
	Container  string `json:"container"`
	Namespace  string `json:"namespace"`
	PodName    string `json:"pod"`
	Prometheus string `json:"prometheus"`
	Severity   string `json:"severity"`
}

//description 告警详情
type desMsg struct {
	Description string `json:"description"`
	RunbookUrl  string `json:"runbook_url"`
	Summary     string `json:"summary"`
}
