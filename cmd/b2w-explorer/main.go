package main

import (
	"b2w-explorer/internal/app/service"
	api_routes "b2w-explorer/internal/infra/api-routes"
	"b2w-explorer/internal/infra/clients"
	"b2w-explorer/internal/infra/util"
	"github.com/igorralexsander/httpcircuited"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	e := echo.New()

	httpClient := httpcircuited.NewHttpClient()

	config := httpClient.NewConfigBuilder().
		BaseUrl("https://www.americanas.com.br").
		WithName("default").
		AddHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36").
		Timeout(800).
		CircuitFailureRatio(0.3).
		Build()

	httpClient.AddDownstream(*config)

	b2wClient := clients.NewB2WClient(httpClient.GetDownstream("default"))

	productPageService := service.NewProductPage(b2wClient)

	productPageRoute := api_routes.NewProductPageRoute(productPageService)

	productPageRoute.RegisterEndpoints(e)

	e.Use(util.MetricsMiddleware())

	startMetricsServer()

	go func() {
		log.Fatal(e.Start(":8080"))
	}()

	gracefullyShutdown()
}

func startMetricsServer() {
	metricsServer := echo.New()
	metricsServer.HideBanner = true
	metricsServer.HidePort = true
	metricsServer.GET("/prometheus", echo.WrapHandler(promhttp.Handler()))
	go func() {
		log.Info("Starting metrics server in port: 8081")
		log.Fatal(metricsServer.Start(":8081"))
	}()

}

func gracefullyShutdown() {
	// Listen for system signals to gracefully stop the application
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	switch <-signalChannel {
	case os.Interrupt:
		log.Info("Received SIGINT, stopping...")
	case syscall.SIGTERM:
		log.Info("Received SIGTERM, stopping...")
	}
}
