package solver

import (
	"math/bits"
	"server/utils"
)

type freeVariables struct {
	indexes      []uint8
	affectedRows []uint32
}

type FreeVariableFixer interface {
	fixFreeVariables(augmentedMatrix *[MatrixSize]uint32, finalRow uint8)
}

type freeVariableFixer struct {
	optimizer Optimizer
}

func NewFreeVariableFixer(optimizer Optimizer) FreeVariableFixer {
	return &freeVariableFixer{optimizer: optimizer}
}

func (f *freeVariableFixer) fixFreeVariables(augmentedMatrix *[MatrixSize]uint32, finalRow uint8) {
	// Find the free variables
	freeVariables := findFreeVariables(augmentedMatrix, finalRow)
	if len(freeVariables.indexes) == 0 {
		return
	}

	// Find the optimal values for the free variables using brute force
	optimalValues := f.optimizer.determineOptimalValues(&freeVariables)

	// Set the free variables and do back-substitution according to the optimal values
	for i := uint8(0); i < uint8(len(freeVariables.indexes)); i++ {
		value := utils.TestBit(optimalValues, i)
		vector := getFreeVariableVector(freeVariables.indexes[i], value)

		for t := uint8(0); t < finalRow; t++ {
			if utils.TestBit(augmentedMatrix[t], freeVariables.indexes[i]) {
				augmentedMatrix[t] ^= vector
			}
		}

		augmentedMatrix[finalRow+i] = vector
	}
}

func findFreeVariables(augmentedMatrix *[MatrixSize]uint32, finalRow uint8) freeVariables {
	// Check for the edge case when all rows are empty
	if finalRow == 0 {
		return getFreeVariablesOfEmptyMatrix()
	}

	// Using signed index variables to eliminate the overflow at 0
	signedI := int8(finalRow - 1)
	signedJ := int8(MatrixSize - 1)

	// Check the matrix for free variables
	indexes := make([]uint8, 0)
	for signedI >= 0 && signedJ >= 0 && signedI != signedJ {
		i := uint8(signedI)
		j := uint8(signedJ)
		pivotColumn := uint8(bits.TrailingZeros32(augmentedMatrix[i]))

		// Store the free variables
		for ; j > pivotColumn; j-- {
			indexes = append(indexes, j)
		}

		signedI--
		signedJ = int8(j) - 1
	}

	if len(indexes) == 0 {
		return freeVariables{indexes: indexes, affectedRows: make([]uint32, 0)}
	}

	// Find the rows in the matrix do not just have bits set in the pivot column
	affectedRows := make([]uint32, 0, MatrixSize)
	for i := uint8(0); i < finalRow; i++ {
		if bits.OnesCount32(utils.ClearBit(augmentedMatrix[i], constantRow)) > 1 {
			affectedRows = append(affectedRows, augmentedMatrix[i])
		}
	}

	return freeVariables{indexes, affectedRows}
}

func getFreeVariablesOfEmptyMatrix() freeVariables {
	// All rows of the matrix being empty means that all variables are free
	indexes := make([]uint8, 0, MatrixSize)
	for i := uint8(0); i < MatrixSize; i++ {
		indexes = append(indexes, i)
	}

	return freeVariables{indexes: indexes, affectedRows: make([]uint32, 0)}
}

func getFreeVariableVector(index uint8, value bool) (vector uint32) {
	vector = utils.SetBit(vector, index)
	if value {
		vector = utils.SetBit(vector, constantRow)
	}

	return vector
}
