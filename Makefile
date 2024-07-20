run: build
	go run ./server/main.go

build:
	#buf lint
	#buf check breaking --against-input '.git#branch=main'
	buf generate