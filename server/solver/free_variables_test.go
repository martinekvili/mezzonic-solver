package solver

import (
	"reflect"
	"sort"
	"testing"
)

// Methods that implement sort.Interface so we can sort affectedRows in freeVariables
func (f *freeVariables) Len() int {
	return len(f.affectedRows)
}

func (f *freeVariables) Less(i, j int) bool {
	return f.affectedRows[i] < f.affectedRows[j]
}

func (f *freeVariables) Swap(i, j int) {
	f.affectedRows[i], f.affectedRows[j] = f.affectedRows[j], f.affectedRows[i]
}

// Mock implementation of the optimizer interface
type mockOptimizer struct {
	t             *testing.T
	allowCall     bool
	freeVariables *freeVariables
	optimalValues uint32

	wasCalled bool
}

func (m *mockOptimizer) determineOptimalValues(freeVariables *freeVariables) uint32 {
	// Save that the mock was called
	m.wasCalled = true

	// Check whether it's OK to call the mock
	if !m.allowCall {
		m.t.Fatal("Mock optimizer should not be called")
		return 0
	}

	// Sort the affected rows in freeVariable values, as the order does not matter
	sort.Sort(m.freeVariables)
	sort.Sort(freeVariables)

	// Check whether we got the expected input
	if reflect.DeepEqual(m.freeVariables, freeVariables) {
		return m.optimalValues
	} else {
		m.t.Fatal("Calling mock optimizer with unexpected input")
		return 0
	}
}

