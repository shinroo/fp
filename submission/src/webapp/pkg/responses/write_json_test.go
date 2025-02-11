package responses

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SampleStruct struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func TestWriteJSON(t *testing.T) {
	testCases := []struct {
		name           string
		payload        any
		statusCode     int
		expectedStatus int
		expectedBody   string
		expectError    bool
	}{
		{
			name: "Success with list of structs",
			payload: []SampleStruct{
				{ID: 1, Name: "Alice"},
				{ID: 2, Name: "Bob"},
			},
			statusCode:     http.StatusOK,
			expectedStatus: http.StatusOK,
			expectedBody:   `[{"id":1,"name":"Alice"},{"id":2,"name":"Bob"}]`,
			expectError:    false,
		},
		{
			name:           "Success with empty payload",
			payload:        nil,
			statusCode:     http.StatusOK,
			expectedStatus: http.StatusOK,
			expectedBody:   "null",
			expectError:    false,
		},
		{
			name:           "Marshal error with invalid payload",
			payload:        func() {}, // Functions cannot be marshaled to JSON
			statusCode:     http.StatusOK,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"Msg":"Failed to marshal response","Error":"See server log for more information"}`,
			expectError:    true,
		},
	}

	// Iterate over the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a ResponseRecorder to record the response
			w := httptest.NewRecorder()

			// Call the function
			WriteJSON(tc.payload, tc.statusCode, w)

			// Check the status code
			assert.Equal(t, tc.expectedStatus, w.Code, "status code should match")

			// Check the Content-Type header
			assert.Equal(t, "application/json", w.Header().Get("Content-Type"), "Content-Type should be application/json")

			// Check the response body
			if tc.expectError {
				// For error cases, check the exact error message
				assert.JSONEq(t, tc.expectedBody, w.Body.String(), "response body should match")
			} else {
				// For success cases, check the JSON-encoded body
				assert.JSONEq(t, tc.expectedBody, w.Body.String(), "response body should match")
			}
		})
	}
}
