# Docker Compose parameters
# Start the Docker Compose stack in detached mode
dev-stack:
	docker-compose -f docker-compose-dev.yaml up -d

# Stop and remove the Docker Compose stack
dev-stack-down:
	docker-compose -f docker-compose-dev.yaml down

# Build or rebuild services in the Docker Compose stack
stack-build:
	docker-compose build 

stack:
	docker-compose up -d

# Stop and remove the Docker Compose stack
stack-down:
	docker-compose down

# Build or rebuild services in the Docker Compose stack
dev-stack-build:
	docker-compose -f docker-compose-dev.yaml build 

# dev environment
dev: dev-stack-build dev-stack
dev-down: dev-stack-down

# prod environment
prod: stack-build dev-stack
prod-down: stack-down