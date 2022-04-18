package solver

import (
	"reflect"
	"server/utils"
	"testing"
)

// Mock implementation of the gaussian eliminator interface
type mockGaussianEliminator struct {
	t         *testing.T
	allowCall bool
	board     uint32
	solvable  bool
	finalRow  uint8
	result    *[MatrixSize]uint32

	wasCalled bool
}

func (m *mockGaussianEliminator) gaussianEliminate(augmentedMatrix *[MatrixSize]uint32) (bool, uint8) {
	// Save that the mock was called
	m.wasCalled = true

	// Check whether it's OK to call the mock
	if !m.allowCall {
		m.t.Fatalf("Mock gaussian eliminator should not be called")
		return false, 0
	}

	// Set up the expected coefficient matrix
	expectedCoefficients := [MatrixSize]uint32{
		0b00000_00000_00000_00001_00011,
		0b00000_00000_00000_00010_00111,
		0b00000_00000_00000_00100_01110,
		0b00000_00000_00000_01000_11100,
		0b00000_00000_00000_10000_11000,
		0b00000_00000_00001_00011_00001,
		0b00000_00000_00010_00111_00010,
		0b00000_00000_00100_01110_00100,
		0b00000_00000_01000_11100_01000,
		0b00000_00000_10000_11000_10000,
		0b00000_00001_00011_00001_00000,
		0b00000_00010_00111_00010_00000,
		0b00000_00100_01110_00100_00000,
		0b00000_01000_11100_01000_00000,
		0b00000_10000_11000_10000_00000,
		0b00001_00011_00001_00000_00000,
		0b00010_00111_00010_00000_00000,
		0b00100_01110_00100_00000_00000,
		0b01000_11100_01000_00000_00000,
		0b10000_11000_10000_00000_00000,
		0b00011_00001_00000_00000_00000,
		0b00111_00010_00000_00000_00000,
		0b01110_00100_00000_00000_00000,
		0b11100_01000_00000_00000_00000,
		0b11000_10000_00000_00000_00000,
	}

	// Verify the input is correct
	for i := uint8(0); i < MatrixSize; i++ {
		expectedRow := expectedCoefficients[i]
		if utils.TestBit(m.board, i) {
			expectedRow = utils.SetBit(expectedRow, MatrixSize)
		}

		if augmentedMatrix[i] != expectedRow {
			m.t.Fatalf("Calling mock gaussian eliminator with incorrect input: expected %v, got %v (at matrix line %v)", expectedRow, augmentedMatrix[i], i)
		}
	}

	// Return the configured results
	*augmentedMatrix = *m.result
	return m.solvable, m.finalRow
}

// Mock implementation of the free variable fixer interface
type mockFreeVariableFixer struct {
	t         *testing.T
	allowCall bool
	matrix    *[MatrixSize]uint32
	finalRow  uint8
	result    *[MatrixSize]uint32

	wasCalled bool
}

func (m *mockFreeVariableFixer) fixFreeVariables(augmentedMatrix *[MatrixSize]uint32, finalRow uint8) {
	// Save that the mock was called
	m.wasCalled = true

	// Check whether it's OK to call the mock
	if !m.allowCall {
		m.t.Fatalf("Mock free variable fixer should not be called")
		return
	}

	// Check whether we got the expected input
	if finalRow != m.finalRow {
		m.t.Fatalf("Calling mock free variable fixer with incorrect input (finalRow): expected %v, got %v", m.finalRow, finalRow)
		return
	}

	if !reflect.DeepEqual(m.matrix, augmentedMatrix) {
		m.t.Fatal("Calling mock free variable fixer with incorrect input (augmentedMatrix)")
		return
	}

	// Set the configured result
	*augmentedMatrix = *m.result
}

func TestNoSolution(t *testing.T) {
	// Arrange
	board := uint32(0b00101_00011_10001_01100_10011)

	gaussianResult := [MatrixSize]uint32{
		0b1_00000_00000_00000_00000_00000,
		0b1_00000_00000_00000_00000_00000,
		0b1_00000_00000_00000_00000_00000,
		0b1_00000_00000_00000_00000_00000,
		0b1_00000_00000_00000_00000_00000,
		0b1_00000_00000_00000_00000_00000,
		0b1_00000_00000_00000_00000_00000,
		0b1_00000_00000_00000_00000_00000,
		0b1_00000_00000_00000_00000_00000,
		0b1_00000_00000_00000_00000_00000,
		0b1_00000_00000_00000_00000_00000,
		0b1_00000_00000_00000_00000_00000,
		0b1_00000_00000_00000_00000_00000,
		0b1_00000_00000_00000_00000_00000,
		0b1_00000_00000_00000_00000_00000,
		0b1_00000_00000_00000_00000_00000,
		0b1_00000_00000_00000_00000_00000,
		0b1_00000_00000_00000_00000_00000,
		0b1_00000_00000_00000_00000_00000,
		0b1_00000_00000_00000_00000_00000,
		0b1_00000_00000_00000_00000_00000,
		0b1_00000_00000_00000_00000_00000,
		0b1_00000_00000_00000_00000_00000,
		0b1_00000_00000_00000_00000_00000,
		0b1_00000_00000_00000_00000_00000,
	}
	gaussianEliminator := &mockGaussianEliminator{
		t:         t,
		allowCall: true,
		board:     board,
		solvable:  false,
		finalRow:  0,
		result:    &gaussianResult,
	}

	freeVariableFixer := &mockFreeVariableFixer{
		t:         t,
		allowCall: false,
		matrix:    nil,
		finalRow:  0,
		result:    nil,
	}

	solver := NewBoardSolver(gaussianEliminator, freeVariableFixer)

	// Act
	solvable, _ := solver.SolveBoard(board)

	// Assert
	if solvable {
		t.Error("Incorrect result: expected not solvable, got solvable")
	}

	if !gaussianEliminator.wasCalled {
		t.Error("Gaussian eliminator mock was not called")
	}

	if freeVariableFixer.wasCalled {
		t.Error("Free variable fixer was called when it should not have been")
	}
}

