NGINX_FILE_PATH = $(shell readlink -f out/nginx.conf)
COMPOSE_PATH = "out/docker-compose.yml"

conf-generator:
	go run cmd/main.go
stop: 
	@NGINX_FILE=$(NGINX_FILE_PATH) docker-compose -f $(COMPOSE_PATH) down
clean:
	-@NGINX_FILE=$(NGINX_FILE_PATH) docker-compose -f $(COMPOSE_PATH) down --volumes --remove-orphans
nginx: clean conf-generator
	@NGINX_FILE=$(NGINX_FILE_PATH) docker-compose -f $(COMPOSE_PATH) up -d
	