func TestFixFreeVariables(t *testing.T) {
	testCases := []struct {
		name                string
		matrix              [MatrixSize]uint32
		finalRow            uint8
		shouldCallOptimizer bool
		freeVariables       *freeVariables
		optimalValues       uint32
		expectedResult      [MatrixSize]uint32
	}{
		{
			name: "Empty matrix",
			matrix: [MatrixSize]uint32{
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
			},
			finalRow:            0,
			shouldCallOptimizer: true,
			freeVariables: &freeVariables{
				indexes:      []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
				affectedRows: make([]uint32, 0),
			},
			optimalValues: 0b1_1000_0111_1101_1110_1110_1101,
			expectedResult: [MatrixSize]uint32{
				0b10_0000_0000_0000_0000_0000_0001,
				0b00_0000_0000_0000_0000_0000_0010,
				0b10_0000_0000_0000_0000_0000_0100,
				0b10_0000_0000_0000_0000_0000_1000,
				0b00_0000_0000_0000_0000_0001_0000,
				0b10_0000_0000_0000_0000_0010_0000,
				0b10_0000_0000_0000_0000_0100_0000,
				0b10_0000_0000_0000_0000_1000_0000,
				0b00_0000_0000_0000_0001_0000_0000,
				0b10_0000_0000_0000_0010_0000_0000,
				0b10_0000_0000_0000_0100_0000_0000,
				0b10_0000_0000_0000_1000_0000_0000,
				0b10_0000_0000_0001_0000_0000_0000,
				0b00_0000_0000_0010_0000_0000_0000,
				0b10_0000_0000_0100_0000_0000_0000,
				0b10_0000_0000_1000_0000_0000_0000,
				0b10_0000_0001_0000_0000_0000_0000,
				0b10_0000_0010_0000_0000_0000_0000,
				0b10_0000_0100_0000_0000_0000_0000,
				0b00_0000_1000_0000_0000_0000_0000,
				0b00_0001_0000_0000_0000_0000_0000,
				0b00_0010_0000_0000_0000_0000_0000,
				0b00_0100_0000_0000_0000_0000_0000,
				0b10_1000_0000_0000_0000_0000_0000,
				0b11_0000_0000_0000_0000_0000_0000,
			},
		},
		{
			name: "No free variables",
			matrix: [MatrixSize]uint32{
				0b00_0000_0000_0000_0000_0000_0001,
				0b00_0000_0000_0000_0000_0000_0010,
				0b00_0000_0000_0000_0000_0000_0100,
				0b10_0000_0000_0000_0000_0000_1000,
				0b10_0000_0000_0000_0000_0001_0000,
				0b10_0000_0000_0000_0000_0010_0000,
				0b00_0000_0000_0000_0000_0100_0000,
				0b00_0000_0000_0000_0000_1000_0000,
				0b00_0000_0000_0000_0001_0000_0000,
				0b10_0000_0000_0000_0010_0000_0000,
				0b10_0000_0000_0000_0100_0000_0000,
				0b10_0000_0000_0000_1000_0000_0000,
				0b10_0000_0000_0001_0000_0000_0000,
				0b00_0000_0000_0010_0000_0000_0000,
				0b00_0000_0000_0100_0000_0000_0000,
				0b00_0000_0000_1000_0000_0000_0000,
				0b10_0000_0001_0000_0000_0000_0000,
				0b10_0000_0010_0000_0000_0000_0000,
				0b10_0000_0100_0000_0000_0000_0000,
				0b00_0000_1000_0000_0000_0000_0000,
				0b10_0001_0000_0000_0000_0000_0000,
				0b10_0010_0000_0000_0000_0000_0000,
				0b00_0100_0000_0000_0000_0000_0000,
				0b00_1000_0000_0000_0000_0000_0000,
				0b11_0000_0000_0000_0000_0000_0000,
			},
			finalRow:            MatrixSize,
			shouldCallOptimizer: false,
			freeVariables:       nil,
			optimalValues:       0b0,
			expectedResult: [MatrixSize]uint32{
				0b00_0000_0000_0000_0000_0000_0001,
				0b00_0000_0000_0000_0000_0000_0010,
				0b00_0000_0000_0000_0000_0000_0100,
				0b10_0000_0000_0000_0000_0000_1000,
				0b10_0000_0000_0000_0000_0001_0000,
				0b10_0000_0000_0000_0000_0010_0000,
				0b00_0000_0000_0000_0000_0100_0000,
				0b00_0000_0000_0000_0000_1000_0000,
				0b00_0000_0000_0000_0001_0000_0000,
				0b10_0000_0000_0000_0010_0000_0000,
				0b10_0000_0000_0000_0100_0000_0000,
				0b10_0000_0000_0000_1000_0000_0000,
				0b10_0000_0000_0001_0000_0000_0000,
				0b00_0000_0000_0010_0000_0000_0000,
				0b00_0000_0000_0100_0000_0000_0000,
				0b00_0000_0000_1000_0000_0000_0000,
				0b10_0000_0001_0000_0000_0000_0000,
				0b10_0000_0010_0000_0000_0000_0000,
				0b10_0000_0100_0000_0000_0000_0000,
				0b00_0000_1000_0000_0000_0000_0000,
				0b10_0001_0000_0000_0000_0000_0000,
				0b10_0010_0000_0000_0000_0000_0000,
				0b00_0100_0000_0000_0000_0000_0000,
				0b00_1000_0000_0000_0000_0000_0000,
				0b11_0000_0000_0000_0000_0000_0000,
			},
		},
		{
			name: "10 free variables",
			matrix: [MatrixSize]uint32{
				0b00_0000_0000_0000_1000_0001_0001,
				0b00_0000_0000_0000_0000_0000_0100,
				0b10_0000_0000_0000_0000_0011_1000,
				0b00_0000_0001_0000_0001_0100_0000,
				0b10_0000_0000_0000_0100_1000_0000,
				0b10_0000_0000_0001_0000_0000_0000,
				0b00_0000_0000_0010_0000_0000_0000,
				0b00_0000_0000_0100_0000_0000_0000,
				0b00_0000_1001_1000_0000_0000_0000,
				0b00_0000_0010_0000_0000_0000_0000,
				0b10_0000_0100_0000_0000_0000_0000,
				0b00_0001_0000_0000_0000_0000_0000,
				0b10_0010_0000_0000_0000_0000_0000,
				0b11_0100_0000_0000_0000_0000_0000,
				0b00_1000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000, // finalRow (first empty row)
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
				0b00_0000_0000_0000_0000_0000_0000,
			},
			finalRow:            15,
			shouldCallOptimizer: true,
			freeVariables: &freeVariables{
				indexes: []uint8{24, 19, 16, 11, 10, 9, 8, 5, 4, 1},
				affectedRows: []uint32{
					0b00_0000_0000_0000_1000_0001_0001,
					0b10_0000_0000_0000_0000_0011_1000,
					0b00_0000_0001_0000_0001_0100_0000,
					0b10_0000_0000_0000_0100_1000_0000,
					0b00_0000_1001_1000_0000_0000_0000,
					0b11_0100_0000_0000_0000_0000_0000,
				},
			},
			optimalValues: 0b01_1111_0010,
			expectedResult: [MatrixSize]uint32{
				0b10_0000_0000_0000_0000_0000_0001,
				0b00_0000_0000_0000_0000_0000_0100,
				0b10_0000_0000_0000_0000_0000_1000,
				0b10_0000_0000_0000_0000_0100_0000,
				0b00_0000_0000_0000_0000_1000_0000,
				0b10_0000_0000_0001_0000_0000_0000,
				0b00_0000_0000_0010_0000_0000_0000,
				0b00_0000_0000_0100_0000_0000_0000,
				0b10_0000_0000_1000_0000_0000_0000,
				0b00_0000_0010_0000_0000_0000_0000,
				0b10_0000_0100_0000_0000_0000_0000,
				0b00_0001_0000_0000_0000_0000_0000,
				0b10_0010_0000_0000_0000_0000_0000,
				0b10_0100_0000_0000_0000_0000_0000,
				0b00_1000_0000_0000_0000_0000_0000,
				0b01_0000_0000_0000_0000_0000_0000, // finalRow (first row with a fixed free variable)
				0b10_0000_1000_0000_0000_0000_0000,
				0b00_0000_0001_0000_0000_0000_0000,
				0b00_0000_0000_0000_1000_0000_0000,
				0b10_0000_0000_0000_0100_0000_0000,
				0b10_0000_0000_0000_0010_0000_0000,
				0b10_0000_0000_0000_0001_0000_0000,
				0b10_0000_0000_0000_0000_0010_0000,
				0b10_0000_0000_0000_0000_0001_0000,
				0b00_0000_0000_0000_0000_0000_0010,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Arrange
			optimizer := mockOptimizer{
				t:             t,
				allowCall:     testCase.shouldCallOptimizer,
				freeVariables: testCase.freeVariables,
				optimalValues: testCase.optimalValues,
			}
			freeVariableFixer := NewFreeVariableFixer(&optimizer)

			// Act
			freeVariableFixer.fixFreeVariables(&testCase.matrix, testCase.finalRow)

			// Assert
			if !reflect.DeepEqual(testCase.expectedResult, testCase.matrix) {
				t.Error("Incorrect result")
			}

			if testCase.shouldCallOptimizer != optimizer.wasCalled {
				t.Errorf("Incorrect optimizer usage: should be called: %v, actually called: %v", testCase.shouldCallOptimizer, optimizer.wasCalled)
			}
		})
	}
}
