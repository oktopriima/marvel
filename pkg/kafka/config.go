package kafka

import (
	"github.com/oktopriima/marvel/pkg/config"
	"github.com/oktopriima/marvel/pkg/util"
	"strings"
)

const (
	defaultVersion = "2.1.1"
)

// Config entity of kafka broker
type Config struct {
	// Brokers list of brokers connection hostname or ip address
	Brokers []string `json:"brokers" yaml:"brokers"`
	SASL    SASL     `json:"sasl" yaml:"sasl"`
	// kafka broker Version
	Version  string         `json:"version" yaml:"version"`
	ClientID string         `json:"client_id" yaml:"client_id"`
	Producer ProducerConfig `json:"producer" yaml:"producer"`
	Consumer ConsumerConfig `json:"consumer" yaml:"consumer"`
	TLS      TLS            `json:"tls" yaml:"tls"`
	// The number of events to buffer in internal and external channels. This
	// permits the producer and consumer to continue processing some messages
	// in the background while user code is working, greatly improving throughput.
	// Defaults to 256.
	ChannelBufferSize int `json:"channel_buffer_size" yaml:"channel_buffer_size"`
}

type ProducerConfig struct {
	// The maximum duration the broker will wait the receipt of the number of
	// RequiredActs (defaults to 10 seconds). This is only relevant when
	// RequiredActs are set to WaitForAll or a number > 1. Only supports
	// millisecond resolution, nanoseconds will be truncated. Equivalent to
	// the JVM producer's `request.timeout.ms` setting.
	TimeoutSecond int `json:"timeout_second" yaml:"timeout_second"`
	// RequireACK
	// 0 = NoResponse doesn't send any response; the TCP ACK is all you get.
	// 1 = WaitForLocal waits for only the local commit to succeed before responding.
	// -1 = WaitForAll
	//  waits for all in-sync replicas to commit before responding.
	// The minimum number of in-sync replicas is configured on the broker via
	// the `min.in sync.replicas` configuration key.
	RequireACK int16 `json:"require_ack" yaml:"require_ack"`
	// If enabled, the producer will ensure that exactly one copy of each message is
	// written.
	IdemPotent bool `json:"idem_potent" yaml:"idem_potent"`

	// Generates partitions for choosing the partition to send messages to
	// (defaults to hashing the message key). Similar to the `partitioner.class`
	// setting for the JVM producer.
	PartitionStrategy string `json:"partition_strategy" yaml:"partition_strategy"`
}

type ConsumerConfig struct {
	// Minimum is 10s
	SessionTimeoutSecond int    `json:"session_timeout_second" yaml:"session_timeout_second"`
	OffsetInitial        int64  `json:"offset_initial" yaml:"offset_initial"`
	HeartbeatInterval    int    `json:"heartbeat_interval" yaml:"heartbeat_interval"`
	RebalancedStrategy   string `json:"rebalanced_strategy" yaml:"rebalanced_strategy"`
	AutoCommit           bool   `json:"auto_commit" yaml:"auto_commit"`
	IsolationLevel       int8   `json:"isolation_level" yaml:"isolation_level"`
}

type SASL struct {
	// Whether to use SASL authentication when connecting to the broker
	// (defaults to false).
	Enable bool `json:"enable" yaml:"enable"`
	// SASLMechanism is the name of the enabled SASL mechanism.
	// Possible values: OAUTH BEARER, PLAIN (defaults to PLAIN).
	Mechanism string `json:"mechanism" yaml:"mechanism"`
	// Version is the SASL Protocol Version to use
	// Kafka > 1.x should use V1, except on Azure EventHub, which use V0
	Version int16 `json:"version" yaml:"version"`
	// Whether to send the Kafka SASL handshake first if enabled
	// (defaults to true). You should only set this to false if you're using
	// a non-Kafka SASL proxy.
	Handshake bool `json:"handshake" yaml:"handshake"`
	// User is the authentication identity (authid) to present for
	// SASL/PLAIN or SASL/SCRAM authentication
	User string `json:"user" yaml:"user"`
	// Password for SASL/PLAIN authentication
	Password string `json:"password" yaml:"password"`
	// auth id used for SASL/SCRAM authentication

}

func NewKafkaConfig(c config.AppConfig) *Config {
	return &Config{
		Brokers: strings.Split(c.Kafka.Brokers, ","),
		SASL: SASL{
			Enable:    util.StringToBool(c.Kafka.Sasl.Enabled),
			User:      c.Kafka.Sasl.User,
			Password:  c.Kafka.Sasl.Password,
			Mechanism: c.Kafka.Sasl.Mechanism,
			Version:   int16(util.StringToInt(c.Kafka.Sasl.Version)),
			Handshake: util.StringToBool(c.Kafka.Sasl.Handshake),
		},
		Version:  c.Kafka.Version,
		ClientID: c.Kafka.ClientID,
		Producer: ProducerConfig{
			TimeoutSecond:     util.StringToInt(c.Kafka.Producer.Timeout),
			RequireACK:        int16(util.StringToInt(c.Kafka.Producer.RequiredAck)),
			IdemPotent:        util.StringToBool(c.Kafka.Producer.IdemPotent),
			PartitionStrategy: c.Kafka.Producer.PartitionStrategy,
		},
		Consumer: ConsumerConfig{
			SessionTimeoutSecond: util.StringToInt(c.Kafka.Consumer.SessionTimeout),
			HeartbeatInterval:    util.StringToInt(c.Kafka.Consumer.HeartbeatInterval),
			RebalancedStrategy:   c.Kafka.Consumer.RebalancedStrategy,
			OffsetInitial:        util.StringToInt64(c.Kafka.Consumer.OffsetInitial),
			IsolationLevel:       int8(util.StringToInt(c.Kafka.Consumer.IsolationLevel)),
		},
		TLS: TLS{
			Enable:     util.StringToBool(c.Kafka.Tls.Enabled),
			CaFile:     c.Kafka.Tls.CaFile,
			KeyFile:    c.Kafka.Tls.KeyFile,
			CertFile:   c.Kafka.Tls.CertFile,
			SkipVerify: util.StringToBool(c.Kafka.Tls.SkipVerify),
		},
		ChannelBufferSize: util.StringToInt(c.Kafka.ChannelBufferSize),
	}
}
