.PHONY: build start stop

build: stop build-server build-client create-network run-server create-client

stop:
	@docker stop word_of_wisdom_server word_of_wisdom_client 2>/dev/null || true
	@docker rm word_of_wisdom_client 2>/dev/null || true

build-server:
	docker build -t word_of_wisdom_server -f Dockerfile_server .

build-client:
	docker build -t word_of_wisdom_client -f Dockerfile_client .

create-network:
	docker network rm word_of_wisdom_network 2>/dev/null || true
	docker network create word_of_wisdom_network --driver bridge

run-server:
	docker run --rm --network word_of_wisdom_network -p 8080:8080 --name word_of_wisdom_server -itd -e SERVER_PORT=8080 word_of_wisdom_server

create-client:
	docker rm word_of_wisdom_client 2>/dev/null || true
	docker create --network word_of_wisdom_network --name word_of_wisdom_client -i -e SERVER_HOST=word_of_wisdom_server -e SERVER_PORT=8080 word_of_wisdom_client

start:
	@sleep 1
	docker start -i word_of_wisdom_client
