DB_PORT=8086
DB_NAME=mydb

init-db:
	curl -i -XPOST http://localhost:$(DB_PORT)/query --data-urlencode "q=CREATE DATABASE $(DB_NAME)"

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down
