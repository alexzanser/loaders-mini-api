
build:
		docker-compose up --build

start:
		docker-compose start

loaders_register:
		curl -d "username=VasyaIvanova&password=1234&role=loader" -H "Content-Type: application/x-www-form-urlencoded" -X POST http://localhost:8080/register
		curl -d "username=PetyaSidorov&password=1234&role=loader" -H "Content-Type: application/x-www-form-urlencoded" -X POST http://localhost:8080/register
		curl -d "username=KolyaFedorov&password=1234&role=loader" -H "Content-Type: application/x-www-form-urlencoded" -X POST http://localhost:8080/register

loaders_login:
		curl -d "username=VasyaIvanova&password=1234&role=loader" -H "Content-Type: application/x-www-form-urlencoded" -X POST http://localhost:8080/login
		curl -d "username=PetyaSidorov&password=1234&role=loader" -H "Content-Type: application/x-www-form-urlencoded" -X POST http://localhost:8080/login
		curl -d "username=KolyaFedorov&password=1234&role=loader" -H "Content-Type: application/x-www-form-urlencoded" -X POST http://localhost:8080/login

customer_register:
		curl -d "username=JackBlack&password=1234&role=customer&balance=1000000" -H "Content-Type: application/x-www-form-urlencoded" -X POST http://localhost:8080/register

customer_login:
		curl -d "username=JackBlack&password=1234&role=customer&balance=1000000" -H "Content-Type: application/x-www-form-urlencoded" -X POST http://localhost:8080/login

generate_tasks:
		curl -X POST http://localhost:8080/tasks

stop:
		docker-compose stop

clean:
		docker rm loaders db

.PHONY: build, start, stop, clean