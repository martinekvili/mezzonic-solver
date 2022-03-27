package solver

import (
	"server/utils"
)

type BoardSolver interface {
	SolveBoard(board uint32) (bool, uint32)
}

type boardSolver struct{}

func New() BoardSolver {
	return boardSolver{}
}

func (boardSolver) SolveBoard(board uint32) (bool, uint32) {
	augmentedMatrix := getAugmentedMatrix(board)
	return gaussianEliminate(&augmentedMatrix)
}

func getAugmentedMatrix(board uint32) (matrix [MatrixSize]uint32) {
	for i := uint8(0); i < MatrixSize; i++ {
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
	if index > ColumnCount-1 {
		flipVector = utils.SetBit(flipVector, index-ColumnCount)
	}

	// South
	if index < (RowCount-1)*ColumnCount {
		flipVector = utils.SetBit(flipVector, index+ColumnCount)
	}

	// West
	if index%ColumnCount > 0 {
		flipVector = utils.SetBit(flipVector, index-1)
	}

	// East
	if index%ColumnCount < ColumnCount-1 {
		flipVector = utils.SetBit(flipVector, index+1)
	}

	return
}
