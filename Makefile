APP_BINARY=app_binary

up: 
	@echo "starting docker images"
	docker compose up -d
	@echo "docker images started"

up_build: build_app
	@echo "stopping any running docker images"
	docker compose down
	@echo "building and starting docker images"
	docker compose up -d --build
	@echo "docker images built and started"

down:
	@echo "stopping docker images"
	docker compose down
	@echo "docker images stopped"

build_app:
	@echo "building app"
	go build -o $(APP_BINARY) .
	@echo "app built"