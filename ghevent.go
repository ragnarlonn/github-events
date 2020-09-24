package ghevent

import "time"

type PushEvent struct {
	Ref        *string     `json:"ref"`
	Before     *string     `json:"before"`
	After      *string     `json:"after"`
	Created    *bool       `json:"created"`
	Deleted    *bool       `json:"deleted"`
	Forced     *bool       `json:"forced"`
	BaseRef    *string     `json:"base_ref"`
	Compare    *string     `json:"compare"`
	Commits    []Commit    `json:"commits"`
	HeadCommit *Commit     `json:"head_commit"`
	Repository *Repository `json:"repository"`
	Pusher     *EmailUser  `json:"pusher"`
	Sender     *User       `json:"sender"`
}

type ForkEvent struct {
	Forkee     *Repository `json:"forkee"`
	Repository *Repository `json:"repository"`
	Sender     *User       `json:"sender"`
}

type IssueCommentEvent struct {
	Action     *string       `json:"action"`
	Issue      *Issue        `json:"issue"`
	Comment    *IssueComment `json:"comment"`
	Repository *Repository   `json:"repository"`
	Sender     *User         `json:"sender"`
}

type IssuesEvent struct {
	Action *string `json:"action"`
	Issue  *Issue  `json:"issue"`
	// Changes ??? object `json:"changes"`
	Repository *Repository `json:"repository"`
	Sender     *User       `json:"sender"`
}

type LabelEvent struct {
	Action     *string     `json:"action"`
	Label      *Label      `json:"label"`
	Repository *Repository `json:"repository"`
	Sender     *User       `json:"sender"`
}

// type MarketplacePurchaseEvent struct {

type PullRequestEvent struct {
	Action      *string      `json:"action"`
	Number      *int         `json:"number"`
	PullRequest *PullRequest `json:"pull_request"`
	Repository  *Repository  `json:"repository"`
	Sender      *User        `json:"sender"`
}

//
// Objects sent in the various events
//

type IssueComment struct {
	URL               *string    `json:"url"`
	HTMLURL           *string    `json:"html_url"`
	IssueURL          *string    `json:"issue_url"`
	ID                *int       `json:"id"`
	NodeID            *string    `json:"node_id"`
	User              *User      `json:"user"`
	CreatedAt         *time.Time `json:"created_at"`
	UpdatedAt         *time.Time `json:"updated_at"`
	AuthorAssociation *string    `json:"author_association"`
	Body              *string    `json:"body"`
}

type Issue struct {
	ID               *int              `json:"id"`
	NodeID           *string           `json:"node_id"`
	URL              *string           `json:"url"`
	RepositoryURL    *string           `json:"repository_url"`
	LabelsURL        *string           `json:"labels_url"`
	CommentsURL      *string           `json:"comments_url"`
	EventsURL        *string           `json:"events_url"`
	HTMLURL          *string           `json:"html_url"`
	Number           *int              `json:"number"`
	State            *string           `json:"state"`
	Title            *string           `json:"title"`
	Body             *string           `json:"body"`
	User             *User             `json:"user"`
	Labels           []Label           `json:"labels"`
	Assignee         *User             `json:"assignee"`
	Assignees        []User            `json:"assignees"`
	Milestone        *Milestone        `json:"milestone"`
	Locked           *bool             `json:"locked"`
	ActiveLockReason *string           `json:"active_lock_reason"`
	Comments         *int              `json:"comments"`
	PullRequest      *IssuePullRequest `json:"pull_request"`
	ClosedAt         *time.Time        `json:"closed_at"`
	CreatedAt        *time.Time        `json:"created_at"`
	UpdatedAt        *time.Time        `json:"updated_at"`
	ClosedBy         *User             `json:"closed_by"`
}

type IssuePullRequest struct {
	URL      *string `json:"url"`
	HTMLURL  *string `json:"html_url"`
	DiffURL  *string `json:"diff_url"`
	PatchURL *string `json:"patch_url"`
}

type Label struct {
	ID          *int    `json:"id"`
	NodeID      *string `json:"node_id"`
	URL         *string `json:"url"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Color       *string `json:"color"`
	Default     *bool   `json:"default"`
}

type Milestone struct {
	URL          *string    `json:"url"`
	HTMLURL      *string    `json:"html_url"`
	LabelsURL    *string    `json:"labels_url"`
	ID           *int       `json:"id"`
	NodeID       *string    `json:"node_id"`
	Number       *int       `json:"number"`
	State        *string    `json:"state"`
	Title        *string    `json:"title"`
	Description  *string    `json:"description"`
	Creator      *User      `json:"creator"`
	OpenIssues   *int       `json:"open_issues"`
	ClosedIssues *int       `json:"closed_issues"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	ClosedAt     *time.Time `json:"closed_at"`
	DueOn        *time.Time `json:"due_on"`
}

