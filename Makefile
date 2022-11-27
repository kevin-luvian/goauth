#!/bin/bash

export NOW=$(shell date +"%Y/%m/%d")

.PHONY: dev
dev:
	@echo "${NOW} === RUNNING DEVEL ENV ==="
	@docker-compose stop gogin && docker-compose up -d gogin yacht
	@echo "click this link to open the page http://localhost:8000"

configure:
	@echo "🛠 CONFIGURING YOUR MACHINE FOR DEVELOPMENT 🛠"
	@echo "1⃣ SETUP DOCKER NETWORKS"
	@echo "🚀 Done, You are ready to Go 🚀"

clean:
	@echo "🛠 CLEANING MACHINE FOR DEVELOPMENT 🛠"
	@echo "1⃣ REMOVING BIN FOLDER"
	@rm -r ./bin
	@echo "🚀 Done, You are ready to Go 🚀"

down:
	@docker-compose down