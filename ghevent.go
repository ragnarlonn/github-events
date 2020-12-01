package ghevent

import "time"

type PushEvent struct {
	Ref        *string     `json:"ref,omitempty"`
	Before     *string     `json:"before,omitempty"`
	After      *string     `json:"after,omitempty"`
	Created    *bool       `json:"created,omitempty"`
	Deleted    *bool       `json:"deleted,omitempty"`
	Forced     *bool       `json:"forced,omitempty"`
	BaseRef    *string     `json:"base_ref,omitempty"`
	Compare    *string     `json:"compare,omitempty"`
	Commits    []Commit    `json:"commits,omitempty"`
	HeadCommit *Commit     `json:"head_commit,omitempty"`
	Repository *Repository `json:"repository,omitempty"`
	Pusher     *EmailUser  `json:"pusher,omitempty"`
	Sender     *User       `json:"sender,omitempty"`
}

type ForkEvent struct {
	Forkee     *Repository `json:"forkee,omitempty"`
	Repository *Repository `json:"repository,omitempty"`
	Sender     *User       `json:"sender,omitempty"`
}

type IssueCommentEvent struct {
	Action     *string       `json:"action,omitempty"`
	Issue      *Issue        `json:"issue,omitempty"`
	Comment    *IssueComment `json:"comment,omitempty"`
	Repository *Repository   `json:"repository,omitempty"`
	Sender     *User         `json:"sender,omitempty"`
}

type IssuesEvent struct {
	Action *string `json:"action,omitempty"`
	Issue  *Issue  `json:"issue,omitempty"`
	// Changes ??? object `json:"changes,omitempty"`
	Repository *Repository `json:"repository,omitempty"`
	Sender     *User       `json:"sender,omitempty"`
}

type LabelEvent struct {
	Action     *string     `json:"action,omitempty"`
	Label      *Label      `json:"label,omitempty"`
	Repository *Repository `json:"repository,omitempty"`
	Sender     *User       `json:"sender,omitempty"`
}

// type MarketplacePurchaseEvent struct {

type PullRequestEvent struct {
	Action      *string      `json:"action,omitempty"`
	Number      *int         `json:"number,omitempty"`
	PullRequest *PullRequest `json:"pull_request,omitempty"`
	Repository  *Repository  `json:"repository,omitempty"`
	Sender      *User        `json:"sender,omitempty"`
}

//
// Objects sent in the various events
//

type IssueComment struct {
	URL               *string    `json:"url,omitempty"`
	HTMLURL           *string    `json:"html_url,omitempty"`
	IssueURL          *string    `json:"issue_url,omitempty"`
	ID                *int       `json:"id,omitempty"`
	NodeID            *string    `json:"node_id,omitempty"`
	User              *User      `json:"user,omitempty"`
	CreatedAt         *time.Time `json:"created_at,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
	AuthorAssociation *string    `json:"author_association,omitempty"`
	Body              *string    `json:"body,omitempty"`
}

type Issue struct {
	ID               *int              `json:"id,omitempty"`
	NodeID           *string           `json:"node_id,omitempty"`
	URL              *string           `json:"url,omitempty"`
	RepositoryURL    *string           `json:"repository_url,omitempty"`
	LabelsURL        *string           `json:"labels_url,omitempty"`
	CommentsURL      *string           `json:"comments_url,omitempty"`
	EventsURL        *string           `json:"events_url,omitempty"`
	HTMLURL          *string           `json:"html_url,omitempty"`
	Number           *int              `json:"number,omitempty"`
	State            *string           `json:"state,omitempty"`
	Title            *string           `json:"title,omitempty"`
	Body             *string           `json:"body,omitempty"`
	User             *User             `json:"user,omitempty"`
	Labels           []Label           `json:"labels,omitempty"`
	Assignee         *User             `json:"assignee,omitempty"`
	Assignees        []User            `json:"assignees,omitempty"`
	Milestone        *Milestone        `json:"milestone,omitempty"`
	Locked           *bool             `json:"locked,omitempty"`
	ActiveLockReason *string           `json:"active_lock_reason,omitempty"`
	Comments         *int              `json:"comments,omitempty"`
	PullRequest      *IssuePullRequest `json:"pull_request,omitempty"`
	ClosedAt         *time.Time        `json:"closed_at,omitempty"`
	CreatedAt        *time.Time        `json:"created_at,omitempty"`
	UpdatedAt        *time.Time        `json:"updated_at,omitempty"`
	ClosedBy         *User             `json:"closed_by,omitempty"`
}

type IssuePullRequest struct {
	URL      *string `json:"url,omitempty"`
	HTMLURL  *string `json:"html_url,omitempty"`
	DiffURL  *string `json:"diff_url,omitempty"`
	PatchURL *string `json:"patch_url,omitempty"`
}

type Label struct {
	ID          *int    `json:"id,omitempty"`
	NodeID      *string `json:"node_id,omitempty"`
	URL         *string `json:"url,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Color       *string `json:"color,omitempty"`
	Default     *bool   `json:"default,omitempty"`
}

