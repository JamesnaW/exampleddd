package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sevenNt/echo-pprof"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"golang.org/x/time/rate"

	midware "exampleddd/middleware"

	authRepository "exampleddd/pkg/auth/repository"
	authService "exampleddd/pkg/auth/service"
	authTransport "exampleddd/pkg/auth/transport"

	profileRepository "exampleddd/pkg/profile/repository"
	profileService "exampleddd/pkg/profile/service"
	profileTransport "exampleddd/pkg/profile/transport"
)

func main() {
	dotenv := flag.Bool("dotenv", false, "Load env variable from .env file")
	flag.Parse()
	if *dotenv {
		if err := godotenv.Load(); err != nil {
			log.Fatal("unable to read .env file")
		}
	}

	// ENV business Config
	// SECRET := os.Getenv("SECRET")

	// ENV connection
	connStr := os.Getenv("DBCONN")

	// ENV server
	PORT := os.Getenv("PORT")
	Environment := os.Getenv("ENVIRONMENT")

	// Init Middleware
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// pprof trace
	if Environment != "production" {
		echopprof.Wrap(e)
	}

	// API Group
	api := e.Group("/api")

	// health
	api.GET("/health", health)

	// limiter
	limiter := rate.NewLimiter(10, 1)
	e.Use(midware.RateLimit(limiter))

	// e.Pre(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey: []byte(SECRET),
	// }))

	// Init connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// Init Endpoint
	// Profile
	profileRepo := profileRepository.NewRepo(db)
	profileServ := profileService.NewService(profileRepo)
	profileHandler := profileTransport.NewHandler(profileServ)
	profileTransport.Route(api, profileHandler)

	// Auth
	authRepo := authRepository.NewRepo(db)
	authServ := authService.NewService(authRepo, profileRepo)
	authHandler := authTransport.NewHandler(authServ)
	authTransport.Route(api, authHandler)

	// Ensuer access token
	e.Use(middleware.KeyAuth(midware.AccessToken(authRepo)))

	// Init HTTP Server
	go func() {
		if err := e.Start(":" + PORT); err != nil {
			e.Logger.Info("Shutting down the server")
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
