package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"yat-blog-notes/internal/configs"
	"yat-blog-notes/internal/infra/http_server"
	"yat-blog-notes/internal/repositories/postgres"

	http_controllers "yat-blog-notes/internal/infra/http_server/handlers"
)

func main() {
	config, err := configs.ReadConfig()
	if err != nil {
		log.Fatalf("could not read config: %v", err)
	}
	fmt.Println(config)

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Postgres.User,
		config.Postgres.Password,
		config.Postgres.Host,
		config.Postgres.Port,
		config.Postgres.DB,
	)
	notesDb, err := postgres.NewDB(dsn)
	if err != nil {
		log.Fatalf("could not start db: %v", err)
	}
	defer notesDb.Close()

	controller := http_controllers.NewNotesController()
	router := http_server.NewRouter(controller)
	server := http_server.NewServer(":8080", router)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	errorCh := make(chan error, 2)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			errs := []error{err}
			if cErr := server.Close(); cErr != nil {
				errs = append(errs, cErr)
			}
			errorCh <- errors.Join(errs...)
		}
	}()

	select {
	case sig := <-signalCh:
		log.Fatalf("received signal: %v", sig)
	case err := <-errorCh:
		log.Fatalf("received error: %v", err)
	}
	// finish up working
}
