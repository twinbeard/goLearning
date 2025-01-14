package setting

type Config struct {
	Server ServerSetting `mapstructure:"server"`
	Logger LoggerSetting `mapstructure:"logger"`
	Mysql  MySqlSetting  `mapstructure:"mysql"`
	Redis  RedisSetting  `mapstructure:"redis"`
}

type ServerSetting struct {
	Mode string `mapstructure:"mode"`
	Port int    `mapstructure:"port"`
}

type LoggerSetting struct {
	Log_level     string `mapstructure:"log_Level"`
	File_log_name string `mapstructure:"file_log_name"`
	Max_size      int    `mapstructure:"max_size"`
	Max_backups   int    `mapstructure:"max_backups"`
	Max_Age       int    `mapstructure:"max_age"`
	Compress      bool   `mapstructure:"compress"`
}

type MySqlSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}
type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}
