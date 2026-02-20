package infrastructure

import (
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Env struct {
	DbHost     string `mapstructure:"DB_SERVER_PROD"`
	DbPort     int    `mapstructure:"DB_PORT_PROD"`
	DbSid      string `mapstructure:"DB_SERVICE_PROD"`
	DbUser     string `mapstructure:"DB_USER_PROD"`
	DbPassword string `mapstructure:"DB_PASS_PROD"`
	Port       int    `mapstructure:"PORT"`
	Host       string `mapstructure:"HOST"`
	BasePath   string `mapstructure:"BASE_PATH"`
}

func NewEnv(envPath string, logger *zap.SugaredLogger) *Env {
	logger.Info("Iniciando leitura de variáveis de ambiente.")
	defer logger.Info("Leitura das variáveis de ambiente completada com sucesso.")

	env := Env{}

	v := viper.New()
	v.SetConfigFile(envPath)

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		logger.Infow("Não foi possível ler algumas variáveis de ambiente.", "erro", err)
	}

	err = v.Unmarshal(&env)
	if err != nil {
		logger.Fatalw("Erro ao fazer unmarshal das variáveis de ambiente.", "erro", err)
	}

	return &env
}
