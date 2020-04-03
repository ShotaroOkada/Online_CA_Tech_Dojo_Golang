DEV_DOCKER_COMPOSE_FILE='./deployments/docker-compose.dev.yml'

up-user-dev:
	sudo docker-compose -f $(DEV_DOCKER_COMPOSE_FILE) up -d mongo mongo-express
	sleep 5
	sudo docker-compose -f $(DEV_DOCKER_COMPOSE_FILE) up -d user
	sudo docker-compose -f $(DEV_DOCKER_COMPOSE_FILE) logs -f user