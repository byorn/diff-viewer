package main

import (
	"diffviewer/dir_details"
	"diffviewer/github_int"
	"net/http"
)

func main() {
	http.HandleFunc("/dir_tree", corsMiddleware(dir_details.DirTreeHandler))
	http.HandleFunc("/create_issue", corsMiddleware(github_int.CreateGitHubIssue))
	// Start the HTTP server
	port := ":8080"
	println("Server is running on http://localhost" + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}

// corsMiddleware adds CORS headers to the response
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Add CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Change '*' to a specific origin if needed
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next(w, r)
	}
}
