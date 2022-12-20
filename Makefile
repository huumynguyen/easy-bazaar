run:
	go run .
up:
	docker-compose up --build
down:
	docker-compose down
docker-build:
	docker build -t easy-bazaar:1.0 .                                