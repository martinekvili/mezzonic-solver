package api

import (
	"encoding/json"
	"net/http"
	"server/solver"
	"server/utils"
)

type solution struct {
	HasSolution bool  `json:"hasSolution"`
	Solution    []int `json:"solution"`
}

func writeSolution(w http.ResponseWriter, solvable bool, solutionNumber uint32) {
	solution := createSolution(solvable, solutionNumber)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(solution)
}

func createSolution(solvable bool, solutionNumber uint32) solution {
	if !solvable {
		return solution{false, nil}
	}

	indexes := make([]int, 0)
	for i := uint8(0); i < solver.MatrixSize; i++ {
		if utils.TestBit(solutionNumber, i) {
			indexes = append(indexes, int(i))
		}
	}

	return solution{true, indexes}
}
