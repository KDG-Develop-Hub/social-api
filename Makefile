run:
	air
install:
	go mod tidy
	go get
	go install github.com/cosmtrek/air@latest