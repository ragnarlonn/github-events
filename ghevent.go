package ghevent

//
// This package tries to describe the Github REST API, to make JSON encoding/decoding
// of Github data easier.
//
// - The goal is to include all *possible* fields in every object definition, even
//   though Github doesn't always include all fields for an object it returns in
//   response to an API call. I.e. the same type of object (e.g. the Repository
//   object) may contain different fields depending on what API end point it was
//   returned by. We actually try to model the Github *data structure* rather than
//   the API responses. It is up to the user to verify that they got all the fields
//   they needed in each response.
//
// - JSON fields used by the Github API will be CamelCase:d variables or types here
//   (i.e. a field called "head_commit" will become a Golang struct field called "HeadCommit"
//
// - All struct fields are pointers, which means that a missing JSON field will turn into
//   a nil pointer. This allows the user to detect which fields were present in the JSON
//   and which weren't (i.e. nil fields were NOT present)
//
// - When the structs are used to encode JSON data, empty fields will not be encoded -
//   all struct fields have an "omitempty" in their struct tags
//

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//
// Data wrappers for inconsistencies in the Github API
//

//
// Sometimes timestamps are integers (Unix epoch timestamps) and sometimes they're RFC3339. Sigh.
//
type TimeWrapper struct {
	t time.Time
}

func (tw *TimeWrapper) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	if t, err := time.Parse("2006-01-02T15:04:05Z", s); err == nil {
		tw.t = t
		return nil
	}
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		tw.t = t
		return nil
	}
	if epochSecs, err := strconv.Atoi(s); err != nil {
		return errors.New(fmt.Sprintf("Failed to parse time \"%s\" as either seconds since Epoch, or RFC3339-style string", s))
	} else {
		tw.t = time.Unix(int64(epochSecs), 0)
		return nil
	}
}
func (tw *TimeWrapper) MarshalJSON() ([]byte, error) {
	return tw.t.MarshalJSON()
}
func (tw *TimeWrapper) Time() time.Time {
	return tw.t
}

//
// Events
//

