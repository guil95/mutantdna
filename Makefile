migrate:
	docker run -ti --rm \
		--name mutant-migrate \
		--network mutant \
		-v $(PWD)/migrations:/migrations \
		migrate/migrate:v4.14.1 \
		-path=/migrations/ -database postgres://mutant_user:mut4ant@database:5432/mutant?sslmode=disable up

start:
	go mod vendor
	go mod tidy
	docker-compose up --build -d

start-migrate:
	make start
	make migrate

test:
	touch count.out
	go test -covermode=count -coverprofile=count.out -v ./...
	$(MAKE) coverage

coverage:
	go tool cover -func=count.out

mock-generate:
	go mod vendor
	docker run --rm -v "$(PWD):/app" -w /app -t vektra/mockery --all --dir ./internal/domain --case underscore
