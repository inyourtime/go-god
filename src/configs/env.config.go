package configs

type Env struct {
	ServerPort string         `mapstructure:"server_port"`
	Discord    DisCordWebHook `mapstructure:"discord"`
	Database   DataBase       `mapstructure:"database"`
	JwtSecret  string         `mapstructure:"jwt_secret"`
}

type DisCordWebHook struct {
	WebHookID    string `mapstructure:"webhook_id"`
	Token        string `mapstructure:"token"`
	SqlWebHookID string `mapstructure:"sql_webhook_id"`
	SqlToken     string `mapstructure:"sql_token"`
}

type DataBase struct {
	Postgres Postgres `mapstructure:"postgres"`
	Mongo    Mongo    `mapstructure:"mongodb"`
}

type Postgres struct {
	Host       string `mapstructure:"host"`
	User       string `mapstructure:"user"`
	Dbname     string `mapstructure:"dbname"`
	Port       string `mapstructure:"port"`
	Password   string `mapstructure:"password"`
	SearchPath string `mapstructure:"search_path"`
	Sslmode    string `mapstructure:"sslmode"`
}

type Mongo struct {
	Uri string `mapstructure:"uri"`
}