type Repository struct {
	ID               *int       `json:"id"`
	NodeID           *string    `json:"node_id"`
	Name             *string    `json:"name"`
	FullName         *string    `json:"full_name"`
	Private          *bool      `json:"private"`
	Owner            *User      `json:"owner"`
	HTMLURL          *string    `json:"html_url"`
	Description      *string    `json:"description"`
	Fork             *bool      `json:"fork"`
	URL              *string    `json:"url"`
	ForksURL         *string    `json:"forks_url"`
	KeysURL          *string    `json:"keys_url"`
	CollaboratorsURL *string    `json:"collaborators_url"`
	TeamsURL         *string    `json:"teams_url"`
	HooksURL         *string    `json:"hooks_url"`
	IssueEventsURL   *string    `json:"issue_events_url"`
	EventsURL        *string    `json:"events_url"`
	AssigneesURL     *string    `json:"assignees_url"`
	BranchesURL      *string    `json:"branches_url"`
	TagsURL          *string    `json:"tags_url"`
	BlobsURL         *string    `json:"blobs_url"`
	GitTagsURL       *string    `json:"git_tags_url"`
	GitRefsURL       *string    `json:"git_refs_url"`
	TreesURL         *string    `json:"trees_url"`
	StatusesURL      *string    `json:"statuses_url"`
	LanguagesURL     *string    `json:"languages_url"`
	StargazersURL    *string    `json:"stargazers_url"`
	ContributorsURL  *string    `json:"contributors_url"`
	SubscribersURL   *string    `json:"subscribers_url"`
	SubscriptionURL  *string    `json:"subscription_url"`
	CommitsURL       *string    `json:"commits_url"`
	GitCommitsURL    *string    `json:"git_commits_url"`
	CommentsURL      *string    `json:"comments_url"`
	IssueCommentsURL *string    `json:"issue_comment_url"`
	ContentsURL      *string    `json:"contents_url"`
	CompareURL       *string    `json:"compare_url"`
	MergesURL        *string    `json:"merges_url"`
	ArchiveURL       *string    `json:"archive_url"`
	DownloadsURL     *string    `json:"downloads_url"`
	IssuesURL        *string    `json:"issues_url"`
	PullsURL         *string    `json:"pulls_url"`
	MilestonesURL    *string    `json:"milestones_url"`
	NotificationURL  *string    `json:"notifications_url"`
	LabelsURL        *string    `json:"labels_url"`
	ReleasesURL      *string    `json:"releases_url"`
	DeploymentsURL   *string    `json:"deployments_url"`
	CreatedAt        *uint32    `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at"`
	PushedAt         *uint32    `json:"pushed_at"`
	GitURL           *string    `json:"git_url"`
	SSHURL           *string    `json:"ssh_url"`
	CloneURL         *string    `json:"clone_url"`
	SvnURL           *string    `json:"svn_url"`
	Homepage         *string    `json:"homepage"`
	Size             *int       `json:"size"`
	StargazersCount  *int       `json:"stargazers_count"`
	WatchersCount    *int       `json:"watchers_count"`
	Languages        *string    `json:"language"`
	HasIssues        *bool      `json:"has_issues"`
	HasProjects      *bool      `json:"has_projects"`
	HasDownloads     *bool      `json:"has_downloads"`
	HasWiki          *bool      `json:"has_wiki"`
	HasPages         *bool      `json:"has_pages"`
	ForksCount       *int       `json:"forks_count"`
	MirrorURL        *string    `json:"mirror_url"`
	Archived         *bool      `json:"archived"`
	Disabled         *bool      `json:"disabled"`
	OpenIssuesCount  *int       `json:"open_issues_count"`
	License          *string    `json:"license"`
	Forks            *int       `json:"forks"`
	OpenIssues       *int       `json:"open_issues"`
	Watchers         *int       `json:"watchers"`
	DefaultBranch    *string    `json:"default_branch"`
	Stargazers       *int       `json:"stargazers"`
	MasterBranch     *string    `json:"master_branch"`
}

