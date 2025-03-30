build:
	docker build --no-cache --platform=linux/amd64 --network host -t scraper_base:latest .

test:
	python3 -m unittest -v breed_identifier_test.py