package app

import (
	"C8/config"
	"C8/repository"
	"C8/route"
	"C8/service"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApplication() {
	repo := repository.NewRepo(config.PSQL.DB)
	app := service.NewService(repo)
	route.RegisterApi(router, app)

	address := os.Getenv("APP_ADDRESS")
	port := os.Getenv("APP_PORT")
	fmt.Println(port)
	router.Run(fmt.Sprintf("%s:%s", address, port))
}
