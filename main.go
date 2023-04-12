package main

import (
	"deploybot-service-monitor/task"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServerPort int `envconfig:"SERVER_PORT"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		panic(err)
	}

	g := gin.Default()

	r := task.NewRunner()
	g.GET("/serviceLogs", r.ServiceLogHandler())
	g.GET("/healthCheck", r.HealthCheckHandler())

	g.Run(fmt.Sprintf(":%d", cfg.ServerPort))
}
