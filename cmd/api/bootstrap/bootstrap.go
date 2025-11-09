package bootstrap

import (
	"github.com/m4t1t0/GoCoinTracker/internal/asset"
	assetpg "github.com/m4t1t0/GoCoinTracker/internal/asset/repository/postgres"
	"github.com/m4t1t0/GoCoinTracker/internal/platform/db"
	"github.com/m4t1t0/GoCoinTracker/internal/platform/server"
	"os"
	"strconv"
)

func Run() error {
	port, _ := strconv.Atoi(os.Getenv("HTTP_PORT"))

	gormDB, err := db.Connect()
	if err != nil {
		return err
	}

	repo := assetpg.New(gormDB)
	assetsSvc := asset.NewService(repo)

	srv := server.New(uint(port), assetsSvc)
	return srv.Run()
}
