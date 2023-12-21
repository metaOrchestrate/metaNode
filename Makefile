# Variable to hold the tag value. It can be overridden when calling make.
# Example usage: make create TAG=v0.12.1
TAG ?= v0.12.1 # <- choose your tag release 

networks:
	git remote add -f networks git@github.com:celestiaorg/networks.git
	git subtree add --prefix networks networks --squash
	git commit -m "Add celestia-networks subtree"

# Add the remote repository as a new subtree in the directory 'app'.
create:
	git remote add -f celestia-app git@github.com:celestiaorg/celestia-app.git
	git subtree add --prefix app celestia-app tags/$(TAG) --squash
	git commit -m "Add celestia-app subtree at tags/$(TAG)"

# Fetch and merge updates from the metalgo remote repository's specified tag into the 'validator' subtree.
sync:
	git fetch celestia-app $(TAG) 
	git subtree pull --prefix app celestia-app tags/$(TAG) --squash # Pull changes from metalgo's specified tag into the validator subtree, squashing commits.
	git commit -m "Sync app subtree with tags/$(TAG)" || true    # Commit the changes to sync the metalgo subtree, ignoring if there are no changes.

# Remove the 'validator' subtree and its associated remote.
remove:
	git rm -r app
	git commit -m "Remove validator folder"
	git remote rm celestia-app

# Run the Docker container from the pulled image
pull:
	@echo "Pulling ghcr.io/metaorchestrate/metalgo:$(TAG)..."
	@docker pull ghcr.io/metaorchestrate/metalgo:$(TAG) || echo "Failed to pull the image with tag $(TAG). Please check the tag exists and you have permissions to pull the image."

magic: pull
	@echo "Running metalgo instance..."
	@docker run -ti --name validator -p 127.0.0.1:9650:9650 -p 127.0.0.1:9651:9651 -v metalgo:/root/.metalgo -e API_PASSWORD=as0d98d9f8s09f8da098 ghcr.io/metaorchestrate/metalgo:$(TAG) || echo "Failed to run the image with tag $(TAG). Please check the Docker run command and the image tag."

# Docker Compose parameters
# Start the Docker Compose stack in detached mode
stack:
	docker-compose up -d

# Stop and remove the Docker Compose stack
stack-down:
	docker-compose down

# Build or rebuild services in the Docker Compose stack
stack-build:
	docker-compose build	