type User struct {
	Login             *string    `json:"login"`
	ID                *int       `json:"id"`
	NodeID            *string    `json:"node_id"`
	AvatarURL         *string    `json:"avatar_url"`
	GravatarID        *string    `json:"gravatar_id"`
	URL               *string    `json:"url"`
	HTMLURL           *string    `json:"html_url"`
	FollowersURL      *string    `json:"followers_url"`
	FollowingURL      *string    `json:"following_url"`
	GistsURL          *string    `json:"gists_url"`
	StarredURL        *string    `json:"starred_url"`
	SubscriptionsURL  *string    `json:"subscriptions_url"`
	OrganizationsURL  *string    `json:"organizations_url"`
	ReposURL          *string    `json:"repos_url"`
	EventsURL         *string    `json:"events_url"`
	ReceivedEventsURL *string    `json:"received_events_url"`
	Type              *string    `json:"type"`
	SiteAdmin         *bool      `json:"site_admin"`
	Name              *string    `json:"name"`
	Company           *string    `json:"company"`
	Blog              *string    `json:"blog"`
	Location          *string    `json:"location"`
	Email             *string    `json:"email"`
	Hireable          *bool      `json:"hireable"`
	Bio               *string    `json:"bio"`
	TwitterUsername   *string    `json:"twitter_username"`
	PublicRepos       *int       `json:"public_repos"`
	PublicGists       *int       `json:"public_gists"`
	Followers         *int       `json:"followers"`
	Following         *int       `json:"following"`
	CreatedAt         *time.Time `json:"created_at"`
	UpdatedAt         *time.Time `json:"updated_at"`
}

type CommitUser struct {
	Email *string    `json:"email"`
	Name  *string    `json:"name"`
	Date  *time.Time `json:"date"`
}

type EmailUser struct {
	Email *string `json:"email"`
	Name  *string `json:"name"`
}

type TreeObject struct {
	URL *string `json:"url"`
	SHA *string `json:"sha"`
}

type VerificationObject struct {
	Verified  *bool   `json:"verified"`
	Reason    *string `json:"reason"`
	Signature *string `json:"signature"`
	Payload   *string `json:"payload"`
}

type CommitData struct {
	URL          *string             `json:"url"`
	Author       *CommitUser         `json:"author"`
	Committer    *CommitUser         `json:"committer"`
	Message      *string             `json:"message"`
	Tree         *TreeObject         `json:"tree"`
	CommentCount *int                `json:"comment_count"`
	Verification *VerificationObject `json:"verification"`
}

type Commit struct {
	URL         *string      `json:"url"`
	SHA         *string      `json:"sha"`
	NodeID      *string      `json:"node_id"`
	HTMLURL     *string      `json:"html_url"`
	CommentsURL *string      `json:"comments_url"`
	Commit      *CommitData  `json:"commit"`
	Author      *User        `json:"author"`
	Committer   *User        `json:"committer"`
	Parents     []TreeObject `json:"parents"`
}

type PullRequest struct {
	// XXX TODO
}