type PushEvent struct {
	Ref          *string       `json:"ref,omitempty"`
	Before       *string       `json:"before,omitempty"`
	After        *string       `json:"after,omitempty"`
	Created      *bool         `json:"created,omitempty"`
	Deleted      *bool         `json:"deleted,omitempty"`
	Forced       *bool         `json:"forced,omitempty"`
	BaseRef      *string       `json:"base_ref,omitempty"`
	Compare      *string       `json:"compare,omitempty"`
	Commits      []Commit      `json:"commits,omitempty"`
	HeadCommit   *Commit       `json:"head_commit,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
	Organization *Account      `json:"organization,omitempty"`
	Repository   *Repository   `json:"repository,omitempty"`
	Pusher       *EmailUser    `json:"pusher,omitempty"`
	Sender       *Account      `json:"sender,omitempty"`
}

type ForkEvent struct {
	Forkee     *Repository `json:"forkee,omitempty"`
	Repository *Repository `json:"repository,omitempty"`
	Sender     *Account    `json:"sender,omitempty"`
}

type IssueCommentEvent struct {
	Action     *string       `json:"action,omitempty"`
	Issue      *Issue        `json:"issue,omitempty"`
	Comment    *IssueComment `json:"comment,omitempty"`
	Repository *Repository   `json:"repository,omitempty"`
	Sender     *Account      `json:"sender,omitempty"`
}

type IssuesEvent struct {
	Action *string `json:"action,omitempty"`
	Issue  *Issue  `json:"issue,omitempty"`
	// Changes ??? object `json:"changes,omitempty"`
	Repository *Repository `json:"repository,omitempty"`
	Sender     *Account    `json:"sender,omitempty"`
}

type LabelEvent struct {
	Action     *string     `json:"action,omitempty"`
	Label      *Label      `json:"label,omitempty"`
	Repository *Repository `json:"repository,omitempty"`
	Sender     *Account    `json:"sender,omitempty"`
}

type PullRequestEvent struct {
	Action      *string      `json:"action,omitempty"`
	Number      *int         `json:"number,omitempty"`
	PullRequest *PullRequest `json:"pull_request,omitempty"`
	Repository  *Repository  `json:"repository,omitempty"`
	Sender      *Account     `json:"sender,omitempty"`
}

// X-Github-Event: "installation"
type InstallationEvent struct {
	Action       *string       `json:"action,omitempty"` // "created" | "deleted"
	Installation *Installation `json:"installation,omitempty"`
	Repositories []Repository  `json:"repositories,omitempty"`
	Sender       *Account      `json:"sender,omitempty"`
}

// X-Github-Event: "installation_repositories"
type InstallationRepositoriesEvent struct {
	Action              *string       `json:"action,omitempty"` // "added" | "removed"
	Installation        *Installation `json:"installation,omitempty"`
	RepositoriesAdded   []Repository  `json:"repositories_added,omitempty"`
	RepositoriesRemoved []Repository  `json:"repositories_removed,omitempty"`
	Sender              *Account      `json:"sender,omitempty"`
}

//
// Objects
//

type IssueComment struct {
	URL               *string      `json:"url,omitempty"`
	HTMLURL           *string      `json:"html_url,omitempty"`
	IssueURL          *string      `json:"issue_url,omitempty"`
	ID                *uint32      `json:"id,omitempty"`
	NodeID            *string      `json:"node_id,omitempty"`
	User              *Account     `json:"user,omitempty"`
	CreatedAt         *TimeWrapper `json:"created_at,omitempty"`
	UpdatedAt         *TimeWrapper `json:"updated_at,omitempty"`
	AuthorAssociation *string      `json:"author_association,omitempty"`
	Body              *string      `json:"body,omitempty"`
}

type Issue struct {
	ID               *uint32           `json:"id,omitempty"`
	NodeID           *string           `json:"node_id,omitempty"`
	URL              *string           `json:"url,omitempty"`
	RepositoryURL    *string           `json:"repository_url,omitempty"`
	LabelsURL        *string           `json:"labels_url,omitempty"`
	CommentsURL      *string           `json:"comments_url,omitempty"`
	EventsURL        *string           `json:"events_url,omitempty"`
	HTMLURL          *string           `json:"html_url,omitempty"`
	Number           *uint32           `json:"number,omitempty"`
	State            *string           `json:"state,omitempty"`
	Title            *string           `json:"title,omitempty"`
	Body             *string           `json:"body,omitempty"`
	User             *Account          `json:"user,omitempty"`
	Labels           []Label           `json:"labels,omitempty"`
	Assignee         *Account          `json:"assignee,omitempty"`
	Assignees        []Account         `json:"assignees,omitempty"`
	Milestone        *Milestone        `json:"milestone,omitempty"`
	Locked           *bool             `json:"locked,omitempty"`
	ActiveLockReason *string           `json:"active_lock_reason,omitempty"`
	Comments         *uint32           `json:"comments,omitempty"`
	PullRequest      *IssuePullRequest `json:"pull_request,omitempty"`
	ClosedAt         *TimeWrapper      `json:"closed_at,omitempty"`
	CreatedAt        *TimeWrapper      `json:"created_at,omitempty"`
	UpdatedAt        *TimeWrapper      `json:"updated_at,omitempty"`
	ClosedBy         *Account          `json:"closed_by,omitempty"`
}

type IssuePullRequest struct {
	URL      *string `json:"url,omitempty"`
	HTMLURL  *string `json:"html_url,omitempty"`
	DiffURL  *string `json:"diff_url,omitempty"`
	PatchURL *string `json:"patch_url,omitempty"`
}

type Label struct {
	ID          *uint32 `json:"id,omitempty"`
	NodeID      *string `json:"node_id,omitempty"`
	URL         *string `json:"url,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Color       *string `json:"color,omitempty"`
	Default     *bool   `json:"default,omitempty"`
}

