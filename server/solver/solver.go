package solver

import (
	"server/utils"
)

func SolveBoard(board uint32) (bool, uint32) {
	augmentedMatrix := getAugmentedMatrix(board)
	return gaussianEliminate(&augmentedMatrix)
}

func getAugmentedMatrix(board uint32) (matrix [matrixSize]uint32) {
	for i := uint8(0); i < matrixSize; i++ {
		flipVector := getFlipVector(i)
		if utils.TestBit(board, i) {
			flipVector = utils.SetBit(flipVector, constantRow)
		}

		matrix[i] = flipVector
	}

	return
}

func getFlipVector(index uint8) (flipVector uint32) {
	// The current position
	flipVector = utils.SetBit(flipVector, index)

	// North
	if index > 4 {
		flipVector = utils.SetBit(flipVector, index-5)
	}

	// South
	if index < 20 {
		flipVector = utils.SetBit(flipVector, index+5)
	}

	// West
	if index%5 > 0 {
		flipVector = utils.SetBit(flipVector, index-1)
	}

	// East
	if index%5 < 4 {
		flipVector = utils.SetBit(flipVector, index+1)
	}

	return
}
