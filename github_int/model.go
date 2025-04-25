package github_int

type IssueRequest struct {
	Owner string `json:"owner"` // GitHub username
	Repo  string `json:"repo"`  // Repo name
	Title string `json:"title"`
	Body  string `json:"body"`
}
