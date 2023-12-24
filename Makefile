APP_NAME=new_app
BINARY_PATH = $(GOPATH)/bin
DSN = postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable

migrate-setup:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate-clean:
	export PATH=$(BINARY_PATH):$$PATH && migrate -path ./pkg/database/migrations -database "$(DSN)" force 1

migrate-up:
	export PATH=$(BINARY_PATH):$$PATH && migrate -path ./pkg/database/migrations -database "$(DSN)" up

migrate-down:
	export PATH=$(BINARY_PATH):$$PATH && migrate -path ./pkg/database/migrations -database "$(DSN)" down

build:
	go build -o bin/$(APP_NAME) cmd/app/main.go

run: build
	./bin/$(APP_NAME)

clean:
	rm -f $(APP_NAME)