func TestHasSolution(t *testing.T) {
	// Arrange
	board := uint32(0b00101_00011_10001_01100_10011)

	gaussianResult := [MatrixSize]uint32{
		0b0_00000_00000_00000_00000_00001,
		0b0_00000_00000_00000_00000_00010,
		0b0_00000_00000_00000_00000_00100,
		0b0_00000_00000_00000_00000_01000,
		0b0_00000_00000_00000_00000_10000,
		0b0_00000_00000_00000_00001_00000,
		0b0_00000_00000_00000_00010_00000,
		0b0_00000_00000_00000_00100_00000,
		0b0_00000_00000_00000_01000_00000,
		0b0_00000_00000_00000_10000_00000,
		0b0_00000_00000_00000_00000_00000,
		0b0_00000_00000_00000_00000_00000,
		0b0_00000_00000_00000_00000_00000,
		0b0_00000_00000_00000_00000_00000,
		0b0_00000_00000_00000_00000_00000,
		0b0_00000_00000_00000_00000_00000,
		0b0_00000_00000_00000_00000_00000,
		0b0_00000_00000_00000_00000_00000,
		0b0_00000_00000_00000_00000_00000,
		0b0_00000_00000_00000_00000_00000,
		0b0_00000_00000_00000_00000_00000,
		0b0_00000_00000_00000_00000_00000,
		0b0_00000_00000_00000_00000_00000,
		0b0_00000_00000_00000_00000_00000,
		0b0_00000_00000_00000_00000_00000,
	}
	gaussianEliminator := &mockGaussianEliminator{
		t:         t,
		allowCall: true,
		board:     board,
		solvable:  true,
		finalRow:  10,
		result:    &gaussianResult,
	}

	fixedFreeVariablesResult := [MatrixSize]uint32{
		0b0_00000_00000_00000_00000_00001,
		0b0_00000_00000_00000_00000_00010,
		0b1_00000_00000_00000_00000_00100,
		0b1_00000_00000_00000_00000_01000,
		0b0_00000_00000_00000_00000_10000,
		0b0_00000_00000_00000_00001_00000,
		0b1_00000_00000_00000_00010_00000,
		0b0_00000_00000_00000_00100_00000,
		0b1_00000_00000_00000_01000_00000,
		0b1_00000_00000_00000_10000_00000,
		0b1_10000_00000_00000_00000_00000,
		0b1_01000_00000_00000_00000_00000,
		0b0_00100_00000_00000_00000_00000,
		0b0_00010_00000_00000_00000_00000,
		0b0_00001_00000_00000_00000_00000,
		0b0_00000_10000_00000_00000_00000,
		0b0_00000_01000_00000_00000_00000,
		0b1_00000_00100_00000_00000_00000,
		0b1_00000_00010_00000_00000_00000,
		0b0_00000_00001_00000_00000_00000,
		0b0_00000_00000_10000_00000_00000,
		0b1_00000_00000_01000_00000_00000,
		0b0_00000_00000_00100_00000_00000,
		0b0_00000_00000_00010_00000_00000,
		0b0_00000_00000_00001_00000_00000,
	}
	freeVariableFixer := &mockFreeVariableFixer{
		t:         t,
		allowCall: true,
		matrix:    &gaussianResult,
		finalRow:  10,
		result:    &fixedFreeVariablesResult,
	}

	solver := NewBoardSolver(gaussianEliminator, freeVariableFixer)

	// Act
	solvable, solution := solver.SolveBoard(board)

	// Assert
	if !solvable {
		t.Error("Incorrect result: expected solvable, got not solvable")
	}

	expectedSolution := uint32(0b_11000_00110_01000_11010_01100)
	if solution != expectedSolution {
		t.Errorf("Incorrect result for solution: expected %v, got %v", expectedSolution, solution)
	}

	if !gaussianEliminator.wasCalled {
		t.Error("Gaussian eliminator mock was not called")
	}

	if !freeVariableFixer.wasCalled {
		t.Error("Free variable fixer was not called")
	}
}
