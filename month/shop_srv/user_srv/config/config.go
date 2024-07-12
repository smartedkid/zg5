package config

type NacosMysqlConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type NacosRedisConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
	DB   int    `json:"db"`
}

type NacosElasticSearchConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type NacosConsulConfig struct {
	Host string   `json:"host"`
	Port string   `json:"port"`
	Tags []string `json:"tags"`
}

type NacosConfig struct {
	Host      string `json:"host"`
	Port      uint64 `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Namespace string `json:"namespace"`
	DataId    string `json:"dataid"`
	Group     string `json:"group"`
}

type ServerConfig struct {
	Host   string                   `json:"host"`
	Name   string                   `json:"name"`
	Mysql  NacosMysqlConfig         `json:"mysql"`
	Redis  NacosRedisConfig         `json:"redis"`
	Es     NacosElasticSearchConfig `json:"elasticsearch"`
	Consul NacosConsulConfig        `json:"consul"`
}
