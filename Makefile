cfg=env
pkg=mux

setup:
	cp configs/.env.example .env
	cp configs/example_configs.yml config.yml

# example make start-server cfg=env pkg=echo (or exclude the arguments)
start-server:
	go run main.go $(cfg) $(pkg)

start-docker:
	docker-compose up --remove-orphans

stop-docker:
	docker-compose down

# can ignore pkg argument
migrate-db:
	go run internal/app/job/userdb/userdb.go $(cfg) $(pkg)

clear-log:
	cp /dev/null log/server.log