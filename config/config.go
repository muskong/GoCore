package config

import (
	"sync"

	"github.com/spf13/viper"
)

var (
	Db  *viper.Viper
	Log *viper.Viper
	App *viper.Viper
	Rsa *viper.Viper

	configOnce sync.Once
)

func Init() {
	configOnce.Do(func() {
		Db = load("db")
		Log = load("log")
		App = load("app")
		Rsa = load("rsa")
	})
}

func load(configureName string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigType("toml")
	conf.SetConfigName(configureName)
	conf.AddConfigPath("./conf")
	err := conf.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(err)
	}
	return conf
}
