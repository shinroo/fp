build:
	pushd scrapers && make build && popd
	pushd scrapers/amanzimtoti && make build && popd
	pushd scrapers/cape && make build && popd
	pushd webapp && make build && popd
	
up:
	docker-compose -p spca_pet_matchmaker -f airflow/docker-compose.yaml -f database/docker-compose.yaml -f webapp/docker-compose.yaml up -d

down:
	docker-compose -p spca_pet_matchmaker -f airflow/docker-compose.yaml -f database/docker-compose.yaml -f webapp/docker-compose.yaml down

