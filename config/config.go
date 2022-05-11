package config

// 全局配置文件
func NewConfig() *Config {
	return &Config{
		App:      newApp(),
		Dingding: map[string]Dingtoken{},
	}
}

// 全局配置 对象
type Config struct {
	App *app `toml:"app"`
	Dingding map[string]Dingtoken `toml:"dingding"`
}

// 应用端口相关信息
type app struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

// app 构造函数
func newApp() *app {
	return &app{
		Host: "0.0.0.0",
		Port: "8080",
	}
}

func NewDingToken() *Dingtoken {
	return &Dingtoken{}
}
// 钉钉token 对象
type Dingtoken struct {
	Token string
	Name string
}