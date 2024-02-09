package conf

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func ReadConf() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// 读取配置
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	pflag.String("app.cacheDir", "./cache/", "cache directory")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine) // 绑定命令行参数
}

func GetValue(key, defaultValue string) string {
	if viper.GetString(key) == "" {
		return defaultValue
	}
	return viper.GetString(key)
}

func GetIntValue(key string, defaultValue int) int {
	if viper.GetInt(key) == 0 {
		return defaultValue
	}
	return viper.GetInt(key)
}