type Milestone struct {
	URL          *string    `json:"url,omitempty"`
	HTMLURL      *string    `json:"html_url,omitempty"`
	LabelsURL    *string    `json:"labels_url,omitempty"`
	ID           *int       `json:"id,omitempty"`
	NodeID       *string    `json:"node_id,omitempty"`
	Number       *int       `json:"number,omitempty"`
	State        *string    `json:"state,omitempty"`
	Title        *string    `json:"title,omitempty"`
	Description  *string    `json:"description,omitempty"`
	Creator      *User      `json:"creator,omitempty"`
	OpenIssues   *int       `json:"open_issues,omitempty"`
	ClosedIssues *int       `json:"closed_issues,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	ClosedAt     *time.Time `json:"closed_at,omitempty"`
	DueOn        *time.Time `json:"due_on,omitempty"`
}

type Repository struct {
	ID               *int       `json:"id,omitempty"`
	NodeID           *string    `json:"node_id,omitempty"`
	Name             *string    `json:"name,omitempty"`
	FullName         *string    `json:"full_name,omitempty"`
	Private          *bool      `json:"private,omitempty"`
	Owner            *User      `json:"owner,omitempty"`
	HTMLURL          *string    `json:"html_url,omitempty"`
	Description      *string    `json:"description,omitempty"`
	Fork             *bool      `json:"fork,omitempty"`
	URL              *string    `json:"url,omitempty"`
	ForksURL         *string    `json:"forks_url,omitempty"`
	KeysURL          *string    `json:"keys_url,omitempty"`
	CollaboratorsURL *string    `json:"collaborators_url,omitempty"`
	TeamsURL         *string    `json:"teams_url,omitempty"`
	HooksURL         *string    `json:"hooks_url,omitempty"`
	IssueEventsURL   *string    `json:"issue_events_url,omitempty"`
	EventsURL        *string    `json:"events_url,omitempty"`
	AssigneesURL     *string    `json:"assignees_url,omitempty"`
	BranchesURL      *string    `json:"branches_url,omitempty"`
	TagsURL          *string    `json:"tags_url,omitempty"`
	BlobsURL         *string    `json:"blobs_url,omitempty"`
	GitTagsURL       *string    `json:"git_tags_url,omitempty"`
	GitRefsURL       *string    `json:"git_refs_url,omitempty"`
	TreesURL         *string    `json:"trees_url,omitempty"`
	StatusesURL      *string    `json:"statuses_url,omitempty"`
	LanguagesURL     *string    `json:"languages_url,omitempty"`
	StargazersURL    *string    `json:"stargazers_url,omitempty"`
	ContributorsURL  *string    `json:"contributors_url,omitempty"`
	SubscribersURL   *string    `json:"subscribers_url,omitempty"`
	SubscriptionURL  *string    `json:"subscription_url,omitempty"`
	CommitsURL       *string    `json:"commits_url,omitempty"`
	GitCommitsURL    *string    `json:"git_commits_url,omitempty"`
	CommentsURL      *string    `json:"comments_url,omitempty"`
	IssueCommentsURL *string    `json:"issue_comment_url,omitempty"`
	ContentsURL      *string    `json:"contents_url,omitempty"`
	CompareURL       *string    `json:"compare_url,omitempty"`
	MergesURL        *string    `json:"merges_url,omitempty"`
	ArchiveURL       *string    `json:"archive_url,omitempty"`
	DownloadsURL     *string    `json:"downloads_url,omitempty"`
	IssuesURL        *string    `json:"issues_url,omitempty"`
	PullsURL         *string    `json:"pulls_url,omitempty"`
	MilestonesURL    *string    `json:"milestones_url,omitempty"`
	NotificationURL  *string    `json:"notifications_url,omitempty"`
	LabelsURL        *string    `json:"labels_url,omitempty"`
	ReleasesURL      *string    `json:"releases_url,omitempty"`
	DeploymentsURL   *string    `json:"deployments_url,omitempty"`
	CreatedAt        *uint32    `json:"created_at,omitempty"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty"`
	PushedAt         *uint32    `json:"pushed_at,omitempty"`
	GitURL           *string    `json:"git_url,omitempty"`
	SSHURL           *string    `json:"ssh_url,omitempty"`
	CloneURL         *string    `json:"clone_url,omitempty"`
	SvnURL           *string    `json:"svn_url,omitempty"`
	Homepage         *string    `json:"homepage,omitempty"`
	Size             *int       `json:"size,omitempty"`
	StargazersCount  *int       `json:"stargazers_count,omitempty"`
	WatchersCount    *int       `json:"watchers_count,omitempty"`
	Languages        *string    `json:"language,omitempty"`
	HasIssues        *bool      `json:"has_issues,omitempty"`
	HasProjects      *bool      `json:"has_projects,omitempty"`
	HasDownloads     *bool      `json:"has_downloads,omitempty"`
	HasWiki          *bool      `json:"has_wiki,omitempty"`
	HasPages         *bool      `json:"has_pages,omitempty"`
	ForksCount       *int       `json:"forks_count,omitempty"`
	MirrorURL        *string    `json:"mirror_url,omitempty"`
	Archived         *bool      `json:"archived,omitempty"`
	Disabled         *bool      `json:"disabled,omitempty"`
	OpenIssuesCount  *int       `json:"open_issues_count,omitempty"`
	License          *string    `json:"license,omitempty"`
	Forks            *int       `json:"forks,omitempty"`
	OpenIssues       *int       `json:"open_issues,omitempty"`
	Watchers         *int       `json:"watchers,omitempty"`
	DefaultBranch    *string    `json:"default_branch,omitempty"`
	Stargazers       *int       `json:"stargazers,omitempty"`
	MasterBranch     *string    `json:"master_branch,omitempty"`
}

