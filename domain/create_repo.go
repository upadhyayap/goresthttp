package domain

type CreateRepoRequest struct {
	Name        string
	Description string
	Homepage    string
	private     bool
	HasIssues   bool
	HasProjects bool
	HasWiki     bool
}

// Modal fields which are required to us
type CreateRespoResponse struct {
	Id       int64     `json: "id"`
	Name     string    `json: "name"`
	FullName string    `json: "full_name"`
	Owner    RepoOwner `json: "owner"`
}

type RepoOwner struct {
	Id      int64
	Login   string
	Url     string
	HtmlUrl string
}
