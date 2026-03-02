package registrations

import (
	"github.com/samforedev/asignads/core/tenant-middleware/internal/api"
	"github.com/samforedev/asignads/core/tenant-middleware/internal/business"
	"github.com/samforedev/asignads/core/tenant-middleware/internal/config"
	"github.com/samforedev/asignads/core/tenant-middleware/internal/services"
)

func SetUp() {
	cfg := config.LoadConfig()
	rdbClient := services.InitRedis(cfg)
	dbClient := services.InitPostgres(cfg)

	redisRepo := services.NewRedisService(rdbClient)
	postgresRepo := services.NewPostgresService(dbClient)

	manager := business.NewTenantManager(redisRepo, postgresRepo)
	resolver := business.NewTenantResolver(redisRepo, postgresRepo)
	server := api.NewServer(cfg, resolver, manager)

	if err := server.Run(); err != nil {
		panic("Fatal error to run server: " + err.Error())
	}

}
