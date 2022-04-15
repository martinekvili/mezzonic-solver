package solver

import "server/utils"

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

	optimalValues := uint32(0)
	optimalResult := calculateResultForValues(freeVariables, optimalValues)

	for values := uint32(1); values < 1<<len(freeVariables.indexes); values++ {
		result := calculateResultForValues(freeVariables, values)
		if optimalResult > result {
			optimalValues = values
			optimalResult = result
		}
	}

	return optimalValues
}

func calculateResultForValues(freeVariables *freeVariables, values uint32) (result uint8) {
	rows := make([]uint32, len(freeVariables.affectedRows))
	copy(rows, freeVariables.affectedRows)

	// Do back-substitution according to the current values
	for i := uint8(0); i < uint8(len(freeVariables.indexes)); i++ {
		value := utils.TestBit(values, i)
		vector := getFreeVariableVector(freeVariables.indexes[i], value)

		for t := uint8(0); t < uint8(len(rows)); t++ {
			if utils.TestBit(rows[t], freeVariables.indexes[i]) {
				rows[t] ^= vector
			}
		}
	}

	// Calculate the amount of "clicks" required in case of this solution
	// The "clicks" needed for the free variables
	for i := uint8(0); i < uint8(len(freeVariables.indexes)); i++ {
		if utils.TestBit(values, i) {
			result++
		}
	}

	// The "clicks" needed for the other affected variables
	for i := uint8(0); i < uint8(len(rows)); i++ {
		if utils.TestBit(rows[i], constantRow) {
			result++
		}
	}

	return result
}
