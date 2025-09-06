package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"technoCredits/internal/routes"
	"technoCredits/internal/security"
	"technoCredits/pkg/brokers"
	"technoCredits/pkg/db"
	"technoCredits/pkg/logger"
)

var mainServer *Server

func ServiceStart() (err error) {
	gin.SetMode(security.AppSettings.AppParams.GinMode)
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     security.AppSettings.Cors.AllowOrigins,
		AllowMethods:     security.AppSettings.Cors.AllowMethods,
		AllowHeaders:     security.AppSettings.Cors.AllowHeaders,
		ExposeHeaders:    security.AppSettings.Cors.ExposeHeaders,
		AllowCredentials: security.AppSettings.Cors.AllowCredentials,
	}))

	mainServer = new(Server)
	go func() {
		if err = mainServer.Run(security.AppSettings.AppParams.PortRun, routes.InitRoutes(router)); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error while starting HTTP Service: %s", err)
		}
	}()

	return nil
}

func ServiceShutdown() {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	fmt.Printf("\n%s\n", yellow("Start of service termination"))

	err := db.CloseDBConn()
	if err != nil {
		strErr := fmt.Sprintf("Error closing database connection: %s", err.Error())
		fmt.Println(red(strErr))
		logger.Error.Println(strErr)
	}

	err = brokers.CloseRabbitMQ()
	if err != nil {
		strErr := fmt.Sprintf("Error closing rabbitmq connection: %s", err.Error())
		fmt.Println(red(strErr))
		logger.Error.Println(strErr)
		return
	}

	if err = mainServer.Shutdown(context.Background()); err != nil {
		strErr := fmt.Sprintf("Error shutting down server: %s", err.Error())
		fmt.Println(red(strErr))
		logger.Error.Println(strErr)
	} else {
		strSuccess := "HTTP-service termination successfully"
		fmt.Println(green(strSuccess))
		logger.Info.Println(strSuccess)
	}
}
