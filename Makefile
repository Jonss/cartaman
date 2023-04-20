run:
	go run cmd/cartaman/main.go

test:
	go test -v --cover ./...

build-mocks:
# usecases
	mockgen -destination pkg/usecase/decks/mocks/create_deck_usecase.go github.com/Jonss/cartaman/pkg/usecase/decks DeckUseCase
# repositories


new-migration: # new-migration name=migration_name
	migrate create -ext sql -dir pkg/adapters/repository/postgres/migrations -seq $(name)
