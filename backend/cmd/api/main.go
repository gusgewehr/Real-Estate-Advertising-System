package main

import (
	"real-estate-api/internal/adapters/gateway"
	"real-estate-api/internal/adapters/handler"
	"real-estate-api/internal/adapters/repository"
	"real-estate-api/internal/adapters/route"
	"real-estate-api/internal/application/usecase"
	"real-estate-api/internal/infrastructure"
)

func main() {
	logger := infrastructure.NewLogger()

	env := infrastructure.NewEnv(".env", logger)

	db := infrastructure.InitDB(env, logger)

	engine := infrastructure.NewEngine(env)

	repoPagination := repository.NewPaginationRepo(db, logger)
	ucPagination := usecase.NewPaginationUseCase(repoPagination, logger)

	repoRealEstate := repository.NewRealEstateRepository(db, logger)
	ucRealEstate := usecase.NewRealEstateUseCase(repoRealEstate, ucPagination, logger)
	handleRealEstate := handler.NewRealEstateHandler(ucRealEstate, logger)
	route.RealEstateRoutes(engine, handleRealEstate)

	repoFileStorage := repository.NewLocalFileStorage()
	ucFileStorage := usecase.NewFileStorageUseCase(repoFileStorage, env.FilePath, env.Host, env.FileUrl, logger)
	handleFileStorage := handler.NewFileStorageHandler(ucFileStorage, logger, env.MaxUploadSize, env.AllowedMIMETypes)
	route.FileStorageRoutes(engine, handleFileStorage)

	repoExchangeRate := repository.NewExchangeRateRepository(db, logger)
	ucExchangeRate := usecase.NewExchangeRateUseCase(repoExchangeRate, logger)
	handleExchangeRate := handler.NewExchangeRateHandler(ucExchangeRate, logger)
	route.ExchangeRateRoutes(engine, handleExchangeRate)

	repoZipcode := gateway.NewZipCodeGateway()
	ucZipcode := usecase.NewZipCodeUseCase(repoZipcode, logger)
	handleZipcode := handler.NewZipCodeHandler(ucZipcode)
	route.ZipCodeRoutes(engine, handleZipcode)

	infrastructure.StartHttpServer(engine, logger, env)

}
