#!/bin/bash

export NOW=$(shell date +"%Y/%m/%d")

configure:
	@echo "${NOW} === CONFIGURING FILES ==="
	@cp ./conf/app.ini.example conf/app.ini
	@echo "${NOW} === CONFIGURED ==="

generate:
	@echo "${NOW} === GENERATING FILES ==="
	@go generate ./...
	@echo "${NOW} === GENERATED ==="

.PHONY: dev
dev:
	@echo "${NOW} === RUNNING DEVELOPMENT ENV ==="
	@docker-compose stop goauth-be goauth-fe && docker-compose up -d goauth-be goauth-fe
	@echo "click this link to open the backend http://localhost:8000"
	@echo "click this link to open the frontend http://localhost:8001"

dev-all:
	@echo "${NOW} === RUNNING DEVELOPMENT ALL ==="
	@cd tools/ && docker-compose stop && docker-compose up -d
	@docker-compose stop && docker-compose up -d

dev-fe:
	@echo "${NOW} === RUNNING DEVELOPMENT ENV ==="
	@docker-compose stop goauth-fe && docker-compose up -d goauth-fe
	@echo "click this link to open the page http://localhost:8001"

dev-tools:
	@echo "${NOW} === RUNNING DEVELOPMENT TOOLS ==="
	@cd tools/ && docker-compose stop && docker-compose up -d
	@echo "click this link to open yacht page http://localhost:5000"
	@echo "click this link to open prometheus http://localhost:5001"

clean:
	@echo "🛠 CLEANING MACHINE FOR DEVELOPMENT 🛠"
	@echo "1⃣ REMOVING BIN FOLDER"
	@rm -r ./bin
	@echo "🚀 Done, You are ready to Go 🚀"

down:
	@docker-compose stop
	@cd tools/ && docker-compose stop
	@docker-compose down
	@cd tools/ && docker-compose down

down-tools:
	@cd tools/ && docker-compose stop
	@cd tools/ && docker-compose down