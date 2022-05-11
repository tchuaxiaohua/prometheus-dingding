package prometheus

import (
	"bytes"
	"fmt"
	"text/template"
)

// 模板解析
func ParseTemplate(fileName string,data interface{}) (*bytes.Buffer,error) {
	// 保存解析后数据
	var buf bytes.Buffer
	tml,err := template.ParseFiles(fileName)
	if err != nil {
		return nil,fmt.Errorf("模板解析失败：%s",err)
	}

	err = tml.Execute(&buf,data)
	if err != nil {
		return nil,fmt.Errorf("模板渲染失败：%s",err)
	}
	return &buf,nil
}
