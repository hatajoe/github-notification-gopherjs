all: main/index.go renderer/github.go
	gopherjs build renderer/github.go -o github.js
	gopherjs build main/index.go -o index.js

clean: 
	rm -f *.js && rm -f *.map
