package infrastructure

import (
	"log"

	"go.uber.org/zap"
)

func NewLogger() *zap.SugaredLogger {

	encoder := zap.NewDevelopmentEncoderConfig()
	logLevel := zap.NewAtomicLevelAt(zap.InfoLevel)

	config := zap.Config{
		Level:            logLevel,
		Development:      false,
		Encoding:         "json",
		EncoderConfig:    encoder,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := config.Build() //zap.NewProduction()

	if err != nil {
		log.Fatalf("Erro ao inicializar o logger: %v", err)
	}

	defer logger.Sync() // flushes buffer, if any

	return logger.Sugar() //, logExec.Sugar()
}
