package app

import (
	"context"
	"fmt"

	"github.com/MORE.Tech-5.0/server/internal/controller/http"
	"github.com/MORE.Tech-5.0/server/internal/repo"
	"github.com/MORE.Tech-5.0/server/internal/service"
	"github.com/MORE.Tech-5.0/server/pkg/log"
	"github.com/MORE.Tech-5.0/server/pkg/postgres"
)

func Run(cfg *Config, logger log.Logger) error {
	postgres, err := postgres.New(context.Background(), cfg.Postgres.URI)
	if err != nil {
		return err
	}
	defer postgres.Close()
	logger.Info("database connection established")

	officeRepo := repo.NewOfficePostgres(postgres)
	ohRepo := repo.NewOpenHoursPostgres(postgres)
	osRepo := repo.NewOfficeServicePostgres(postgres)
	atmRepo := repo.NewATMPostgres(postgres)
	asRepo := repo.NewServicePostgres(postgres)

	officeCache := repo.NewOfficeCache()
	atmCache := repo.NewATMCache()

	office := service.NewOffice(service.OfficeRepos{officeRepo, ohRepo, osRepo, officeCache})
	atm := service.NewATM(service.ATMRepos{atmRepo, asRepo, atmCache})

	server := http.NewServer(logger, http.Services{office, atm}, http.ServerOptions{
		Addr:    fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port),
		Origins: cfg.HTTP.Origins,
	})

	logger.Info("server created with address " + server.Addr)
	return fmt.Errorf("server down: %w", server.ListenAndServe())
}
