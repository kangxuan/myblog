package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Mode         string `mapstructure:"mode"`
	Port         int    `mapstructure:"port"`
	PageSize     int    `mapstructure:"page_size"`
	ReadTimeOut  int    `mapstructure:"read_time_out"`
	WriteTimeOut int    `mapstructure:"write_time_out"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

var ServerConf = &ServerConfig{}

type MysqlConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Pass        string `mapstructure:"pass"`
	Db          string `mapstructure:"db"`
	MaxOpenConn int    `mapstructure:"max_open_conn"`
	MaxIdleConn int    `mapstructure:"max_idle_conn"`
}

type RedisConfig struct {
	Host          string `mapstructure:"host"`
	Port          int    `mapstructure:"port"`
	Pass          string `mapstructure:"pass"`
	Db            int    `mapstructure:"db"`
	MaxIdleConn   int    `mapstructure:"max_idle_conn"`
	MaxActiveConn int    `mapstructure:"max_active_conn"`
}

func SetUp() {
	// viper设置配置文件路径是以根目录为准
	viper.SetConfigFile("./conf/app.yaml")

	// 观察配置文件是否发生变化
	viper.WatchConfig()

	// 配置文件发生变化后要同步到全局变量Conf
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(ServerConf); err != nil {
			panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
		}
	})

	// 加载配置
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("ReadInConfig failed, err: %v", err))
	}
	// 解析配置
	if err := viper.Unmarshal(ServerConf); err != nil {
		panic(fmt.Errorf("unmarshal to Conf failed, err:%v", err))
	}
}
