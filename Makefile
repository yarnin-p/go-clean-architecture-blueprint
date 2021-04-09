build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o main .

run: build
	docker-compose -f docker-compose.yml up -d --build
