package resp

import (
	"fmt"
	"net/http"
)

// WriteJSON - Send byte string json response
func WriteJSON(js []byte, code int, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(js)
}

func WriteErrorJSON(err error, code int, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(
		fmt.Sprintf("{'err': '%s', 'status_code': %d}", err.Error(), code)))
}