type Milestone struct {
	URL          *string      `json:"url,omitempty"`
	HTMLURL      *string      `json:"html_url,omitempty"`
	LabelsURL    *string      `json:"labels_url,omitempty"`
	ID           *uint32      `json:"id,omitempty"`
	NodeID       *string      `json:"node_id,omitempty"`
	Number       *uint32      `json:"number,omitempty"`
	State        *string      `json:"state,omitempty"`
	Title        *string      `json:"title,omitempty"`
	Description  *string      `json:"description,omitempty"`
	Creator      *Account     `json:"creator,omitempty"`
	OpenIssues   *uint32      `json:"open_issues,omitempty"`
	ClosedIssues *uint32      `json:"closed_issues,omitempty"`
	CreatedAt    *TimeWrapper `json:"created_at,omitempty"`
	UpdatedAt    *TimeWrapper `json:"updated_at,omitempty"`
	ClosedAt     *TimeWrapper `json:"closed_at,omitempty"`
	DueOn        *TimeWrapper `json:"due_on,omitempty"`
}

type License struct {
	Key    *string `json:"key,omitempty"`
	Name   *string `json:"name,omitempty"`
	SpdxID *string `json:"spdx_id,omitempty"`
	URL    *string `json:"url,omitempty"`
	NodeID *string `json:"node_id,omitempty"`
}

type Repository struct {
	ID               *uint32      `json:"id,omitempty"`
	NodeID           *string      `json:"node_id,omitempty"`
	Name             *string      `json:"name,omitempty"`
	FullName         *string      `json:"full_name,omitempty"`
	Private          *bool        `json:"private,omitempty"`
	Owner            *Account     `json:"owner,omitempty"`
	HTMLURL          *string      `json:"html_url,omitempty"`
	Description      *string      `json:"description,omitempty"`
	Fork             *bool        `json:"fork,omitempty"`
	URL              *string      `json:"url,omitempty"`
	ForksURL         *string      `json:"forks_url,omitempty"`
	KeysURL          *string      `json:"keys_url,omitempty"`
	CollaboratorsURL *string      `json:"collaborators_url,omitempty"`
	TeamsURL         *string      `json:"teams_url,omitempty"`
	HooksURL         *string      `json:"hooks_url,omitempty"`
	IssueEventsURL   *string      `json:"issue_events_url,omitempty"`
	EventsURL        *string      `json:"events_url,omitempty"`
	AssigneesURL     *string      `json:"assignees_url,omitempty"`
	BranchesURL      *string      `json:"branches_url,omitempty"`
	TagsURL          *string      `json:"tags_url,omitempty"`
	BlobsURL         *string      `json:"blobs_url,omitempty"`
	GitTagsURL       *string      `json:"git_tags_url,omitempty"`
	GitRefsURL       *string      `json:"git_refs_url,omitempty"`
	TreesURL         *string      `json:"trees_url,omitempty"`
	StatusesURL      *string      `json:"statuses_url,omitempty"`
	LanguagesURL     *string      `json:"languages_url,omitempty"`
	StargazersURL    *string      `json:"stargazers_url,omitempty"`
	ContributorsURL  *string      `json:"contributors_url,omitempty"`
	SubscribersURL   *string      `json:"subscribers_url,omitempty"`
	SubscriptionURL  *string      `json:"subscription_url,omitempty"`
	CommitsURL       *string      `json:"commits_url,omitempty"`
	GitCommitsURL    *string      `json:"git_commits_url,omitempty"`
	CommentsURL      *string      `json:"comments_url,omitempty"`
	IssueCommentsURL *string      `json:"issue_comment_url,omitempty"`
	ContentsURL      *string      `json:"contents_url,omitempty"`
	CompareURL       *string      `json:"compare_url,omitempty"`
	MergesURL        *string      `json:"merges_url,omitempty"`
	ArchiveURL       *string      `json:"archive_url,omitempty"`
	DownloadsURL     *string      `json:"downloads_url,omitempty"`
	IssuesURL        *string      `json:"issues_url,omitempty"`
	PullsURL         *string      `json:"pulls_url,omitempty"`
	MilestonesURL    *string      `json:"milestones_url,omitempty"`
	NotificationURL  *string      `json:"notifications_url,omitempty"`
	LabelsURL        *string      `json:"labels_url,omitempty"`
	ReleasesURL      *string      `json:"releases_url,omitempty"`
	DeploymentsURL   *string      `json:"deployments_url,omitempty"`
	CreatedAt        *TimeWrapper `json:"created_at,omitempty"`
	UpdatedAt        *TimeWrapper `json:"updated_at,omitempty"`
	PushedAt         *TimeWrapper `json:"pushed_at,omitempty"`
	GitURL           *string      `json:"git_url,omitempty"`
	SSHURL           *string      `json:"ssh_url,omitempty"`
	CloneURL         *string      `json:"clone_url,omitempty"`
	SvnURL           *string      `json:"svn_url,omitempty"`
	Homepage         *string      `json:"homepage,omitempty"`
	Size             *uint32      `json:"size,omitempty"`
	StargazersCount  *uint32      `json:"stargazers_count,omitempty"`
	WatchersCount    *uint32      `json:"watchers_count,omitempty"`
	Languages        *string      `json:"language,omitempty"`
	HasIssues        *bool        `json:"has_issues,omitempty"`
	HasProjects      *bool        `json:"has_projects,omitempty"`
	HasDownloads     *bool        `json:"has_downloads,omitempty"`
	HasWiki          *bool        `json:"has_wiki,omitempty"`
	HasPages         *bool        `json:"has_pages,omitempty"`
	ForksCount       *uint32      `json:"forks_count,omitempty"`
	MirrorURL        *string      `json:"mirror_url,omitempty"`
	Archived         *bool        `json:"archived,omitempty"`
	Disabled         *bool        `json:"disabled,omitempty"`
	OpenIssuesCount  *uint32      `json:"open_issues_count,omitempty"`
	License          *License     `json:"license,omitempty"`
	Forks            *uint32      `json:"forks,omitempty"`
	OpenIssues       *uint32      `json:"open_issues,omitempty"`
	Watchers         *uint32      `json:"watchers,omitempty"`
	DefaultBranch    *string      `json:"default_branch,omitempty"`
	Stargazers       *uint32      `json:"stargazers,omitempty"`
	MasterBranch     *string      `json:"master_branch,omitempty"`
}

