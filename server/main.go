package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"server/api"
	"server/solver"
	"syscall"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	api := setupApiWithDependencies()

	handler := api.SetupHttpHandler()
	srv := &http.Server{
		Addr: "0.0.0.0:" + port,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handler, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Printf("Starting web server, listening on port %v\n", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println(err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	log.Println("Starting graceful shutdown process")
	srv.Shutdown(ctx)
	log.Println("Shutdown successful")
}

func setupApiWithDependencies() api.Api {
	gaussianEliminator := solver.NewGaussianEliminator()

	optimizer := solver.NewBruteForceOptimizer()
	freeVariableFixer := solver.NewFreeVariableFixer(optimizer)

	solver := solver.NewBoardSolver(gaussianEliminator, freeVariableFixer)

	return api.New(solver)
}
