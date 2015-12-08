all: main/index.go renderer/github.go renderer/structs.go
	gopherjs build renderer/github.go renderer/structs.go -o github.js
	gopherjs build main/index.go -o index.js

clean: 
	rm -f *.js && rm -f *.map
