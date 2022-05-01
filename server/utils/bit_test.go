package utils

import "testing"

func TestTestBit(t *testing.T) {
	testCases := []struct {
		name           string
		vector         uint32
		index          uint8
		expectedResult bool
	}{
		{
			name:           "Empty vector, position 0",
			vector:         0b0,
			index:          0,
			expectedResult: false,
		},
		{
			name:           "Empty vector, position 15",
			vector:         0b0,
			index:          15,
			expectedResult: false,
		},
		{
			name:           "Full vector, position 0",
			vector:         0b1_1111_1111_1111_1111_1111_1111,
			index:          0,
			expectedResult: true,
		},
		{
			name:           "Full vector, position 22",
			vector:         0b1_1111_1111_1111_1111_1111_1111,
			index:          22,
			expectedResult: true,
		},
		{
			name:           "Random vector, position 6",
			vector:         0b1111_1111_0001_0011_1100,
			index:          6,
			expectedResult: false,
		},
		{
			name:           "Random vector, position 17",
			vector:         0b1111_1111_0001_0011_1100,
			index:          17,
			expectedResult: true,
		},
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
		{
			name:           "Empty vector, position 0",
			vector:         0b0,
			index:          0,
			expectedResult: 0b1,
		},
		{
			name:           "Empty vector, position 15",
			vector:         0b0,
			index:          15,
			expectedResult: 0b1000_0000_0000_0000,
		},
		{
			name:           "Full vector, position 0",
			vector:         0b1_1111_1111_1111_1111_1111_1111,
			index:          0,
			expectedResult: 0b1_1111_1111_1111_1111_1111_1111,
		},
		{
			name:           "Full vector, position 22",
			vector:         0b1_1111_1111_1111_1111_1111_1111,
			index:          22,
			expectedResult: 0b1_1111_1111_1111_1111_1111_1111,
		},
		{
			name:           "Random vector, position 6",
			vector:         0b1111_1111_0001_0011_1100,
			index:          6,
			expectedResult: 0b1111_1111_0001_0111_1100,
		},
		{
			name:           "Random vector, position 17",
			vector:         0b1111_1111_0001_0011_1100,
			index:          17,
			expectedResult: 0b1111_1111_0001_0011_1100,
		},
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
		{
			name:           "Empty vector, position 0",
			vector:         0b0,
			index:          0,
			expectedResult: 0b0,
		},
		{
			name:           "Empty vector, position 15",
			vector:         0b0,
			index:          15,
			expectedResult: 0b0,
		},
		{
			name:           "Full vector, position 0",
			vector:         0b1_1111_1111_1111_1111_1111_1111,
			index:          0,
			expectedResult: 0b1_1111_1111_1111_1111_1111_1110,
		},
		{
			name:           "Full vector, position 22",
			vector:         0b1_1111_1111_1111_1111_1111_1111,
			index:          22,
			expectedResult: 0b1_1011_1111_1111_1111_1111_1111},
		{
			name:           "Random vector, position 6",
			vector:         0b1111_1111_0001_0011_1100,
			index:          6,
			expectedResult: 0b1111_1111_0001_0011_1100,
		},
		{
			name:           "Random vector, position 17",
			vector:         0b1111_1111_0001_0011_1100,
			index:          17,
			expectedResult: 0b1101_1111_0001_0011_1100,
		},
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

func TestFlipBit(t *testing.T) {
	testCases := []struct {
		name           string
		vector         uint32
		index          uint8
		expectedResult uint32
	}{
		{
			name:           "Empty vector, position 0",
			vector:         0b0,
			index:          0,
			expectedResult: 0b1,
		},
		{
			name:           "Empty vector, position 15",
			vector:         0b0,
			index:          15,
			expectedResult: 0b1000_0000_0000_0000,
		},
		{
			name:           "Full vector, position 0",
			vector:         0b1_1111_1111_1111_1111_1111_1111,
			index:          0,
			expectedResult: 0b1_1111_1111_1111_1111_1111_1110,
		},
		{
			name:           "Full vector, position 22",
			vector:         0b1_1111_1111_1111_1111_1111_1111,
			index:          22,
			expectedResult: 0b1_1011_1111_1111_1111_1111_1111},
		{
			name:           "Random vector, position 6",
			vector:         0b1111_1111_0001_0011_1100,
			index:          6,
			expectedResult: 0b1111_1111_0001_0111_1100,
		},
		{
			name:           "Random vector, position 17",
			vector:         0b1111_1111_0001_0011_1100,
			index:          17,
			expectedResult: 0b1101_1111_0001_0011_1100,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Act
			result := FlipBit(testCase.vector, testCase.index)

			// Assert
			if result != testCase.expectedResult {
				t.Errorf("Incorrect result: expected %v, got %v", testCase.expectedResult, result)
			}
		})
	}
}
