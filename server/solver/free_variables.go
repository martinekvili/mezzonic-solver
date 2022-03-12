package solver

import (
	"math/bits"
	"server/utils"
)

type freeVariables struct {
	indexes      []uint8
	affectedRows []uint32
}

func fixFreeVariables(augmentedMatrix *[matrixSize]uint32, finalRow uint8) {
	// Find the free variables
	freeVariables := findFreeVariables(augmentedMatrix, finalRow)
	if len(freeVariables.indexes) == 0 {
		return
	}

	// Find the optimal values for the free variables using brute force
	optimalValues := determineOptimalValues(&freeVariables)

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

func findFreeVariables(augmentedMatrix *[matrixSize]uint32, finalRow uint8) freeVariables {
	indexes := make([]uint8, 0)
	affectedRowsMap := make(map[uint8]uint32)

	// Using signed index variables to eliminate the overflow at 0
	signedI := int8(finalRow - 1)
	signedJ := int8(matrixSize - 1)

	// Check the matrix for free variables
	for signedI >= 0 && signedJ >= 0 {
		i := uint8(signedI)
		j := uint8(signedJ)
		pivotColumn := uint8(bits.TrailingZeros32(augmentedMatrix[i]))

		// Store the free variables
		for ; j > pivotColumn; j-- {
			indexes = append(indexes, j)

			for t := uint8(0); t <= i; t++ {
				if utils.TestBit(augmentedMatrix[t], j) {
					affectedRowsMap[t] = augmentedMatrix[t]
				}
			}
		}

		signedI--
		signedJ = int8(j) - 1
	}

	affectedRows := make([]uint32, 0, len(affectedRowsMap))
	for _, affectedRow := range affectedRowsMap {
		affectedRows = append(affectedRows, affectedRow)
	}

	return freeVariables{indexes, affectedRows}
}

func determineOptimalValues(freeVariables *freeVariables) uint32 {
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
	for i := uint8(0); i < uint8(len(rows)); i++ {
		if utils.TestBit(rows[i], constantRow) {
			result++
		}
	}

	return result
}

func getFreeVariableVector(index uint8, value bool) (vector uint32) {
	vector = utils.SetBit(vector, index)
	if value {
		vector = utils.SetBit(vector, constantRow)
	}

	return vector
}