type Account struct {
	Login                    *string                   `json:"login,omitempty"`
	ID                       *uint32                   `json:"id,omitempty"`
	NodeID                   *string                   `json:"node_id,omitempty"`
	AvatarURL                *string                   `json:"avatar_url,omitempty"`
	GravatarID               *string                   `json:"gravatar_id,omitempty"`
	URL                      *string                   `json:"url,omitempty"`
	HTMLURL                  *string                   `json:"html_url,omitempty"`
	FollowersURL             *string                   `json:"followers_url,omitempty"`
	FollowingURL             *string                   `json:"following_url,omitempty"`
	GistsURL                 *string                   `json:"gists_url,omitempty"`
	StarredURL               *string                   `json:"starred_url,omitempty"`
	SubscriptionsURL         *string                   `json:"subscriptions_url,omitempty"`
	OrganizationsURL         *string                   `json:"organizations_url,omitempty"`
	ReposURL                 *string                   `json:"repos_url,omitempty"`
	EventsURL                *string                   `json:"events_url,omitempty"`
	ReceivedEventsURL        *string                   `json:"received_events_url,omitempty"`
	Type                     *string                   `json:"type,omitempty"`
	SiteAdmin                *bool                     `json:"site_admin,omitempty"`
	Name                     *string                   `json:"name,omitempty"`
	Company                  *string                   `json:"company,omitempty"`
	Blog                     *string                   `json:"blog,omitempty"`
	Location                 *string                   `json:"location,omitempty"`
	Email                    *string                   `json:"email,omitempty"`
	Hireable                 *bool                     `json:"hireable,omitempty"`
	Bio                      *string                   `json:"bio,omitempty"`
	TwitterUsername          *string                   `json:"twitter_username,omitempty"`
	PublicRepos              *uint32                   `json:"public_repos,omitempty"`
	PublicGists              *uint32                   `json:"public_gists,omitempty"`
	Followers                *uint32                   `json:"followers,omitempty"`
	Following                *uint32                   `json:"following,omitempty"`
	CreatedAt                *TimeWrapper              `json:"created_at,omitempty"`
	UpdatedAt                *TimeWrapper              `json:"updated_at,omitempty"`
	OrganizationBillingEmail *string                   `json:"organization_billing_email,omitempty"`
	MarketplacePendingChange *MarketplacePendingChange `json:"marketplace_pending_change,omitempty"`
	MarketplacePurchase      *MarketplacePurchase      `json:"marketplace_purchase,omitempty"`
}

