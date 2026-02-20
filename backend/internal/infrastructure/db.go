package infrastructure

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

func InitDB(env *Env, logger *zap.SugaredLogger) *Database {

	db := Database{}

	dsn := fmt.Sprintf(`host=%s user=%s password=%s  port=%d dbname=%s sslmode=disable`, env.DbHost, env.DbUser, env.DbPassword, env.DbPort, env.DbSid)

	orm, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatalw("Erro ao conectar ao banco de dados.", "erro", err)
	}

	db.Db = orm

	return &db
}
