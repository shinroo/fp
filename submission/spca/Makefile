migrations:
	python3 manage.py makemigrations
	python3 manage.py migrate

nuke_db:
	rm db.sqlite3
	touch db.sqlite3

run_dev:
	python3 manage.py runserver 0.0.0.0:8090

freeze_requirements:
	pip freeze > requirements.txt

run_tests:
	python3 manage.py test --force-color -v 2