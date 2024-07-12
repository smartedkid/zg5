package config

type NacosConfig struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Namespace string `json:"namespace"`
	DataId    string `json:"data_id"`
	Group     string `json:"group"`
}

type ServerConfig struct {
	Name  string   `json:"name"`
	Host  string   `json:"host"`
	Tags  []string `json:"tags"`
	Mysql struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Db       string `json:"db"`
	} `json:"mysql"`
	Es struct {
		Hort string `json:"hort"`
		Port int    `json:"port"`
	} `json:"es"`
	Redis struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Password string `json:"password"`
		Db       int    `json:"db"`
	} `json:"redis"`
	Consul struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"consul"`
}
