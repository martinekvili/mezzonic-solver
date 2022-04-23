package solver

import (
	"math/bits"
	"server/utils"
)

type Optimizer interface {
	determineOptimalValues(freeVariables *freeVariables) uint32
}

type zeroValueOptimizer struct{}

func NewZeroValueOptimizer() Optimizer {
	return zeroValueOptimizer{}
}

func (zeroValueOptimizer) determineOptimalValues(freeVariables *freeVariables) uint32 {
	return uint32(0)
}

type bruteForceOptimizer struct{}

func NewBruteForceOptimizer() Optimizer {
	return bruteForceOptimizer{}
}

func (bruteForceOptimizer) determineOptimalValues(freeVariables *freeVariables) uint32 {
	if len(freeVariables.indexes) == 0 || len(freeVariables.affectedRows) == 0 {
		return uint32(0)
	}

	affectedSolution := getAffectedSolution(freeVariables)

	optimalValues := uint32(0)
	optimalResult := calculateResultForValues(freeVariables, affectedSolution, optimalValues)

	for values := uint32(1); values < 1<<len(freeVariables.indexes); values++ {
		result := calculateResultForValues(freeVariables, affectedSolution, values)
		if optimalResult > result {
			optimalValues = values
			optimalResult = result
		}
	}

	return optimalValues
}

func getAffectedSolution(freeVariables *freeVariables) (result uint32) {
	for i, affectedRow := range freeVariables.affectedRows {
		if utils.TestBit(affectedRow, constantRow) {
			result = utils.SetBit(result, uint8(i))
		}
	}

	return
}

func calculateResultForValues(freeVariables *freeVariables, affectedSolution uint32, values uint32) (result uint8) {
	// Do back-substitution according to the current values
	for i := uint8(0); i < uint8(len(freeVariables.indexes)); i++ {
		if !utils.TestBit(values, i) {
			continue
		}

		for t := uint8(0); t < uint8(len(freeVariables.affectedRows)); t++ {
			if utils.TestBit(freeVariables.affectedRows[t], freeVariables.indexes[i]) {
				affectedSolution = utils.FlipBit(affectedSolution, t)
			}
		}
	}

	// Calculate the amount of "clicks" required in case of this solution
	// The "clicks" needed for the free variables
	result = uint8(bits.OnesCount32(values))

	// The "clicks" needed for the other affected variables
	result += uint8(bits.OnesCount32(affectedSolution))

	return result
}
