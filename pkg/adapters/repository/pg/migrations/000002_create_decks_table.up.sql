CREATE TABLE decks(
    id bigserial PRIMARY KEY,
    external_id varchar UNIQUE,
    is_shuffled boolean NOT NULL
);

CREATE TABLE decks_cards(
    card_id int,
    deck_id int,
    is_drew boolean NOT NULL DEFAULT false,
    CONSTRAINT fk_decks_cards_cards FOREIGN KEY(card_id) REFERENCES cards(id),
    CONSTRAINT fk_decks_cards_decks FOREIGN KEY(deck_id) REFERENCES decks(id)
);
