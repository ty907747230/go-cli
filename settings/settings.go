package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigFile("../config.yaml") // 指定配置文件路径
	viper.SetConfigName("config")         //指定配置文件名称
	viper.SetConfigType("yaml")           //指定配置文件类型
	viper.AddConfigPath(".")              //指定配置文件的路径
	err = viper.ReadInConfig()            //读取配置

	if err != nil {
		fmt.Sprintf("viper failed,err:%v\n", err)
		//panic(fmt.Errorf("Fatal error config file:%s \n", err))
	}
	//r := gin.Default()
	//if err = r.Run(fmt.Sprintf(":%d", viper.GetInt("app.port"))); err != nil {
	//	fmt.Printf("%v111", err)
	//	//panic(err)
	//}
	viper.WatchConfig() //监听配置文件修改
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
	})
	return
}
