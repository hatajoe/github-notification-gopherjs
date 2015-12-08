# github-notification-gopherjs

This app can receive GitHub notifications as desktop notification on darwin.
This is implemented by [gopherjs](https://github.com/gopherjs/gopherjs) + [electron](http://electron.atom.io)

## How to use

Create GitHub Personal access token. (don't forget to enable `notifications` check box)
And create specified file to your `$HOME/.github-notification-gopherjs`

```
{
    "interval": 3,
    "token": "Personal access token"
}
```

- interval: GitHub API execution interval seconds
- token: Personal access token

Run `GitHub-Notification-GopherJS-darwin-x64/GitHub-Notification-GopherJS.app`

## Run as debug

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
