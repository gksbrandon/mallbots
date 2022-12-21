compose-up:
	docker-compose up -d 

compose-down:
	docker-compose down --rmi all

compose-build:
	docker-compose build monolith

compose-reset: compose-down compose-up

compose-update:
	docker-compose down
	docker-compose build monolith
	docker-compose up -d

compose-logs:
	docker-compose logs -f
