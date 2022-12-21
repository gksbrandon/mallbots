dc-up:
	docker-compose up -d 

dc-down:
	docker-compose down --rmi all

dc-build:
	docker-compose build monolith

dc-reset: compose-down compose-up

dc-update:
	docker-compose down
	docker-compose build monolith
	docker-compose up -d

dc-logs:
	docker-compose logs -f
