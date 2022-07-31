package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/lemonlzy/vegetableBlog/configs"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name      string `toml:"name"`
	Mode      string `toml:"mode"`
	Port      string `toml:"port"`
	Version   string `toml:"version"`
	SinceTime string `toml:"since_time"`

	*LogConfig   `toml:"log"`
	*DBConfig    `toml:"db"`
	*CacheConfig `toml:"cache"`
}

type LogConfig struct {
	Level      string `toml:"level"`
	Filename   string `toml:"filename"`
	MaxSize    int    `toml:"max_size"`
	MaxAge     int    `toml:"max_age"`
	MaxBackups int    `toml:"max_backups"`
}

type DBConfig struct {
	Host         string `toml:"host"`
	User         string `toml:"user"`
	Password     string `toml:"password"`
	DBName       string `toml:"db_name"`
	Port         int    `toml:"port"`
	MaxOpenConns int    `toml:"max_open_conns"`
	MaxIdleConns int    `toml:"max_idle_conns"`
	TablePrefix  string `toml:"table_prefix"`
	MigrateTable bool   `toml:"migrate_table"`
}

type CacheConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Password string `toml:"password"`
	DB       int    `toml:"db"`
	PoolSize string `toml:"pool_size"`
}

func Init() error {
	_, err := toml.DecodeFS(configs.EmbedConf, "conf.toml", &Conf)
	if err != nil {
		return err
	}
	return nil
}
