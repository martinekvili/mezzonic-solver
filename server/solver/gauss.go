package solver

import (
	"math/bits"
	"server/utils"
)

func gaussianEliminate(augmentedMatrix *[matrixSize]uint32) (bool, uint32) {
	// Bring to row echelon form
	finalRow := transformToRowEchelon(augmentedMatrix)

	// Check for forbidden rows
	if hasForbiddenRow(augmentedMatrix, finalRow) {
		return false, 0
	}

	// Bring to reduced row echelon form and determine the solution
	solution := determineSolution(augmentedMatrix, finalRow)
	return true, solution
}

func transformToRowEchelon(augmentedMatrix *[matrixSize]uint32) uint8 {
	i := uint8(0)
	j := uint8(0)

	for i < matrixSize && j < matrixSize {
		if !utils.TestBit(augmentedMatrix[i], j) && !swapPivot(augmentedMatrix, i, j) {
			j++
			continue
		}

		for t := i + 1; t < matrixSize; t++ {
			if utils.TestBit(augmentedMatrix[t], i) {
				augmentedMatrix[t] ^= augmentedMatrix[i]
			}
		}

		i++
		j++
	}

	return i
}

func swapPivot(augmentedMatrix *[matrixSize]uint32, i, j uint8) bool {
	for t := i + 1; t < matrixSize; t++ {
		if utils.TestBit(augmentedMatrix[t], j) {
			augmentedMatrix[i], augmentedMatrix[t] = augmentedMatrix[t], augmentedMatrix[i]
			return true
		}
	}

	return false
}

func hasForbiddenRow(augmentedMatrix *[matrixSize]uint32, finalRow uint8) bool {
	for t := finalRow; t < matrixSize; t++ {
		if augmentedMatrix[t] > 0 {
			return true
		}
	}

	return false
}

func determineSolution(augmentedMatrix *[matrixSize]uint32, finalRow uint8) (solution uint32) {
	// Using signed index variables to eliminate the overflow at 0
	signedI := int8(finalRow - 1)

	// Run back-substitution
	for signedI >= 0 {
		i := uint8(signedI)
		pivotColumn := uint8(bits.TrailingZeros32(augmentedMatrix[i]))

		// Back-substitution
		for t := uint8(0); t < i; t++ {
			if utils.TestBit(augmentedMatrix[t], pivotColumn) {
				augmentedMatrix[t] ^= augmentedMatrix[i]
			}
		}

		signedI--
	}

	// Fix the free variables to minimize "clicks" needed in the solution
	fixFreeVariables(augmentedMatrix, finalRow)

	// Determine the solution
	for i := uint8(0); i < matrixSize; i++ {
		pivotColumn := uint8(bits.TrailingZeros32(augmentedMatrix[i]))

		// Update the solution according to the current row
		if utils.TestBit(augmentedMatrix[i], constantRow) {
			solution = utils.SetBit(solution, pivotColumn)
		}
	}

	return solution
}
