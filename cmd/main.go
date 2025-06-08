package main

import (
	"log"

	"github.com/bapakfadil/fastcampus/internal/configs"
	"github.com/bapakfadil/fastcampus/internal/handlers/memberships"
	"github.com/bapakfadil/fastcampus/internal/handlers/posts"
	membershipRepo "github.com/bapakfadil/fastcampus/internal/repositories/memberships"
	postRepo "github.com/bapakfadil/fastcampus/internal/repositories/posts"
	membershipSvc "github.com/bapakfadil/fastcampus/internal/services/memberships"
	postSvc "github.com/bapakfadil/fastcampus/internal/services/posts"
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

	// Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	
	// List of used Repositories
	membershipRepo := membershipRepo.NewRepository(db)
	postRepo := postRepo.NewRepository(db)
	
	// List of used Services
	membershipService := membershipSvc.NewService(cfg, membershipRepo)
	postService := postSvc.NewService(cfg, postRepo)
	 
	// List of used Handlers
	membershipHandler := memberships.NewHandler(r, membershipService)
	membershipHandler.RegisterRoute()
	
	postHandler := posts.NewHandler(r, postService)
	postHandler.RegisterRoute()

	r.Run(cfg.Service.Port) // listen and run the server
}