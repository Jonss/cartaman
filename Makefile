run:
	go run cmd/cartaman/main.go

test:
	go test -v --cover ./...

build-mocks:
# usecases
	mockgen -destination pkg/usecases/decks/mocks/create_deck_usecase.go github.com/Jonss/cartaman/pkg/usecases/decks DeckUseCase
# repositories

