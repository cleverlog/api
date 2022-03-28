all: run_clickhouse

run_clickhouse:
	docker-compose -f ops/docker-compose.yml up
