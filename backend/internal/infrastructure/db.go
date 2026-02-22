package infrastructure

import (
	"fmt"
	"real-estate-api/internal/domain"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Orm *gorm.DB
}

func InitDB(env *Env, logger *zap.SugaredLogger) *Database {

	db := Database{}

	dsn := fmt.Sprintf(`host=%s user=%s password=%s  port=%d dbname=%s sslmode=disable timezone=America/Sao_Paulo`, env.DbHost, env.DbUser, env.DbPassword, env.DbPort, env.DbSid)

	orm, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatalw("Erro ao conectar ao banco de dados.", "erro", err)
	}

	db.Orm = orm

	if !orm.Migrator().HasTable(&domain.RealEstateProperty{}) {
		err = orm.AutoMigrate(&domain.RealEstateProperty{})
		if err != nil {
			logger.Fatal("Error migrating", zap.Error(err))
		}
	}

	if !orm.Migrator().HasTable(&domain.ExchangeRate{}) {
		err = orm.AutoMigrate(&domain.ExchangeRate{})
		if err != nil {
			logger.Fatal("Error migrating", zap.Error(err))
		}
	}

	return &db
}
