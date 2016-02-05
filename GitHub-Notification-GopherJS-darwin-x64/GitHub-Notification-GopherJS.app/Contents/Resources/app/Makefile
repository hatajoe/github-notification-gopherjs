all: main/index.go renderer/github.go
	gopherjs build renderer/github.go -o github.js
	gopherjs build main/index.go -o index.js
	electron-packager . GitHub-Notification-Gopherjs --platform=darwin --arch=x64 --version=0.34.3

clean: 
	rm -f *.js && rm -f *.map && rm -rf GitHub-Notification-Gopherjs-darwin-x64/
