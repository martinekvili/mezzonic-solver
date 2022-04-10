package utils

import "testing"

func TestTestBit(t *testing.T) {
	testCases := []struct {
		name           string
		vector         uint32
		index          uint8
		expectedResult bool
	}{
		{"Empty vector, position 0", 0b0, 0, false},
		{"Empty vector, position 15", 0b0, 15, false},
		{"Full vector, position 0", 0b1_1111_1111_1111_1111_1111_1111, 0, true},
		{"Full vector, position 22", 0b1_1111_1111_1111_1111_1111_1111, 22, true},
		{"Random vector, position 6", 0b1111_1111_0001_0011_1100, 6, false},
		{"Random vector, position 17", 0b1111_1111_0001_0011_1100, 17, true},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Act
			result := TestBit(testCase.vector, testCase.index)

			// Assert
			if result != testCase.expectedResult {
				t.Errorf("Incorrect result: expected %v, got %v", testCase.expectedResult, result)
			}
		})
	}
}

func TestSetBit(t *testing.T) {
	testCases := []struct {
		name           string
		vector         uint32
		index          uint8
		expectedResult uint32
	}{
		{"Empty vector, position 0", 0b0, 0, 0b1},
		{"Empty vector, position 15", 0b0, 15, 0b1000_0000_0000_0000},
		{"Full vector, position 0", 0b1_1111_1111_1111_1111_1111_1111, 0, 0b1_1111_1111_1111_1111_1111_1111},
		{"Full vector, position 22", 0b1_1111_1111_1111_1111_1111_1111, 22, 0b1_1111_1111_1111_1111_1111_1111},
		{"Random vector, position 6", 0b1111_1111_0001_0011_1100, 6, 0b1111_1111_0001_0111_1100},
		{"Random vector, position 17", 0b1111_1111_0001_0011_1100, 17, 0b1111_1111_0001_0011_1100},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Act
			result := SetBit(testCase.vector, testCase.index)

			// Assert
			if result != testCase.expectedResult {
				t.Errorf("Incorrect result: expected %v, got %v", testCase.expectedResult, result)
			}
		})
	}
}

func TestClearBit(t *testing.T) {
	testCases := []struct {
		name           string
		vector         uint32
		index          uint8
		expectedResult uint32
	}{
		{"Empty vector, position 0", 0b0, 0, 0b0},
		{"Empty vector, position 15", 0b0, 15, 0b0},
		{"Full vector, position 0", 0b1_1111_1111_1111_1111_1111_1111, 0, 0b1_1111_1111_1111_1111_1111_1110},
		{"Full vector, position 22", 0b1_1111_1111_1111_1111_1111_1111, 22, 0b1_1011_1111_1111_1111_1111_1111},
		{"Random vector, position 6", 0b1111_1111_0001_0011_1100, 6, 0b1111_1111_0001_0011_1100},
		{"Random vector, position 17", 0b1111_1111_0001_0011_1100, 17, 0b1101_1111_0001_0011_1100},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Act
			result := ClearBit(testCase.vector, testCase.index)

			// Assert
			if result != testCase.expectedResult {
				t.Errorf("Incorrect result: expected %v, got %v", testCase.expectedResult, result)
			}
		})
	}
}
