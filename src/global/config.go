package global

type RabbitMqConfig struct {
	Host        string    `yaml:"host" mapstructure:"host"`
	Port        int32     `yaml:"port" mapstructure:"port"`
	VirtualHost string    `yaml:"virtualHost" mapstructure:"virtualHost"`
	Username    string    `yaml:"username" mapstructure:"username"`
	Password    string    `yaml:"password" mapstructure:"password"`
	ServerOne   ServerOne `yaml:"serverOne" mapstructure:"serverOne"`
}

type ServerOne struct {
	Exchange   string `yaml:"exchange" mapstructure:"exchange"`
	Queue      string `yaml:"queue" mapstructure:"queue"`
	RoutingKey string `yaml:"routingKey" mapstructure:"routingKey"`
}
