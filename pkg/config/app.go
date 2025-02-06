package config

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

type AppConfig struct {
	App struct {
		Port   string `mapstructure:"port"`
		Domain string `mapstructure:"domain"`
	} `mapstructure:"app"`
	Log struct {
		Directory string `mapstructure:"directory"`
		Mysql     string `mapstructure:"mysql"`
	} `mapstructure:"log"`
	Jwt struct {
		Key      string `mapstructure:"key"`
		Issuer   string `mapstructure:"issuer"`
		Duration string `mapstructure:"duration"`
		Audience string `mapstructure:"audience"`
	} `mapstructure:"jwt"`
	Mysql struct {
		Host               string `mapstructure:"host"`
		Database           string `mapstructure:"database"`
		Password           string `mapstructure:"password"`
		Port               string `mapstructure:"port"`
		User               string `mapstructure:"user"`
		MigrationDirectory string `mapstructure:"migration_directory"`
	} `mapstructure:"mysql"`
	Redis struct {
		MaxIdle   string `mapstructure:"max_idle"`
		MaxActive string `mapstructure:"max_active"`
		Address   string `mapstructure:"address"`
		Port      string `mapstructure:"port"`
		Password  string `mapstructure:"password"`
	} `mapstructure:"redis"`
	Mongodb struct {
		Address  string `mapstructure:"address"`
		Database string `mapstructure:"database"`
		Port     string `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
	} `mapstructure:"mongodb"`
	APM struct {
		ServiceName string `mapstructure:"service_name"`
		Version     string `mapstructure:"version"`
		Url         string `mapstructure:"url"`
		SecretToken string `mapstructure:"secret_token"`
	} `mapstructure:"apm"`
	Kafka struct {
		Version           string `mapstructure:"version"`
		Brokers           string `mapstructure:"brokers"`
		ClientID          string `mapstructure:"client_id"`
		ChannelBufferSize string `mapstructure:"channel_buffer_size"`
		Sasl              struct {
			Enabled   string `mapstructure:"enabled"`
			User      string `mapstructure:"user"`
			Password  string `mapstructure:"password"`
			Mechanism string `mapstructure:"mechanism"`
			Version   string `mapstructure:"version"`
			Handshake string `mapstructure:"handshake"`
		} `mapstructure:"sasl"`
		Tls struct {
			Enabled    string `mapstructure:"enabled"`
			CaFile     string `mapstructure:"ca_file"`
			CertFile   string `mapstructure:"cert_file"`
			KeyFile    string `mapstructure:"key_file"`
			SkipVerify string `mapstructure:"skip_verify"`
		}
		Consumer struct {
			SessionTimeout     string `mapstructure:"session_timeout"`
			RebalancedStrategy string `mapstructure:"rebalanced_strategy"`
			OffsetInitial      string `mapstructure:"offset_initial"`
			IsolationLevel     string `mapstructure:"isolation_level"`
			HeartbeatInterval  string `mapstructure:"heartbeat_interval"`
		} `mapstructure:"consumer"`
		Producer struct {
			Timeout           string `mapstructure:"timeout"`
			RequiredAck       string `mapstructure:"required_ack"`
			IdemPotent        string `mapstructure:"idempotent"`
			PartitionStrategy string `mapstructure:"partition_strategy"`
		} `mapstructure:"producer"`
	}
}

func NewAppConfig() (app AppConfig) {
	path := os.Getenv("CONFIG_PATH")
	osEnv := os.Getenv("OS_ENV")

	env := "env"
	if osEnv != "" {
		env = osEnv
	}

	replacer := strings.NewReplacer(`.`, `_`)
	viper.AddConfigPath(path)
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType(`yaml`)
	viper.SetConfigFile(env + `.yaml`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&app)
	if err != nil {
		panic(err)
	}

	return
}
