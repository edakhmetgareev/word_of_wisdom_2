build: stop build.server build.client create.network run.server create.client

stop:
	docker stop server || true && docker stop client || true

build.server:
	docker build -t server -f Dockerfile_server .

build.client:
	docker build -t client -f Dockerfile_client .

create.network:
	docker network rm app_network || true && \
	docker network create app_network --driver bridge

run.server:
	docker stop server || true && \
	docker run --rm --network app_network \
		-p 8000:8000 \
		--name server -itd \
		-e HTTP_PORT=8000 \
		-e CHALLENGE_TTL=1m \
	server

create.client:
	docker rm client || true && \
	docker create --network app_network \
		--name client -i \
		-e SERVER_HOST=server \
		-e HTTP_PORT=8000 \
	client

start:
	docker start -i client