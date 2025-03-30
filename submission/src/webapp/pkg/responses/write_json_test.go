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
			payload:        func() {},
			statusCode:     http.StatusOK,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"Msg":"Failed to marshal response","Error":"See server log for more information"}`,
			expectError:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			WriteJSON(tc.payload, tc.statusCode, w)
			assert.Equal(t, tc.expectedStatus, w.Code, "status code should match")
			assert.Equal(t, "application/json", w.Header().Get("Content-Type"), "Content-Type should be application/json")

			if tc.expectError {
				assert.JSONEq(t, tc.expectedBody, w.Body.String(), "response body should match")
			} else {
				assert.JSONEq(t, tc.expectedBody, w.Body.String(), "response body should match")
			}
		})
	}
}
