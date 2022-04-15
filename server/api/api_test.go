package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockSolver struct {
	t         *testing.T
	solutions map[uint32]struct {
		solvable       bool
		solutionNumber uint32
	}
}

func (m *mockSolver) SolveBoard(board uint32) (bool, uint32) {
	value, exists := m.solutions[board]
	if exists {
		return value.solvable, value.solutionNumber
	} else {
		m.t.Fatalf("Calling mock solver with unregistered input '%v'", board)
		return false, 0
	}
}

func TestInvalidRequest(t *testing.T) {
	testCases := []struct {
		name               string
		httpMethod         string
		httpPath           string
		expectedStatusCode int
	}{
		{
			name:               "Incorrect path",
			httpMethod:         "GET",
			httpPath:           "/api/solution",
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:               "Out-of-bound board (negative)",
			httpMethod:         "GET",
			httpPath:           "/api/solutions/-1",
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:               "Out-of-bound board (too big)",
			httpMethod:         "GET",
			httpPath:           "/api/solutions/100000",
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:               "Out-of-bound board (not a base32 number)",
			httpMethod:         "GET",
			httpPath:           "/api/solutions/cyp",
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:               "Invalid method",
			httpMethod:         "POST",
			httpPath:           "/api/solutions/c1p",
			expectedStatusCode: http.StatusMethodNotAllowed,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Arrage
			solver := &mockSolver{
				t: t,
				solutions: map[uint32]struct {
					solvable       bool
					solutionNumber uint32
				}{},
			}
			api := New(solver)
			handler := api.SetupHttpHandler()

			request := httptest.NewRequest(testCase.httpMethod, testCase.httpPath, nil)
			response := httptest.NewRecorder()

			// Act
			handler.ServeHTTP(response, request)

			// Assert
			result := response.Result()
			if result.StatusCode != testCase.expectedStatusCode {
				t.Errorf("Incorrect status code: expected %v, got %v", testCase.expectedStatusCode, result.StatusCode)
			}
		})
	}
}

func TestSuccessfulRequest(t *testing.T) {
	testCases := []struct {
		name                 string
		boardString          string
		boardNumber          uint32
		solvable             bool
		solutionNumber       uint32
		expectedResponseBody string
	}{
		{
			name:                 "No solution",
			boardString:          "none",
			boardNumber:          778990,
			solvable:             false,
			solutionNumber:       0b0,
			expectedResponseBody: "{\"hasSolution\":false,\"solution\":null}\n",
		},
		{
			name:                 "Empty solution",
			boardString:          "emptv",
			boardNumber:          15427519,
			solvable:             true,
			solutionNumber:       0b0,
			expectedResponseBody: "{\"hasSolution\":true,\"solution\":[]}\n",
		},
		{
			name:                 "Solution with multpile clicks",
			boardString:          "c1p",
			boardNumber:          12345,
			solvable:             true,
			solutionNumber:       0b1_1111,
			expectedResponseBody: "{\"hasSolution\":true,\"solution\":[0,1,2,3,4]}\n",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Arrage
			solver := &mockSolver{
				t: t,
				solutions: map[uint32]struct {
					solvable       bool
					solutionNumber uint32
				}{
					testCase.boardNumber: {
						solvable:       testCase.solvable,
						solutionNumber: testCase.solutionNumber,
					},
				},
			}
			api := New(solver)
			handler := api.SetupHttpHandler()

			request := httptest.NewRequest("GET", "/api/solutions/"+testCase.boardString, nil)
			response := httptest.NewRecorder()

			// Act
			handler.ServeHTTP(response, request)

			// Assert
			result := response.Result()
			if result.StatusCode != http.StatusOK {
				t.Errorf("Incorrect status code: expected %v, got %v", http.StatusOK, result.StatusCode)
			}
			if result.Header.Get("Content-Type") != "application/json" {
				t.Errorf("Incorrect 'Content-Type' header: expected '%v', got '%v'", "application/json", result.Header.Get("Content-Type"))
			}

			bodyBytes, err := io.ReadAll(result.Body)
			if err != nil {
				t.Fatalf("Error while reading response body %v", err)
			}
			body := string(bodyBytes)
			if body != testCase.expectedResponseBody {
				t.Errorf("Incorrect response body: expected '%v', got '%v'", testCase.expectedResponseBody, body)
			}
		})
	}
}