type CommitUser struct {
	Email *string      `json:"email,omitempty"`
	Name  *string      `json:"name,omitempty"`
	Date  *TimeWrapper `json:"date,omitempty"`
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
	CommentCount *uint32             `json:"comment_count,omitempty"`
	Verification *VerificationObject `json:"verification,omitempty"`
}

type Commit struct {
	URL         *string      `json:"url,omitempty"`
	SHA         *string      `json:"sha,omitempty"`
	NodeID      *string      `json:"node_id,omitempty"`
	HTMLURL     *string      `json:"html_url,omitempty"`
	CommentsURL *string      `json:"comments_url,omitempty"`
	Commit      *CommitData  `json:"commit,omitempty"`
	Author      *Account     `json:"author,omitempty"`
	Committer   *Account     `json:"committer,omitempty"`
	Parents     []TreeObject `json:"parents,omitempty"`
}

type Installation struct {
	AccessTokensURL        *string           `json:"access_tokens_url,omitempty"`
	Account                *Account          `json:"account,omitempty"`
	AppID                  *uint32           `json:"app_id,omitempty"`
	AppSlug                *string           `json:"app_slug,omitempty"`
	CreatedAt              *TimeWrapper      `json:"created_at,omitempty"`
	Events                 []string          `json:"events,omitempty"`
	HasMultipleSingleFiles *bool             `json:"has_multiple_single_files,omitempty"`
	HTMLURL                *string           `json:"html_url,omitempty"`
	ID                     *uint32           `json:"id,omitempty"`
	NodeID                 *string           `json:"node_id,omitempty"`
	Permissions            map[string]string `json:"permissions,omitempty"`
	RepositoriesURL        *string           `json:"repositories_url,omitempty"`
	RepositorySelection    *string           `json:"repository_selection,omitempty"`
	SingleFileName         *string           `json:"single_file_name,omitempty"`
	SingleFilePaths        []string          `json:"single_file_paths,omitempty"`
	SuspendedAt            *TimeWrapper      `json:"suspended_at,omitempty"`
	SuspendedBy            *string           `json:"suspended_by,omitempty"`
	TargetID               *uint32           `json:"target_id,omitempty"`
	TargetType             *string           `json:"target_type,omitempty"`
	UpdatedAt              *TimeWrapper      `json:"updated_at,omitempty"`
}

type Plan struct {
	URL                 *string  `json:"url,omitempty"`
	AccountsURL         *string  `json:"accounts_url,omitempty"`
	ID                  *uint32  `json:"id,omitempty"`
	Number              *uint32  `json:"number,omitempty"`
	Name                *string  `json:"name,omitempty"`
	Description         *string  `json:"description,omitempty"`
	MonthlyPriceInCents *uint32  `json:"monthly_price_in_cents,omitempty"`
	YearlyPriceInCents  *uint32  `json:"yearly_price_in_cents,omitempty"`
	PriceModel          *string  `json:"price_model,omitempty"`
	HasFreeTrial        *bool    `json:"has_free_trial,omitempty"`
	State               *string  `json:"state,omitempty"`
	UnitName            *string  `json:"unit_name,omitempty"`
	Bullets             []string `json:"bullets,omitempty"`
}

type MarketplacePendingChange struct {
	EffectiveDate *TimeWrapper `json:"effective_date,omitempty"`
	UnitCount     *uint32      `json:"unit_count,omitempty"`
	ID            *uint32      `json:"id,omitempty"`
	Plan          *Plan        `json:"plan,omitempty"`
}

