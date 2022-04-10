package solver

import (
	"math/bits"
	"server/utils"
)

type BoardSolver interface {
	SolveBoard(board uint32) (bool, uint32)
}

type boardSolver struct {
	gaussianEliminator GaussianEliminator
	freeVariableFixer  FreeVariableFixer
}

func NewBoardSolver(gaussianEliminator GaussianEliminator, freeVariableFixer FreeVariableFixer) BoardSolver {
	return &boardSolver{gaussianEliminator: gaussianEliminator, freeVariableFixer: freeVariableFixer}
}

func (s *boardSolver) SolveBoard(board uint32) (bool, uint32) {
	// Create the initial augmented matrix
	augmentedMatrix := getAugmentedMatrix(board)

	// Run the gaussian elimination algorithm
	solvable, finalRow := s.gaussianEliminator.gaussianEliminate(&augmentedMatrix)
	if !solvable {
		return false, 0
	}

	// Fix the free variables to minimize "clicks" needed in the solution
	s.freeVariableFixer.fixFreeVariables(&augmentedMatrix, finalRow)

	// Determine the solution from the final matrix
	solution := determineSolution(&augmentedMatrix)
	return true, solution
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

func determineSolution(augmentedMatrix *[MatrixSize]uint32) (solution uint32) {
	for i := uint8(0); i < MatrixSize; i++ {
		pivotColumn := uint8(bits.TrailingZeros32(augmentedMatrix[i]))

		// Update the solution according to the current row
		if utils.TestBit(augmentedMatrix[i], constantRow) {
			solution = utils.SetBit(solution, pivotColumn)
		}
	}

	return
}
