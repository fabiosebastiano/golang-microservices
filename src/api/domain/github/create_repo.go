package github

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_project"`
	HasWiki     bool   `json:"has_wiki"`
}

//CreateRepoResponse .
type CreateRepoResponse struct {
	Id          uint64          `json:"id"`
	Name        string          `json:"name"`
	FullName    string          `json:"full_name"`
	Owner       RepoOwner       `json:"owner"`
	Permissions RepoPermissions `json:"permissions"`
}

//RepoOwner .
type RepoOwner struct {
	Id        uint64 `json:"id"`
	Login     string `json:"login"`
	URL       string `json:"url"`
	HtmlURL   string `json:"html_url"`
	ReposURL  string `json:"repos_url"`
	UserType  string `json:"type"`
	SiteAdmin bool   `json:"site_admin"`
}

//RepoPermissions .
type RepoPermissions struct {
	IsAdmin bool `json:"admin"`
	HasPush bool `json:"push"`
	HasPull bool `json:"pull"`
}
