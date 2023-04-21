run:
	go run cmd/cartaman/main.go

test:
	go test -v --cover ./...

build-mocks:
# usecases
	mockgen -destination pkg/usecase/decks/mocks/deck_service.go github.com/Jonss/cartaman/pkg/usecase/decks DeckService
# repositories


new-migration: # new-migration name=migration_name
	migrate create -ext sql -dir pkg/adapters/repository/pg/migrations -seq $(name)
