package infrastructure

import (
	"github.com/jmoiron/sqlx"
	go_ora "github.com/sijms/go-ora/v2"
	"go.uber.org/zap"
)

type Database struct {
	Db    *sqlx.DB
	errDb error
}

func InitDB(env *Env, logger *zap.SugaredLogger) *Database {

	var dbs = Database{}

	logger.Info("Iniciando banco de dados")
	defer logger.Info("Banco de dados iniciado")

	sqlx.BindDriver("oracle", sqlx.NAMED)

	urlOptions := map[string]string{
		"SID": env.DbSid,
	}

	conn := go_ora.BuildUrl(env.DbHost, env.DbPort, "", env.DbUser, env.DbPassword, urlOptions)

	dbs.Db, dbs.errDb = sqlx.Open("oracle", conn)
	if dbs.errDb != nil {
		logger.Fatal("Erro ao iniciar banco de dados: ", zap.Error(dbs.errDb))
	}

	dbs.errDb = dbs.Db.Ping()
	if dbs.errDb == nil {
		return &dbs
	}

	conn = go_ora.BuildUrl(env.DbHost, env.DbPort, env.DbSid, env.DbUser, env.DbPassword, map[string]string{})

	dbs.Db, dbs.errDb = sqlx.Open("oracle", conn)
	if dbs.errDb != nil {
		logger.Fatal("Erro ao iniciar banco de dados: ", zap.Error(dbs.errDb))
	}

	dbs.errDb = dbs.Db.Ping()
	if dbs.errDb == nil {
		return &dbs
	}

	logger.Fatalw("Erro ao iniciar banco de dados", "erro", dbs.errDb)
	return nil
}
