run: # run the application locally
	go run cmd/cartaman/main.go

test: # test the application files
	go test -v --cover ./...

build-mocks:
# usecases
	mockgen -destination pkg/usecase/deck/mocks/deck_service.go github.com/Jonss/cartaman/pkg/usecase/deck DeckService

new-migration: # new-migration name=migration_name
	migrate create -ext sql -dir pkg/adapters/repository/pg/migrations -seq $(name)

env-up: # starts dependencies, postgres only
	docker-compose up --build -d db

run-docker: env-up # starts the db, then run the app in a docker container
	docker-compose up --build app

cover:
	go test -v ./... -covermode=count -coverpkg=./... -coverprofile coverage.out
	go tool cover -html coverage.out -o coverage.html
	open coverage.html