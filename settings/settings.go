package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

//全局定义变量接收配置项
var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}
type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Db       int    `mapstructure:"db"`
	Password string `mapstructure:"password"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	viper.SetConfigFile("./config.yaml") // 指定配置文件路径
	viper.SetConfigName("config")        //指定配置文件名称
	viper.SetConfigType("yaml")          //指定配置文件类型
	viper.AddConfigPath(".")             //指定配置文件的路径,相对位置是相对于执行文件的路径，也就是.exe或者main.go
	err = viper.ReadInConfig()           //读取配置

	if err != nil {
		fmt.Sprintf("viper failed,err:%v\n", err)
		//panic(fmt.Errorf("Fatal error config file:%s \n", err))
	}
	fmt.Printf("conf:%v\n", Conf)
	//读取到的信息反序列化到Conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper Unmarshal failed,err:%v\n", err)
	}
	fmt.Printf("conf:%v\n", Conf)
	//r := gin.Default()
	//if err = r.Run(fmt.Sprintf(":%d", viper.GetInt("app.port"))); err != nil {
	//	fmt.Printf("%v111", err)
	//	//panic(err)
	//}
	viper.WatchConfig() //监听配置文件修改
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		//读取到的信息反序列化到Conf变量中
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper Unmarshal failed,err:%v\n", err)
		}
	})
	return
}
