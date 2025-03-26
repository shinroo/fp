build_base:
	pushd scrapers && make build && popd

build_scrapers:
	pushd scrapers/amanzimtoti && make build && popd
	pushd scrapers/cape && make build && popd
	pushd scrapers/durban_and_coast && make build && popd
	pushd scrapers/lower_south_coast && make build && popd
	pushd scrapers/randburg && make build && popd

build_webapp:
	pushd webapp && make build && popd

build: build_base build_scrapers build_webapp
	
up:
	docker-compose -p spca_pet_matchmaker \
		-f database/docker-compose.yaml \
		-f airflow/docker-compose.yaml \
		-f webapp/docker-compose.yaml \
		--project-directory . \
		up -d

down:
	docker-compose -p spca_pet_matchmaker \
		-f database/docker-compose.yaml \
		-f airflow/docker-compose.yaml \
		-f webapp/docker-compose.yaml \
		--project-directory . \
		down

