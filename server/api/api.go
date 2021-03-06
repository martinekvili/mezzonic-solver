package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/solver"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Api interface {
	SetupHttpHandler() http.Handler
}

type api struct {
	solver solver.BoardSolver
}

func New(solver solver.BoardSolver) Api {
	return &api{solver: solver}
}

func (api *api) SetupHttpHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/api/solutions/{board:[0-9a-v]{1,5}}", api.solutionHandler).Methods("GET")

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	allowedOrigin := os.Getenv("FRONTEND_URL")
	if allowedOrigin == "" {
		return loggedRouter
	}

	corsAllowedOrigins := handlers.AllowedOrigins([]string{allowedOrigin})
	corsAllowedMethods := handlers.AllowedMethods([]string{"GET"})
	return handlers.CORS(corsAllowedOrigins, corsAllowedMethods)(loggedRouter)
}

func (api *api) solutionHandler(w http.ResponseWriter, r *http.Request) {
	board, err := parseBoard(r)
	if err != nil {
		log.Println("Bad request due to invalid board number", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "invalid board number")

		return
	}

	solvable, solutionNumber := api.solver.SolveBoard(board)

	log.Printf("Successful request for board %v, solvable: %v, solution: %v", board, solvable, solutionNumber)
	writeSolution(w, solvable, solutionNumber)
}

func parseBoard(r *http.Request) (uint32, error) {
	vars := mux.Vars(r)
	boardString := vars["board"]

	board, err := strconv.ParseUint(boardString, 32, 32)
	if err != nil {
		return 0, err
	}

	if board >= 1<<solver.MatrixSize {
		return 0, errors.New("invalid board number")
	}

	return uint32(board), nil
}
