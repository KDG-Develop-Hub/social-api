run:
	air
install:
	go mod tidy
	go get
	go install github.com/rubenv/sql-migrate/...@latest
	go install github.com/cosmtrek/air@latest