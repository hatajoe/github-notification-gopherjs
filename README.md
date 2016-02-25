# github-notification-gopherjs

This app can receive GitHub notifications as desktop notification on darwin.
This is implemented by [gopherjs](https://github.com/gopherjs/gopherjs) + [electron](http://electron.atom.io)

## How to use

Create GitHub Personal access token. (don't forget to enable `notifications` check box)
And create specified file to your `$HOME/.github-notification-gopherjs`

```
// For github.com
{
    "interval": 3,
    "token": "Personal access token",
    "githostname": "github.com",
    "apihostname": "api.github.com"
}

For GithubEnterprise Example
{
    "interval": 3,
    "token": "Personal access token",
    "githostname": "ghehost.example.com",
    "apihostname": "ghehost.example.com/api/v3"
}
```

- interval: GitHub API execution interval seconds
- token: Personal access token

Execute `GitHub-Notification-GopherJS-darwin-x64/GitHub-Notification-GopherJS.app`

## Run as debug

Install gopherjs

```
go get -u github.com/gopherjs/gopherjs
```

Install npm and electron

```
% brew install npm 
% npm -g install electron-prebuilt
```

Clone repository and run

```
% git clone git@github.com:hatajo/github-notification-gopherjs
% cd github-notification-gopherjs
% npm install
% make clean && make all
% electron .
```

## Package Build

Install electron-packager

```
% npm -g install electron-packager
% electron-packager . GitHub-Notification-GopherJS --platform=darwin --arch=x64 --version=0.35.4
% open GitHub-Notification-GopherJS
```

## License

MIT
