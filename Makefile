postgres:
	docker run --name postgres-gorm -p 5432:5432 -e POSTGRES_PASSWORD=secret-gorm -d postgres

postgres-stop:
	docker stop postgres-gorm

postgres-remove:
	docker rm postgres-gorm

postgres-logs:
	docker logs postgres-gorm

postgres-connect:
	docker exec -it postgres-gorm psql -U postgres

run:
	go run ./cmd/main.go

.PHONY : postgres postgres-stop postgres-remove postgres-logs postgres-connect run