package solver

import (
	"math/bits"
	"server/utils"
)

type GaussianEliminator interface {
	gaussianEliminate(augmentedMatrix *[MatrixSize]uint32) (bool, uint8)
}

type gaussianEliminator struct{}

func NewGaussianEliminator() GaussianEliminator {
	return gaussianEliminator{}
}

func (gaussianEliminator) gaussianEliminate(augmentedMatrix *[MatrixSize]uint32) (bool, uint8) {
	// Bring to row echelon form
	finalRow := transformToRowEchelon(augmentedMatrix)

	// Check for forbidden rows
	if hasForbiddenRow(augmentedMatrix, finalRow) {
		return false, 0
	}

	// Bring to reduced row echelon form
	backSubstitution(augmentedMatrix, finalRow)
	return true, finalRow
}

func transformToRowEchelon(augmentedMatrix *[MatrixSize]uint32) uint8 {
	i := uint8(0)
	j := uint8(0)

	for i < MatrixSize && j < MatrixSize {
		if !utils.TestBit(augmentedMatrix[i], j) && !swapPivot(augmentedMatrix, i, j) {
			j++
			continue
		}

		for t := i + 1; t < MatrixSize; t++ {
			if utils.TestBit(augmentedMatrix[t], j) {
				augmentedMatrix[t] ^= augmentedMatrix[i]
			}
		}

		i++
		j++
	}

	return i
}

func swapPivot(augmentedMatrix *[MatrixSize]uint32, i, j uint8) bool {
	for t := i + 1; t < MatrixSize; t++ {
		if utils.TestBit(augmentedMatrix[t], j) {
			augmentedMatrix[i], augmentedMatrix[t] = augmentedMatrix[t], augmentedMatrix[i]
			return true
		}
	}

	return false
}

func hasForbiddenRow(augmentedMatrix *[MatrixSize]uint32, finalRow uint8) bool {
	for t := finalRow; t < MatrixSize; t++ {
		if augmentedMatrix[t] > 0 {
			return true
		}
	}

	return false
}

func backSubstitution(augmentedMatrix *[MatrixSize]uint32, finalRow uint8) {
	// Using signed index variable to eliminate the overflow at 0
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
}
