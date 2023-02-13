package config

import (
	"fmt"
	"gorm.io/gorm/logger"
	"os"
	"strconv"
	"time"
)

type (
	// Config project config
	Config struct {
		Consul   Consul           `mapstructure:",squash"`
		HTTP     HTTP             `mapstructure:",squash"`
		Postgres PostgresDatabase `mapstructure:",squash"`
		Mysql    MysqlDatabase    `mapstructure:",squash"`
		MongoDB  MongoDBConfig    `mapstructure:",squash"`
		Kafka    Kafka            `mapstructure:",squash"`
	}

	// Consul contains consul remote config related values
	Consul struct {
		Host  string `mapstructure:"CONSUL_HOST"`
		Port  int    `mapstructure:"CONSUL_PORT"`
		Token string `mapstructure:"CONSUL_HTTP_TOKEN"`
	}

	// HTTP contains HTTP related configuration
	HTTP struct {
		Port             int           `mapstructure:"HTTP_PORT"`
		GracefulDuration time.Duration `mapstructure:"HTTP_GRACEFUL DURATION"`
	}

	// PostgresDatabase contains postgres configuration
	PostgresDatabase struct {
		DSN                   string          `mapstructure:"DB_POSTGRES_DSN"`
		LogLevel              logger.LogLevel `mapstructure:"DB_POSTGRES_LOG_LEVEL"`
		MaxOpenConnections    int             `mapstructure:"DB_POSTGRES_MAX_OPEN_CONNECTION"`
		MaxIdleConnections    int             `mapstructure:"DB_POSTGRES_MAX_IDLE_CONNECTION"`
		MaxConnectionLifeTime time.Duration   `mapstructure:"DB_POSTGRES_MAX_CONNECTION_LIFETIME"`
	}

	// MysqlDatabase contains mysql configuration
	MysqlDatabase struct {
		DSN                   string          `mapstructure:"DB_MYSQL_DSN"`
		LogLevel              logger.LogLevel `mapstructure:"DB_MYSQL_LOG_LEVEL"`
		MaxOpenConnections    int             `mapstructure:"DB_MYSQL_MAX_OPEN_CONNECTION"`
		MaxIdleConnections    int             `mapstructure:"DB_MYSQL_MAX_IDLE_CONNECTION"`
		MaxConnectionLifeTime time.Duration   `mapstructure:"DB_MYSQL_MAX_CONNECTION_LIFETIME"`
	}

	// MongoDBConfig contains mongoDB configuration
	MongoDBConfig struct {
		MongoUri      string `mapstructure:"MONGO_DB_HOST"`
		MongoDatabase string `mapstructure:"MONGO_DB_NAME"`
	}

	// Kafka contains Kafka related configuration
	Kafka struct {
		Brokers string `mapstructure:"KAFKA_BROKERS"`
		Group   string `mapstructure:"KAFKA_CONSUMER_GROUP"`
	}
)

func (c *Config) endpoint() string {
	if os.Getenv("CONSUL_HTTP_TOKEN") == "" && c.Consul.Token != "" {
		_ = os.Setenv("CONSUL_HTTP_TOKEN", c.Consul.Token)
	}

	if c.Consul.Host == "" {
		c.Consul.Host = os.Getenv("CONSUL_HOST")
	}

	if c.Consul.Port == 0 {
		consulPort := os.Getenv("CONSUL_PORT")
		if consulPort != "" {
			if port, err := strconv.Atoi(consulPort); err != nil {
				c.Consul.Port = port
			}
		}
	}

	return fmt.Sprintf("%s:%d", c.Consul.Host, c.Consul.Port)
}

func (c Config) IsValidHTTP() bool {
	return c.HTTP.Port != 0
}