/*

{
  "ref": "refs/tags/simple-tag",
  "before": "6113728f27ae82c7b1a177c8d03f9e96e0adf246",
  "after": "0000000000000000000000000000000000000000",
  "created": false,
  "deleted": true,
  "forced": false,
  "base_ref": null,
  "compare": "https://github.com/Codertocat/Hello-World/compare/6113728f27ae...000000000000",
  "commits": [],
  "head_commit": null,
  "repository": {
    "id": 186853002,
    "node_id": "MDEwOlJlcG9zaXRvcnkxODY4NTMwMDI=",
    "name": "Hello-World",
    "full_name": "Codertocat/Hello-World",
    "private": false,
    "owner": {
      "name": "Codertocat",
      "email": "21031067+Codertocat@users.noreply.github.com",
      "login": "Codertocat",
      "id": 21031067,
      "node_id": "MDQ6VXNlcjIxMDMxMDY3",
      "avatar_url": "https://avatars1.githubusercontent.com/u/21031067?v=4",
      "gravatar_id": "",
      "url": "https://api.github.com/users/Codertocat",
      "html_url": "https://github.com/Codertocat",
      "followers_url": "https://api.github.com/users/Codertocat/followers",
      "following_url": "https://api.github.com/users/Codertocat/following{/other_user}",
      "gists_url": "https://api.github.com/users/Codertocat/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/Codertocat/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/Codertocat/subscriptions",
      "organizations_url": "https://api.github.com/users/Codertocat/orgs",
      "repos_url": "https://api.github.com/users/Codertocat/repos",
      "events_url": "https://api.github.com/users/Codertocat/events{/privacy}",
      "received_events_url": "https://api.github.com/users/Codertocat/received_events",
      "type": "User",
      "site_admin": false
    },
    "html_url": "https://github.com/Codertocat/Hello-World",
    "description": null,
    "fork": false,
    "url": "https://github.com/Codertocat/Hello-World",
    "forks_url": "https://api.github.com/repos/Codertocat/Hello-World/forks",
    "keys_url": "https://api.github.com/repos/Codertocat/Hello-World/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/Codertocat/Hello-World/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/Codertocat/Hello-World/teams",
    "hooks_url": "https://api.github.com/repos/Codertocat/Hello-World/hooks",
    "issue_events_url": "https://api.github.com/repos/Codertocat/Hello-World/issues/events{/number}",
    "events_url": "https://api.github.com/repos/Codertocat/Hello-World/events",
    "assignees_url": "https://api.github.com/repos/Codertocat/Hello-World/assignees{/user}",
    "branches_url": "https://api.github.com/repos/Codertocat/Hello-World/branches{/branch}",
    "tags_url": "https://api.github.com/repos/Codertocat/Hello-World/tags",
    "blobs_url": "https://api.github.com/repos/Codertocat/Hello-World/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/Codertocat/Hello-World/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/Codertocat/Hello-World/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/Codertocat/Hello-World/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/Codertocat/Hello-World/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/Codertocat/Hello-World/languages",
    "stargazers_url": "https://api.github.com/repos/Codertocat/Hello-World/stargazers",
    "contributors_url": "https://api.github.com/repos/Codertocat/Hello-World/contributors",
    "subscribers_url": "https://api.github.com/repos/Codertocat/Hello-World/subscribers",
    "subscription_url": "https://api.github.com/repos/Codertocat/Hello-World/subscription",
    "commits_url": "https://api.github.com/repos/Codertocat/Hello-World/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/Codertocat/Hello-World/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/Codertocat/Hello-World/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/Codertocat/Hello-World/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/Codertocat/Hello-World/contents/{+path}",
    "compare_url": "https://api.github.com/repos/Codertocat/Hello-World/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/Codertocat/Hello-World/merges",
    "archive_url": "https://api.github.com/repos/Codertocat/Hello-World/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/Codertocat/Hello-World/downloads",
    "issues_url": "https://api.github.com/repos/Codertocat/Hello-World/issues{/number}",
    "pulls_url": "https://api.github.com/repos/Codertocat/Hello-World/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/Codertocat/Hello-World/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/Codertocat/Hello-World/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/Codertocat/Hello-World/labels{/name}",
    "releases_url": "https://api.github.com/repos/Codertocat/Hello-World/releases{/id}",
    "deployments_url": "https://api.github.com/repos/Codertocat/Hello-World/deployments",
    "created_at": 1557933565,
    "updated_at": "2019-05-15T15:20:41Z",
    "pushed_at": 1557933657,
    "git_url": "git://github.com/Codertocat/Hello-World.git",
    "ssh_url": "git@github.com:Codertocat/Hello-World.git",
    "clone_url": "https://github.com/Codertocat/Hello-World.git",
    "svn_url": "https://github.com/Codertocat/Hello-World",
    "homepage": null,
    "size": 0,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": "Ruby",
    "has_issues": true,
    "has_projects": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": true,
    "forks_count": 1,
    "mirror_url": null,
    "archived": false,
    "disabled": false,
    "open_issues_count": 2,
    "license": null,
    "forks": 1,
    "open_issues": 2,
    "watchers": 0,
    "default_branch": "master",
    "stargazers": 0,
    "master_branch": "master"
  },
  "pusher": {
    "name": "Codertocat",
    "email": "21031067+Codertocat@users.noreply.github.com"
  },
  "sender": {
    "login": "Codertocat",
    "id": 21031067,
    "node_id": "MDQ6VXNlcjIxMDMxMDY3",
    "avatar_url": "https://avatars1.githubusercontent.com/u/21031067?v=4",
    "gravatar_id": "",
    "url": "https://api.github.com/users/Codertocat",
    "html_url": "https://github.com/Codertocat",
    "followers_url": "https://api.github.com/users/Codertocat/followers",
    "following_url": "https://api.github.com/users/Codertocat/following{/other_user}",
    "gists_url": "https://api.github.com/users/Codertocat/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/Codertocat/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/Codertocat/subscriptions",
    "organizations_url": "https://api.github.com/users/Codertocat/orgs",
    "repos_url": "https://api.github.com/users/Codertocat/repos",
    "events_url": "https://api.github.com/users/Codertocat/events{/privacy}",
    "received_events_url": "https://api.github.com/users/Codertocat/received_events",
    "type": "User",
    "site_admin": false
  }
}

*/

//
// TestStruct is used to test that JSON encoding/decoding works as intended
//

type TestStruct struct {
	IntegerField     *int      `json:"int"`
	FloatField       *float64  `json:"float"`
	StringField      *string   `json:"string"`
	BoolField        *bool     `json:"bool"`
	StringSliceField []string  `json:"stringslice"`
	IntSliceField    []int     `json:"intslice"`
	FloatSliceField  []float64 `json:"floatslice"`
	BoolSliceField   []bool    `json:"boolslice"`
}
