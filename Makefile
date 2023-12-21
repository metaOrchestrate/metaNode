# Docker Compose parameters
# Start the Docker Compose stack in detached mode
stack:
	docker-compose up

# Stop and remove the Docker Compose stack
stack-down:
	docker-compose down

# Build or rebuild services in the Docker Compose stack
stack-build:
	docker-compose build	