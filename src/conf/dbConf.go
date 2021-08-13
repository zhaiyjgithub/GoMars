package conf

type DBConf struct {
	Host string
	Port int
	User string
	Password string
	DBName string
}

var RedisConf = DBConf{
	Host:     "redis",
	Port:     6379,
	User:     "",
	Password: "",
	DBName:   "",
}

