package config

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

type Config interface {
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	Init(string)
}

type viperConfig struct{}

func (v *viperConfig) Init(prefix string) {
	viper.SetEnvPrefix(`go-clean`)
	viper.AutomaticEnv()

	osEnv := os.Getenv("OS_ENV")

	env := "env"
	if osEnv != "" {
		env = osEnv
	}

	if prefix != "" {
		env = prefix + "." + env
	}

	replacer := strings.NewReplacer(`.`, `_`)
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType(`yaml`)
	viper.SetConfigFile(env + `.yaml`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
}

func (v *viperConfig) GetString(key string) string {
	return viper.GetString(key)
}

func (v *viperConfig) GetStringMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}

func NewConfig() Config {
	v := &viperConfig{}
	v.Init("")
	return v
}
