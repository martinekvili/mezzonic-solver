package api

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"server/solver"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func SetupHttpHandler() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/solutions/{board:[0-9a-v]{1,5}}", solutionHandler).Methods("GET")
	r.Use(loggingMiddleware)
	return r
}

func solutionHandler(w http.ResponseWriter, r *http.Request) {
	board, err := parseBoard(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "invalid board number")
	}

	w.WriteHeader(http.StatusOK)
	solvable, solutionNumber := solver.SolveBoard(board)
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

func loggingMiddleware(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, next)
}
