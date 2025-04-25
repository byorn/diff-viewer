package github_int

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func CreateGitHubIssue(w http.ResponseWriter, r *http.Request) {

	// Check if the request method is POST
	if r.Method != http.MethodPost {
		// If not, return 405 Method Not Allowed
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req IssueRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// 🔐 Extract token from Authorization header
	authHeader := r.Header.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
		return
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")

	github_username := getGitHubUsername(token)

	fmt.Printf("github username is %v,", github_username)
	req.Owner = github_username
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", req.Owner, req.Repo)

	payload := map[string]string{
		"title": req.Title,
		"body":  req.Body,
	}
	payloadBytes, _ := json.Marshal(payload)
	fmt.Printf(" api URL %v,", apiURL)
	reqHttp, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadBytes))
	reqHttp.Header.Set("Authorization", "Bearer "+token)
	reqHttp.Header.Set("Accept", "application/vnd.github+json")
	reqHttp.Header.Set("Content-Type", "application/json")
	reqHttp.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	resp, err := http.DefaultClient.Do(reqHttp)

	fmt.Printf("respojnse is %v,", resp)
	if err != nil {
		http.Error(w, "GitHub API request failed", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	respBody := map[string]interface{}{
		"status": resp.Status,
	}
	json.NewEncoder(w).Encode(respBody)
}

func getGitHubUsername(token string) string {
	req, _ := http.NewRequest("GET", "https://api.github.com/user", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return ""
	}
	defer resp.Body.Close()
	var data map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&data)
	return data["login"].(string)
}
