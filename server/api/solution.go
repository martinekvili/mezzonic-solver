package api

import (
	"encoding/json"
	"net/http"
	"server/solver"
	"server/utils"
)

type index struct {
	Row    uint8 `json:"row"`
	Column uint8 `json:"column"`
}

type solution struct {
	HasSolution bool    `json:"hasSolution"`
	Solution    []index `json:"solution"`
}

func writeSolution(w http.ResponseWriter, solvable bool, solutionNumber uint32) {
	solution := createSolution(solvable, solutionNumber)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(solution)
}

func createSolution(solvable bool, solutionNumber uint32) solution {
	if !solvable {
		return solution{false, nil}
	}

	indexes := make([]index, 0)
	for i := uint8(0); i < solver.MatrixSize; i++ {
		if utils.TestBit(solutionNumber, i) {
			index := index{i / solver.ColumnCount, i % solver.ColumnCount}
			indexes = append(indexes, index)
		}
	}

	return solution{true, indexes}
}
