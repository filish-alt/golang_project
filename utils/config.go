package utils

import (
	"context"
	"database/sql"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ENVIRONMENT          string        `mapstructur:"ENVIRONMENT"`
	DBDriver             string        `mapstructur:"DB_DRIVER"`
	DBSource             string        `mapstructur:"DB_SOURCE"`
	RedisAddress         string        `mapstructure:"REDIS_ADDRESS"`
	MigrationURL         string        `mapstructure:"MIGRATION_URL"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	HTTPServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
    EmailSenderName      string        `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress   string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword  string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
 }

// ExecContext implements db.DBTX.
func (c Config) ExecContext(context.Context, string, ...interface{}) (sql.Result, error) {
	panic("unimplemented")
}

// PrepareContext implements db.DBTX.
func (c Config) PrepareContext(context.Context, string) (*sql.Stmt, error) {
	panic("unimplemented")
}

// QueryContext implements db.DBTX.
func (c Config) QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error) {
	panic("unimplemented")
}

// QueryRowContext implements db.DBTX.
func (c Config) QueryRowContext(context.Context, string, ...interface{}) *sql.Row {
	panic("unimplemented")
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
