all: run

run:
	docker-compose -f ops/docker-compose.yml up -d

stop:
	docker-compose -f ops/docker-compose.yml down
