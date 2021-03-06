package solver

import "testing"

func TestZeroValueOptimizer(t *testing.T) {
	// Arrange
	freeVariables := freeVariables{indexes: make([]uint8, 0), affectedRows: make([]uint32, 0)}

	// Act
	optimizer := NewZeroValueOptimizer()
	result := optimizer.determineOptimalValues(&freeVariables)

	// Assert
	if result != 0b0 {
		t.Errorf("Incorrect result: expected 0, got %v", result)
	}
}

func TestBruteForceOptimizer(t *testing.T) {
	testCases := []struct {
		name           string
		indexes        []uint8
		affectedRows   []uint32
		expectedResult uint32
	}{
		{
			name:           "No free variables",
			indexes:        make([]uint8, 0),
			affectedRows:   make([]uint32, 0),
			expectedResult: 0b0,
		},
		// Free state -> Affected state (free + affected = total) | Best
		// 0          -> 0 0            (0 + 0 = 0)                 *
		// 1          -> 1 1            (1 + 2 = 3)
		{
			name:    "Single free variable, expect 0",
			indexes: []uint8{15},
			affectedRows: []uint32{
				0b00_0000_0000_1000_0000_0000_0001,
				0b00_0000_0000_1000_0000_0000_0010,
			},
			expectedResult: 0b0,
		},
		// Free state -> Affected state (free + affected = total) | Best
		// 0          -> 1 1            (0 + 2 = 2)
		// 1          -> 0 0            (1 + 0 = 1)                 *
		{
			name:    "Single free variable, expect 1",
			indexes: []uint8{9},
			affectedRows: []uint32{
				0b10_0000_0000_0000_0010_0000_0001,
				0b10_0000_0000_0000_0010_0000_0010,
			},
			expectedResult: 0b1,
		},
		// Free state -> Affected state (free + affected = total) | Best
		// 0 0        -> 1 0            (0 + 1 = 1)                 *
		// 1 0        -> 0 1            (1 + 1 = 2)
		// 0 1        -> 1 1            (0 + 2 = 3)
		// 1 1        -> 0 0            (2 + 0 = 2)
		{
			name:    "Two free variables, expect (0, 0)",
			indexes: []uint8{15, 19},
			affectedRows: []uint32{
				0b10_0000_0000_1000_0000_0000_0001,
				0b00_0000_1000_1000_0000_0000_0010,
			},
			expectedResult: 0b0_0,
		},
		// Free state -> Affected state (free + affected = total) | Best
		// 0 0        -> 1 1 1 1 1      (0 + 5 = 5)
		// 1 0        -> 1 1 0 0 0      (1 + 2 = 3)
		// 0 1        -> 0 0 0 0 0      (1 + 0 = 1)                 *
		// 1 1        -> 0 0 1 1 1      (2 + 3 = 5)
		{
			name:    "Two free variables, expect (0, 1)",
			indexes: []uint8{19, 15},
			affectedRows: []uint32{
				0b10_0000_0000_1000_0000_0000_0001,
				0b10_0000_0000_1000_0000_0000_0010,
				0b10_0000_1000_1000_0000_0000_0100,
				0b10_0000_1000_1000_0000_0000_1000,
				0b10_0000_1000_1000_0000_0001_0000,
			},
			expectedResult: 0b1_0,
		},
		// Free state -> Affected state (free + affected = total) | Best
		// 0 0 0      -> 1 1 0 1 1      (0 + 4 = 4)
		// 1 0 0      -> 0 0 0 1 1      (1 + 2 = 3)
		// 0 1 0      -> 0 0 1 0 1      (1 + 2 = 3)
		// 1 1 0      -> 1 1 1 0 1      (2 + 4 = 6)
		// 0 0 1      -> 1 1 0 0 0      (1 + 2 = 3)
		// 1 0 1      -> 0 0 0 0 0      (2 + 0 = 2)                 *
		// 0 1 1      -> 0 0 1 1 0      (2 + 2 = 4)
		// 1 1 1      -> 1 1 1 1 0      (3 + 4 = 7)
		{
			name:    "Three free variables, expect (1, 0, 1)",
			indexes: []uint8{20, 21, 22},
			affectedRows: []uint32{
				0b10_0011_0000_0000_0000_0000_0001,
				0b10_0011_0000_0000_0000_0000_0010,
				0b00_0010_0000_0000_0000_0000_0100,
				0b10_0110_0000_0000_0000_0000_1000,
				0b10_0100_0000_0000_0000_0001_0000,
			},
			expectedResult: 0b1_0_1,
		},
		{
			name:           "Multiple free variables, no affected rows",
			indexes:        []uint8{1, 3, 5, 7, 9, 11},
			affectedRows:   make([]uint32, 0),
			expectedResult: 0b0_0_0_0_0_0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Arrange
			freeVariables := freeVariables{indexes: testCase.indexes, affectedRows: testCase.affectedRows}

			// Act
			optimizer := NewBruteForceOptimizer()
			result := optimizer.determineOptimalValues(&freeVariables)

			// Assert
			if result != testCase.expectedResult {
				t.Errorf("Incorrect result: expected %v, got %v", testCase.expectedResult, result)
			}
		})
	}
}