type User struct {
	Login             *string    `json:"login,omitempty"`
	ID                *int       `json:"id,omitempty"`
	NodeID            *string    `json:"node_id,omitempty"`
	AvatarURL         *string    `json:"avatar_url,omitempty"`
	GravatarID        *string    `json:"gravatar_id,omitempty"`
	URL               *string    `json:"url,omitempty"`
	HTMLURL           *string    `json:"html_url,omitempty"`
	FollowersURL      *string    `json:"followers_url,omitempty"`
	FollowingURL      *string    `json:"following_url,omitempty"`
	GistsURL          *string    `json:"gists_url,omitempty"`
	StarredURL        *string    `json:"starred_url,omitempty"`
	SubscriptionsURL  *string    `json:"subscriptions_url,omitempty"`
	OrganizationsURL  *string    `json:"organizations_url,omitempty"`
	ReposURL          *string    `json:"repos_url,omitempty"`
	EventsURL         *string    `json:"events_url,omitempty"`
	ReceivedEventsURL *string    `json:"received_events_url,omitempty"`
	Type              *string    `json:"type,omitempty"`
	SiteAdmin         *bool      `json:"site_admin,omitempty"`
	Name              *string    `json:"name,omitempty"`
	Company           *string    `json:"company,omitempty"`
	Blog              *string    `json:"blog,omitempty"`
	Location          *string    `json:"location,omitempty"`
	Email             *string    `json:"email,omitempty"`
	Hireable          *bool      `json:"hireable,omitempty"`
	Bio               *string    `json:"bio,omitempty"`
	TwitterUsername   *string    `json:"twitter_username,omitempty"`
	PublicRepos       *int       `json:"public_repos,omitempty"`
	PublicGists       *int       `json:"public_gists,omitempty"`
	Followers         *int       `json:"followers,omitempty"`
	Following         *int       `json:"following,omitempty"`
	CreatedAt         *time.Time `json:"created_at,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
}

type CommitUser struct {
	Email *string    `json:"email,omitempty"`
	Name  *string    `json:"name,omitempty"`
	Date  *time.Time `json:"date,omitempty"`
}

type EmailUser struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
}

type TreeObject struct {
	URL *string `json:"url,omitempty"`
	SHA *string `json:"sha,omitempty"`
}

type VerificationObject struct {
	Verified  *bool   `json:"verified,omitempty"`
	Reason    *string `json:"reason,omitempty"`
	Signature *string `json:"signature,omitempty"`
	Payload   *string `json:"payload,omitempty"`
}

type CommitData struct {
	URL          *string             `json:"url,omitempty"`
	Author       *CommitUser         `json:"author,omitempty"`
	Committer    *CommitUser         `json:"committer,omitempty"`
	Message      *string             `json:"message,omitempty"`
	Tree         *TreeObject         `json:"tree,omitempty"`
	CommentCount *int                `json:"comment_count,omitempty"`
	Verification *VerificationObject `json:"verification,omitempty"`
}

type Commit struct {
	URL         *string      `json:"url,omitempty"`
	SHA         *string      `json:"sha,omitempty"`
	NodeID      *string      `json:"node_id,omitempty"`
	HTMLURL     *string      `json:"html_url,omitempty"`
	CommentsURL *string      `json:"comments_url,omitempty"`
	Commit      *CommitData  `json:"commit,omitempty"`
	Author      *User        `json:"author,omitempty"`
	Committer   *User        `json:"committer,omitempty"`
	Parents     []TreeObject `json:"parents,omitempty"`
}

type Installation struct {
	AccessTokensURL        *string           `json:"access_tokens_url,omitempty"`
	Account                *User             `json:"account,omitempty"`
	AppID                  *int              `json:"app_id,omitempty"`
	AppSlug                *string           `json:"app_slug,omitempty"`
	CreatedAt              *time.Time        `json:"created_at,omitempty"`
	Events                 []string          `json:"events,omitempty"`
	HasMultipleSingleFiles *bool             `json:"has_multiple_single_files,omitempty"`
	HTMLURL                *string           `json:"html_url,omitempty"`
	ID                     *int              `json:"id,omitempty"`
	Permissions            map[string]string `json:"permissions,omitempty"`
	RepositoriesURL        *string           `json:"repositories_url,omitempty"`
	RepositorySelection    *string           `json:"repository_selection,omitempty"`
	SingleFileName         *string           `json:"single_file_name,omitempty"`
	SingleFilePaths        []string          `json:"single_file_paths,omitempty"`
	SuspendedAt            *time.Time        `json:"suspended_at,omitempty"`
	SuspendedBy            *string           `json:"suspended_by,omitempty"`
	TargetID               *int              `json:"target_id,omitempty"`
	TargetType             *string           `json:"target_type,omitempty"`
	UpdatedAt              *time.Time        `json:"updated_at,omitempty"`
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
	IntegerField     *int      `json:"int,omitempty"`
	FloatField       *float64  `json:"float,omitempty"`
	StringField      *string   `json:"string,omitempty"`
	BoolField        *bool     `json:"bool,omitempty"`
	StringSliceField []string  `json:"stringslice,omitempty"`
	IntSliceField    []int     `json:"intslice,omitempty"`
	FloatSliceField  []float64 `json:"floatslice,omitempty"`
	BoolSliceField   []bool    `json:"boolslice,omitempty"`
}
