package main

import (
	"log"

	"github.com/bapakfadil/fastcampus/internal/configs"
	"github.com/bapakfadil/fastcampus/internal/handlers/memberships"
	membershipRepo "github.com/bapakfadil/fastcampus/internal/repositories/memberships"
	membershipSvc "github.com/bapakfadil/fastcampus/internal/services/memberships"
	"github.com/bapakfadil/fastcampus/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Gagal inisiasi config!")
	}

	cfg = configs.Get()
	log.Println("config", cfg)

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal inisiasi database!", err)
	}

	membershipRepo := membershipRepo.NewRepository(db)

	membershipService := membershipSvc.NewService(cfg, membershipRepo)
	
	membershipHandler := memberships.NewHandler(r, membershipService)
	membershipHandler.RegisterRoute()

	r.Run(cfg.Service.Port) // listen and run the server
}