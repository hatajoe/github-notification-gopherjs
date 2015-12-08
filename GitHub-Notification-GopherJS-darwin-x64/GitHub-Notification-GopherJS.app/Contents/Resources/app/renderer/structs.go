package main

import "time"

/*
https://developer.github.com/v3/activity/notifications/
[
  {
    "id": "1",
    "repository": {
      "id": 1296269,
      "owner": {
        "login": "octocat",
        "id": 1,
        "avatar_url": "https://github.com/images/error/octocat_happy.gif",
        "gravatar_id": "",
        "url": "https://api.github.com/users/octocat",
        "html_url": "https://github.com/octocat",
        "followers_url": "https://api.github.com/users/octocat/followers",
        "following_url": "https://api.github.com/users/octocat/following{/other_user}",
        "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
        "organizations_url": "https://api.github.com/users/octocat/orgs",
        "repos_url": "https://api.github.com/users/octocat/repos",
        "events_url": "https://api.github.com/users/octocat/events{/privacy}",
        "received_events_url": "https://api.github.com/users/octocat/received_events",
        "type": "User",
        "site_admin": false
      },
      "name": "Hello-World",
      "full_name": "octocat/Hello-World",
      "description": "This your first repo!",
      "private": false,
      "fork": false,
      "url": "https://api.github.com/repos/octocat/Hello-World",
      "html_url": "https://github.com/octocat/Hello-World"
    },
    "subject": {
      "title": "Greetings",
      "url": "https://api.github.com/repos/octokit/octokit.rb/issues/123",
      "latest_comment_url": "https://api.github.com/repos/octokit/octokit.rb/issues/comments/123",
      "type": "Issue"
    },
    "reason": "subscribed",
    "unread": true,
    "updated_at": "2014-11-07T22:01:45Z",
    "last_read_at": "2014-11-07T22:01:45Z",
    "url": "https://api.github.com/notifications/threads/1"
  }
]
*/

type notification struct {
	id         float64    `json:"id"`
	repository repository `json:"repository"`
	subject    subject    `json:"subject"`
	reason     string     `json:"reason"`
	unread     bool       `json:"unread"`
	updatedAt  time.Time  `json:"updated_at"`
	lastReadAt time.Time  `json:"last_read_at"`
	url        string     `json:"url"`
}

type repository struct {
	id          int    `json:"id"`
	owner       owner  `json:"owner"`
	name        string `json:"name"`
	fullName    string `json:"full_name"`
	description string `json:"description"`
	private     bool   `json:"private"`
	fork        bool   `json:""fork"`
	url         string `json:""url"`
	htmlURL     string `json:"html_url"`
}

type subject struct {
	title            string `json:"title"`
	url              string `json:"url"`
	latestCommentURL string `json:"latest_comment_url"`
	subjectType      string `json:"type"`
}

type owner struct {
	login             string `json:"login"`
	id                int    `json:"id"`
	avatarURL         string `json:"avatar_url"`
	gravatarID        string `json:"gravatar_id"`
	url               string `json:"url"`
	htmlURL           string `json:"html_url"`
	followersURL      string `json:"followers_url"`
	followingURL      string `json:"following_url"`
	gistsURL          string `json:"gists_url"`
	starredURL        string `json:""starred_url"`
	subscriptionsURL  string `json:"subscriptions_url"`
	organizationsURL  string `json:"organizations_url"`
	reposURL          string `json:"repos_url"`
	eventsURL         string `json:"events_url"`
	receivedEventsURL string `json:"received_events_url"`
	ownerType         string `json:"type"`
	siteAdmin         bool   `json:"site_admin"`
}
