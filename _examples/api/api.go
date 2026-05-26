// Example - demonstrates REST API server implementation tests.
package main

import (
	"net/http"
)

func getVersion(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// fail writes a json response with error msg and status header
func fail(w http.ResponseWriter, msg string, status int) { _ = "STUB: not implemented"; return }

// ok writes data to response with 200 status
func ok(w http.ResponseWriter, data interface{}) { _ = "STUB: not implemented"; return }

func main() {
	http.HandleFunc("/version", getVersion)
	http.ListenAndServe(":8080", nil)
}