type MarketplacePurchase struct {
	BillingCycle    *string      `json:"billing_cycle,omitempty"`
	NextBillingDate *TimeWrapper `json:"next_billing_date"`
	UnitCount       *uint32      `json:"unit_count,omitempty"`
	OnFreeTrial     *bool        `json:"on_free_trial,omitempty"`
	FreeTrialEndsOn *TimeWrapper `json:"free_trial_ends_on,omitempty"`
	UpdatedAt       *TimeWrapper `json:"updated_at,omitempty"`
	Plan            *Plan        `json:"plan,omitempty"`
}

type PullRequest struct {
	// XXX TODO
}

//
// API responses
//
// Endpoint: /user/installations
type UserInstallationsResponse struct {
	Installations []Installation `json:"installations"`
	TotalCount    uint32         `json:"total_count"`
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
	IntegerField     *uint32   `json:"int,omitempty"`
	FloatField       *float64  `json:"float,omitempty"`
	StringField      *string   `json:"string,omitempty"`
	BoolField        *bool     `json:"bool,omitempty"`
	StringSliceField []string  `json:"stringslice,omitempty"`
	IntSliceField    []uint32  `json:"intslice,omitempty"`
	FloatSliceField  []float64 `json:"floatslice,omitempty"`
	BoolSliceField   []bool    `json:"boolslice,omitempty"`
}

