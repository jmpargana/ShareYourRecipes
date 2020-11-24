help:
	cat Makefile

build:
	docker-compose build

run:
	docker-compose up -d

clean:
	sudo rm -R data/
