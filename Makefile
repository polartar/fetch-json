run:
	go run main.go
build:
	set GOOS=macos && go build -o main
 
compose:
	docker-compose up -d
 