/*

GHEventHandler: Got GHEvent request (X-Github-Event was "push")
GHEventHandler: request body was: {
   "after": "5aa9bdcf37e491356668c84e8c12b722063070ea",
   "base_ref": null,
   "before": "677a1c6abf3dcbd040a81762b7b9d3a869d5caa2",
   "commits": [
      {
         "added": [],
         "author": {
            "email": "ragnar@lonn.org",
            "name": "Ragnar Lonn",
            "username": "ragnarlonn"
         },
         "committer": {
            "email": "ragnar@lonn.org",
            "name": "Ragnar Lonn",
            "username": "ragnarlonn"
         },
         "distinct": true,
         "id": "5aa9bdcf37e491356668c84e8c12b722063070ea",
         "message": "Test commit",
         "modified": [
            "httpserver/database/relay.go"
         ],
         "removed": [],
         "timestamp": "2021-01-14T07:35:08+01:00",
         "tree_id": "fa2c871a795d610ce5589d060899d9123b7acb0b",
         "url": "https://github.com/0ddParity/badgebot/commit/5aa9bdcf37e491356668c84e8c12b722063070ea"
      }
   ],
   "compare": "https://github.com/0ddParity/badgebot/compare/677a1c6abf3d...5aa9bdcf37e4",
   "created": false,
   "deleted": false,
   "forced": false,
   "head_commit": {
      "added": [],
      "author": {
         "email": "ragnar@lonn.org",
         "name": "Ragnar Lonn",
         "username": "ragnarlonn"
      },
      "committer": {
         "email": "ragnar@lonn.org",
         "name": "Ragnar Lonn",
         "username": "ragnarlonn"
      },
      "distinct": true,
      "id": "5aa9bdcf37e491356668c84e8c12b722063070ea",
      "message": "Test commit",
      "modified": [
         "httpserver/database/relay.go"
      ],
      "removed": [],
      "timestamp": "2021-01-14T07:35:08+01:00",
      "tree_id": "fa2c871a795d610ce5589d060899d9123b7acb0b",
      "url": "https://github.com/0ddParity/badgebot/commit/5aa9bdcf37e491356668c84e8c12b722063070ea"
   },
   "installation": {
      "id": 14075073,
      "node_id": "MDIzOkludGVncmF0aW9uSW5zdGFsbGF0aW9uMTQwNzUwNzM="
   },
   "organization": {
      "avatar_url": "https://avatars2.githubusercontent.com/u/69793146?v=4",
      "description": "Introducing odd code to the world since 2018",
      "events_url": "https://api.github.com/orgs/0ddParity/events",
      "hooks_url": "https://api.github.com/orgs/0ddParity/hooks",
      "id": 69793146,
      "issues_url": "https://api.github.com/orgs/0ddParity/issues",
      "login": "0ddParity",
      "members_url": "https://api.github.com/orgs/0ddParity/members{/member}",
      "node_id": "MDEyOk9yZ2FuaXphdGlvbjY5NzkzMTQ2",
      "public_members_url": "https://api.github.com/orgs/0ddParity/public_members{/member}",
      "repos_url": "https://api.github.com/orgs/0ddParity/repos",
      "url": "https://api.github.com/orgs/0ddParity"
   },
   "pusher": {
      "email": "ragnar@lonn.org",
      "name": "ragnarlonn"
   },
   "ref": "refs/heads/master",
   "repository": {
      "archive_url": "https://api.github.com/repos/0ddParity/badgebot/{archive_format}{/ref}",
      "archived": false,
      "assignees_url": "https://api.github.com/repos/0ddParity/badgebot/assignees{/user}",
      "blobs_url": "https://api.github.com/repos/0ddParity/badgebot/git/blobs{/sha}",
      "branches_url": "https://api.github.com/repos/0ddParity/badgebot/branches{/branch}",
      "clone_url": "https://github.com/0ddParity/badgebot.git",
      "collaborators_url": "https://api.github.com/repos/0ddParity/badgebot/collaborators{/collaborator}",
      "comments_url": "https://api.github.com/repos/0ddParity/badgebot/comments{/number}",
      "commits_url": "https://api.github.com/repos/0ddParity/badgebot/commits{/sha}",
      "compare_url": "https://api.github.com/repos/0ddParity/badgebot/compare/{base}...{head}",
      "contents_url": "https://api.github.com/repos/0ddParity/badgebot/contents/{+path}",
      "contributors_url": "https://api.github.com/repos/0ddParity/badgebot/contributors",
      "created_at": 1597749125,
      "default_branch": "master",
      "deployments_url": "https://api.github.com/repos/0ddParity/badgebot/deployments",
      "description": "Create badges for your repo",
      "disabled": false,
      "downloads_url": "https://api.github.com/repos/0ddParity/badgebot/downloads",
      "events_url": "https://api.github.com/repos/0ddParity/badgebot/events",
      "fork": false,
      "forks": 0,
      "forks_count": 0,
      "forks_url": "https://api.github.com/repos/0ddParity/badgebot/forks",
      "full_name": "0ddParity/badgebot",
      "git_commits_url": "https://api.github.com/repos/0ddParity/badgebot/git/commits{/sha}",
      "git_refs_url": "https://api.github.com/repos/0ddParity/badgebot/git/refs{/sha}",
      "git_tags_url": "https://api.github.com/repos/0ddParity/badgebot/git/tags{/sha}",
      "git_url": "git://github.com/0ddParity/badgebot.git",
      "has_downloads": true,
      "has_issues": true,
      "has_pages": false,
      "has_projects": true,
      "has_wiki": true,
      "homepage": null,
      "hooks_url": "https://api.github.com/repos/0ddParity/badgebot/hooks",
      "html_url": "https://github.com/0ddParity/badgebot",
      "id": 288433559,
      "issue_comment_url": "https://api.github.com/repos/0ddParity/badgebot/issues/comments{/number}",
      "issue_events_url": "https://api.github.com/repos/0ddParity/badgebot/issues/events{/number}",
      "issues_url": "https://api.github.com/repos/0ddParity/badgebot/issues{/number}",
      "keys_url": "https://api.github.com/repos/0ddParity/badgebot/keys{/key_id}",
      "labels_url": "https://api.github.com/repos/0ddParity/badgebot/labels{/name}",
      "language": "Go",
      "languages_url": "https://api.github.com/repos/0ddParity/badgebot/languages",
      "license": null,
      "master_branch": "master",
      "merges_url": "https://api.github.com/repos/0ddParity/badgebot/merges",
      "milestones_url": "https://api.github.com/repos/0ddParity/badgebot/milestones{/number}",
      "mirror_url": null,
      "name": "badgebot",
      "node_id": "MDEwOlJlcG9zaXRvcnkyODg0MzM1NTk=",
      "notifications_url": "https://api.github.com/repos/0ddParity/badgebot/notifications{?since,all,participating}",
      "open_issues": 3,
      "open_issues_count": 3,
      "organization": "0ddParity",
      "owner": {
         "avatar_url": "https://avatars2.githubusercontent.com/u/69793146?v=4",
         "email": null,
         "events_url": "https://api.github.com/users/0ddParity/events{/privacy}",
         "followers_url": "https://api.github.com/users/0ddParity/followers",
         "following_url": "https://api.github.com/users/0ddParity/following{/other_user}",
         "gists_url": "https://api.github.com/users/0ddParity/gists{/gist_id}",
         "gravatar_id": "",
         "html_url": "https://github.com/0ddParity",
         "id": 69793146,
         "login": "0ddParity",
         "name": "0ddParity",
         "node_id": "MDEyOk9yZ2FuaXphdGlvbjY5NzkzMTQ2",
         "organizations_url": "https://api.github.com/users/0ddParity/orgs",
         "received_events_url": "https://api.github.com/users/0ddParity/received_events",
         "repos_url": "https://api.github.com/users/0ddParity/repos",
         "site_admin": false,
         "starred_url": "https://api.github.com/users/0ddParity/starred{/owner}{/repo}",
         "subscriptions_url": "https://api.github.com/users/0ddParity/subscriptions",
         "type": "Organization",
         "url": "https://api.github.com/users/0ddParity"
      },
      "private": true,
      "pulls_url": "https://api.github.com/repos/0ddParity/badgebot/pulls{/number}",
      "pushed_at": 1610606114,
      "releases_url": "https://api.github.com/repos/0ddParity/badgebot/releases{/id}",
      "size": 2128,
      "ssh_url": "git@github.com:0ddParity/badgebot.git",
      "stargazers": 0,
      "stargazers_count": 0,
      "stargazers_url": "https://api.github.com/repos/0ddParity/badgebot/stargazers",
      "statuses_url": "https://api.github.com/repos/0ddParity/badgebot/statuses/{sha}",
      "subscribers_url": "https://api.github.com/repos/0ddParity/badgebot/subscribers",
      "subscription_url": "https://api.github.com/repos/0ddParity/badgebot/subscription",
      "svn_url": "https://github.com/0ddParity/badgebot",
      "tags_url": "https://api.github.com/repos/0ddParity/badgebot/tags",
      "teams_url": "https://api.github.com/repos/0ddParity/badgebot/teams",
      "trees_url": "https://api.github.com/repos/0ddParity/badgebot/git/trees{/sha}",
      "updated_at": "2021-01-13T15:23:23Z",
      "url": "https://github.com/0ddParity/badgebot",
      "watchers": 0,
      "watchers_count": 0
   },
   "sender": {
      "avatar_url": "https://avatars2.githubusercontent.com/u/6524809?v=4",
      "events_url": "https://api.github.com/users/ragnarlonn/events{/privacy}",
      "followers_url": "https://api.github.com/users/ragnarlonn/followers",
      "following_url": "https://api.github.com/users/ragnarlonn/following{/other_user}",
      "gists_url": "https://api.github.com/users/ragnarlonn/gists{/gist_id}",
      "gravatar_id": "",
      "html_url": "https://github.com/ragnarlonn",
      "id": 6524809,
      "login": "ragnarlonn",
      "node_id": "MDQ6VXNlcjY1MjQ4MDk=",
      "organizations_url": "https://api.github.com/users/ragnarlonn/orgs",
      "received_events_url": "https://api.github.com/users/ragnarlonn/received_events",
      "repos_url": "https://api.github.com/users/ragnarlonn/repos",
      "site_admin": false,
      "starred_url": "https://api.github.com/users/ragnarlonn/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/ragnarlonn/subscriptions",
      "type": "User",
      "url": "https://api.github.com/users/ragnarlonn"
   }
}

*/
