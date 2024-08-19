package config

// Config 读取application.yaml 文件的配置类
type Config struct {
	NacosConfig NacosConfig `yaml:"nacos" mapstructure:"nacos"`
	DataBases   DataBases   `yaml:"databases" mapstructure:"databases"`
	Server      Server      `yaml:"server" mapstructure:"server"`
	Logs        LogConfig   `yaml:"log" mapstructure:"log"`
}

// NacosConfig 结构体 配置了连接Nacos的参数
type NacosConfig struct {
	Dir                 Directory `yaml:"dir" mapstructure:"dir"`
	LogLevel            string    `yaml:"logLevel" mapstructure:"logLevel"`
	Username            string    `yaml:"username" mapstruture:"username"`
	Password            string    `yaml:"password" mapstruture:"password"`
	TimeoutMs           uint64    `yaml:"timeoutMs" mapstruture:"timeoutMs"`
	IP                  string    `yaml:"ip" mapstruture:"ip"`
	Port                uint64    `yaml:"port" mapstruture:"port"`
	NotLoadCacheAtStart bool      `yaml:"notLoadCacheAtStart" mapstructure:"notLoadCacheAtStart"`
	ContextPath         string    `yaml:"contextPath" mapstructure:"contextPath"`
	Bootstrap           BootStrap `yaml:"bootstrap" mapstructure:"bootstrap"`
}

// Directory nacos 缓存和日志文件配置
type Directory struct {
	Log   string `yaml:"log" mapstruture:"log"`
	Cache string `yaml:"cache" mapstruture:"cache"`
}

type BootStrap struct {
	DataId string `yaml:"dataId" mapstruture:"dataId"`
	Group  string `yaml:"group" mapstruture:"group"`
}

type Server struct {
	IP     string `yaml:"ip" mapstructure:"ip"`
	Port   uint64 `yaml:"port" mapstructure:"port"`
	Name   string `yaml:"name" mapstructure:"name"`
	Group  string `yaml:"group" mapstructure:"group"`
	GinMod string `yaml:"ginMod" mapstructure:"ginMod"`
}
type DataBases struct {
	Mysql MySQLConfig `yaml:"mysql" mapstructure:"mysql"`
}

type NacosBootstrapConfig struct {
	DataBases DataBases `yaml:"databases" mapstructure:"databases"`
	Server    Server    `yaml:"server" mapstructure:"server"`
}

type MySQLConfig struct {
	Username string `yaml:"username" mapstructure:"username"`
	Password string `yaml:"password" mapstructure:"password"`
	Ip       string `yaml:"ip" mapstructure:"ip"`
	Port     int32  `yaml:"port" mapstructure:"port"`
	Database string `yaml:"database" mapstructure:"database"`
}

type LogConfig struct {
	Dir        string `yaml:"dir"  mapstructure:"dir"`
	Level      string `yaml:"level" mapstructure:"level"`
	DayFormat  string `yaml:"dayFormat" mapstructure:"dayFormat"`
	TimeFormat string `yaml:"timeFormat" mapstructure:"timeFormat"`
}
