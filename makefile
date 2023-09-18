.PHONY: build
build:
	@docker build -t final-esp-backend-3-grupo-1 -f docker/Dockerfile .

.PHONY: start
start:
	@docker-compose up
	@echo "Database and application started!"
