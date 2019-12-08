package pkg

type MysqlConfig struct {
	User     string
	Password string
	Host     string
	Database string
	Prefix   string
	Port     int
}

type Config struct {
	Port  int
	Mysql MysqlConfig
}
