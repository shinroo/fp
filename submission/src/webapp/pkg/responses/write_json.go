package responses

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(payload any, statusCode int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	encoded, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(
			[]byte(
				`{"Msg": "Failed to marshal response", "Error": "See server log for more information"}`,
			),
		)
		return
	}

	w.WriteHeader(statusCode)
	_, _ = w.Write(encoded)
}
